package manifest

// Pool of compute and network resources allocated by the blueprint builder.
//
// The blueprint builder populates the pool based on the compute and network
// requirements declared by services in the blueprint. The pool is included in
// the plan so that the executor can provision the resources and apply the
// configurations when deploying the services.
type Infrastructure struct {

	// Allocated machines, by compute entry ID.
	//
	// Each compute entry describes a machine where services run. Deployments
	// refer to entries by ID. The executor provisions the machines and
	// schedules services onto them according to the deployments.
	Computes map[string]Compute `codec:"computes,omitempty"`

	// Allocated network configurations.
	//
	// Each network configuration defines the network settings for a service.
	// The map associates network configuration IDs with their settings, which
	// can be referenced by deployments. The executor applies the configurations
	// to the services' containers according to the deployment specifications.
	Networks map[string]Network `codec:"networks,omitempty"`
}

// Validates all compute and network entries in the pool.
func (i *Infrastructure) Validate() error {
	for _, c := range i.Computes {
		if err := c.Validate(); err != nil {
			return err
		}
	}
	for _, n := range i.Networks {
		if err := n.Validate(); err != nil {
			return err
		}
	}
	return nil
}
