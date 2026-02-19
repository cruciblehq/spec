package plan

import "github.com/cruciblehq/crex"

// Current plan format version.
const Version = 0

// Represents a deployment plan.
//
// Specifies what resources will be deployed and the infrastructure configuration
// required to run them. Generated during the planning phase by resolving
// references, allocating infrastructure, and determining routing.
type Plan struct {
	Version      int           `json:"version"`                // Schema version of the plan format.
	Services     []Service     `json:"services"`               // Services included in the deployment.
	Compute      []Compute     `json:"compute"`                // Compute resources to provision.
	Environments []Environment `json:"environments,omitempty"` // Environment variable sets for service configuration.
	Bindings     []Binding     `json:"bindings"`               // Associations between services and compute resources.
	Gateway      Gateway       `json:"gateway"`                // Gateway routing configuration.
}

// Validates the plan.
//
// The version must match [Version]. At least one service and one compute
// resource are required. Every service must have an ID and reference,
// every compute must have an ID and provider, every binding must
// reference a service and compute, and every route must have a pattern
// and service.
func (p *Plan) Validate() error {
	if p.Version != Version {
		return crex.Wrap(ErrInvalidPlan, ErrUnsupportedVersion)
	}

	if len(p.Services) == 0 {
		return crex.Wrap(ErrInvalidPlan, ErrMissingServices)
	}

	for i := range p.Services {
		if err := p.Services[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidPlan, err)
		}
	}

	if len(p.Compute) == 0 {
		return crex.Wrap(ErrInvalidPlan, ErrMissingCompute)
	}

	for i := range p.Compute {
		if err := p.Compute[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidPlan, err)
		}
	}

	for i := range p.Bindings {
		if err := p.Bindings[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidPlan, err)
		}
	}

	for i := range p.Gateway.Routes {
		if err := p.Gateway.Routes[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidPlan, err)
		}
	}

	return nil
}
