package manifest

import "github.com/cruciblehq/utils-go/crex"

// Current plan format version.
const PlanVersion = 0

// Represents a deployment plan.
//
// Specifies what resources will be deployed and the infrastructure required to
// run them. Generated during the planning phase by resolving references and
// determining routing.
//
// Infrastructure holds separate compute and network pools, each keyed by ID.
// Containers and Environments are lookup maps keyed by ID. Deployments is an
// ordered slice that binds them into deployments.
type Plan struct {
	Version        int                    `codec:"version"`                // Version of the plan format.
	Services       map[string]string      `codec:"services,omitempty"`     // Service lookup table. Key is the service ID, value is the reference target.
	Infrastructure Infrastructure         `codec:"infrastructure"`         // Infrastructure resource pools for compute and network allocations.
	Containers     map[string]Container   `codec:"containers,omitempty"`   // Container spec lookup table. Key is the container ID.
	Environments   map[string]Environment `codec:"environments,omitempty"` // Environment variable set lookup table. Key is the environment ID.
	Deployments    []Deployment           `codec:"deployments,omitempty"`  // Ordered deployment arrangement.
	Gateway        Gateway                `codec:"gateway"`                // Gateway routing configuration.
}

// Validates the plan.
//
// The version must match [PlanVersion]. Every service target must be set.
// Infrastructure entries, environments, and containers are validated
// individually. Each deployment's field values are validated, then its
// service container, compute, network, and environment references are
// checked against the plan's lookup tables. The gateway is validated last.
func (p *Plan) Validate() error {
	if p.Version != PlanVersion {
		return crex.Wrap(ErrInvalidPlan, ErrUnsupportedPlanVersion)
	}

	for _, target := range p.Services {
		if target == "" {
			return crex.Wrap(ErrInvalidPlan, ErrMissingRefTarget)
		}
	}

	if err := p.Infrastructure.Validate(); err != nil {
		return crex.Wrap(ErrInvalidPlan, err)
	}

	for _, e := range p.Environments {
		if err := e.Validate(); err != nil {
			return crex.Wrap(ErrInvalidPlan, err)
		}
	}

	for id, c := range p.Containers {
		if err := c.Validate(); err != nil {
			return crex.Wrapf(ErrInvalidPlan, err, "container %q", id)
		}
	}

	for i := range p.Deployments {
		if err := p.validateDeployment(&p.Deployments[i]); err != nil {
			return err
		}
	}

	if err := p.Gateway.Validate(); err != nil {
		return crex.Wrap(ErrInvalidPlan, err)
	}

	return nil
}

// Validates a deployment's field values and checks that all of its binding
// references resolve to entries in the plan's lookup tables.
func (p *Plan) validateDeployment(d *Deployment) error {
	if err := d.Validate(); err != nil {
		return crex.Wrap(ErrInvalidPlan, err)
	}
	if _, ok := p.Services[d.Service]; !ok {
		return crex.Newf(ErrInvalidPlan, "deployment references undefined service %q", d.Service)
	}
	if _, ok := p.Containers[d.Container]; !ok {
		return crex.Newf(ErrInvalidPlan, "deployment references undefined container %q", d.Container)
	}
	if _, ok := p.Infrastructure.Computes[d.Compute]; !ok {
		return crex.Newf(ErrInvalidPlan, "deployment references undefined compute %q", d.Compute)
	}
	if _, ok := p.Infrastructure.Networks[d.Network]; !ok {
		return crex.Newf(ErrInvalidPlan, "deployment references undefined network %q", d.Network)
	}
	if d.Environment != "" {
		if _, ok := p.Environments[d.Environment]; !ok {
			return crex.Newf(ErrInvalidPlan, "deployment references undefined environment %q", d.Environment)
		}
	}
	return nil
}
