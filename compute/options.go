package compute

import (
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Controls how a container process is executed.
//
// Each field configures one aspect of the process's environment. Fields with
// zero values fall back to defaults from the OCI spec or image configuration.
type Options struct {
	OCI     specs.Spec        `codec:"oci,omitempty"`     // OCI runtime spec applied as the base configuration at container start.
	Shell   string            `codec:"shell,omitempty"`   // Shell for "shell -c command" invocation; defaults to /bin/sh.
	Env     map[string]string `codec:"env,omitempty"`     // Environment variables merged into the spec env.
	Workdir string            `codec:"workdir,omitempty"` // Working directory for the process.
	User    string            `codec:"user,omitempty"`    // User identity as numeric "uid" or "uid:gid".

	// The network baseline for containers is a restricted network. This field
	// relaxes that restriction to use the host network.
	//
	// Note: this is just a placeholder for future network policy controls. For
	// now, it enables the use of the host network while Crucible's security
	// policy is not implemented.
	HostNetwork bool `codec:"host_network,omitempty"`
}
