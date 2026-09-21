package manifest

import (
	"github.com/cruciblehq/utils-go/crex"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Compiled enforcement model for containers.
//
// This model is stored in the plan's Containers map and covers the enforcement
// surface that runc applies at container creation time (OCI runtime spec,
// container-level network spec, and related runtime settings).
type Container struct {

	// Full OCI runtime spec passed to runc at container creation.
	//
	// Contains namespaces, capabilities, seccomp filter, cgroup device rules,
	// mounts, environment, and every other runc-level knob. This is Populated
	// and narrowed by the OCI subsystems (cap, seccomp, cgroup, rlimit, etc).
	OCI specs.Spec `codec:"oci"`

	// Container-level network rules.
	//
	// These entries are derived from the deployment topology and describe the
	// allowed inbound and outbound traffic for the service. The rules are
	// validated as part of plan validation.
	Network Network `codec:"network"`
}

// Validates the compiled container enforcement spec.
func (c *Container) Validate() error {
	if err := c.Network.Validate(); err != nil {
		return crex.Wrap(ErrInvalidContainer, err)
	}
	return nil
}
