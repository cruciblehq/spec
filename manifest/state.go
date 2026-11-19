package manifest

import "github.com/cruciblehq/utils-go/crex"

// Current state format version.
const StateVersion = 0

// Represents the current state of a deployment.
//
// Records what resources have been deployed and the identifiers their provider
// assigned. The executor writes the state after each provisioning step
// and reads it back during updates and teardowns. The state is the single
// source of truth for what exists on the provider.
type State struct {
	Version     int                `codec:"version"`               // Version of the state format.
	Deployments []Deployment       `codec:"deployments"`           // Deployments that were applied.
	Instances   []InstanceRecord   `codec:"instances,omitempty"`   // Compute instances the provider created.
	Volumes     []VolumeRecord     `codec:"volumes,omitempty"`     // Volumes the provider created.
	Networks    []NetworkRecord    `codec:"networks,omitempty"`    // Networks the provider created.
	Attachments []AttachmentRecord `codec:"attachments,omitempty"` // What is attached to each network.
}

// Records a provisioned compute instance.
type InstanceRecord struct {
	ID       string `codec:"id"`       // Instance identifier the provider assigned.
	Image    string `codec:"image"`    // Image identifier used to provision the instance.
	Provider string `codec:"provider"` // Provider type (e.g. "local", "aws").
	Compute  string `codec:"compute"`  // Plan compute ID this instance fulfills.
}

// Records a provisioned volume.
type VolumeRecord struct {
	ID       string `codec:"id"`                 // Volume identifier the provider assigned.
	Instance string `codec:"instance,omitempty"` // Instance the volume is attached to; empty if detached.
	Size     int    `codec:"size,omitempty"`     // Allocated capacity in MiB.
	Service  string `codec:"service"`            // Plan service ID this volume belongs to.
}

// Records a provisioned network.
type NetworkRecord struct {
	ID     string `codec:"id"`     // Network identifier the provider assigned.
	Subnet string `codec:"subnet"` // Assigned CIDR.
}

// Records a container attached to a network.
type AttachmentRecord struct {
	Container string `codec:"container"`      // Container identifier.
	Network   string `codec:"network"`        // Network identifier from [NetworkRecord.ID].
	IP        string `codec:"ip,omitempty"`   // Assigned IP address within the network subnet.
	Veth      string `codec:"veth,omitempty"` // Name of the veth interface on the host (local provider).
	Peer      string `codec:"peer,omitempty"` // Name of the veth interface in the network namespace (local provider).
}

// Validates the state.
//
// The version must match [StateVersion]. Deployments, instances, volumes,
// networks, and attachments are validated for required fields.
func (s *State) Validate() error {
	if s.Version != StateVersion {
		return crex.Wrap(ErrInvalidState, ErrUnsupportedStateVersion)
	}

	if err := validateStateDeployments(s.Deployments); err != nil {
		return err
	}
	if err := validateStateRecords(s.Instances, "instance", func(r *InstanceRecord) bool {
		return r.ID != "" && r.Image != "" && r.Provider != "" && r.Compute != ""
	}); err != nil {
		return err
	}
	if err := validateStateRecords(s.Volumes, "volume", func(r *VolumeRecord) bool {
		return r.ID != "" && r.Service != ""
	}); err != nil {
		return err
	}
	if err := validateStateRecords(s.Networks, "network", func(r *NetworkRecord) bool {
		return r.ID != "" && r.Subnet != ""
	}); err != nil {
		return err
	}
	if err := validateStateRecords(s.Attachments, "attachment", func(r *AttachmentRecord) bool {
		return r.Container != "" && r.Network != ""
	}); err != nil {
		return err
	}

	return nil
}

// Validates all deployment entries in the state.
func validateStateDeployments(deployments []Deployment) error {
	for i := range deployments {
		if err := deployments[i].Validate(); err != nil {
			return crex.Wrap(ErrInvalidState, err)
		}
	}
	return nil
}

// Checks each record of one kind with valid.
//
// A failure names kind and the record's position, counting from 1.
func validateStateRecords[T any](records []T, kind string, valid func(*T) bool) error {
	for i := range records {
		record := &records[i]
		if !valid(record) {
			return crex.At(crex.Wrap(ErrInvalidState, ErrMissingRequiredField), kind, i+1)
		}
	}
	return nil
}
