package manifest

import (
	"github.com/cruciblehq/utils-go/crex"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// Compiled enforcement model for containers.
//
// This model is stored in the plan's Containers map and covers what the OCI
// runtime applies when it creates the container: the OCI runtime spec, the
// network rules of the container, and related settings.
type Container struct {

	// Full OCI runtime spec passed to the OCI runtime at container creation.
	//
	// Contains namespaces, capabilities, seccomp filter, cgroup device rules,
	// mounts, environment, and every other setting the OCI runtime applies.
	OCI specs.Spec `codec:"oci"`

	// Network rules of the container.
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
