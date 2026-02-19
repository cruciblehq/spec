package blueprint

import "github.com/cruciblehq/crex"

// Defines a system composition.
//
// A Blueprint declares a service composition and configuration, being the main
// input to the deployment process. Blueprints prescribe which resources should
// be deployed together and how they are composed and exposed.
type Blueprint struct {

	// Schema version of the blueprint format.
	//
	// Must be the first field in  document. This value determines how the rest
	// of the blueprint is parsed. Currently only version 0 is supported.
	Version int `json:"version"`

	// Services to deploy.
	//
	// Each entry becomes a [plan.Service] after execution. Every service
	// is exposed through the gateway at its declared prefix.
	Services []Service `json:"services"`
}

// Validates the blueprint.
//
// The version must be 0. At least one service is required. Every service
// must have an ID, a reference, and a prefix. Service IDs and prefixes
// must be unique across the blueprint.
func (bp *Blueprint) Validate() error {
	if bp.Version != 0 {
		return crex.Wrap(ErrInvalidBlueprint, ErrUnsupportedVersion)
	}

	if len(bp.Services) == 0 {
		return crex.Wrap(ErrInvalidBlueprint, ErrMissingServices)
	}

	ids := make(map[string]struct{}, len(bp.Services))
	prefixes := make(map[string]struct{}, len(bp.Services))

	for i := range bp.Services {
		if err := bp.Services[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidBlueprint, err)
		}

		if _, exists := ids[bp.Services[i].ID]; exists {
			return crex.Wrap(ErrInvalidBlueprint, ErrDuplicateServiceID)
		}
		ids[bp.Services[i].ID] = struct{}{}

		if _, exists := prefixes[bp.Services[i].Prefix]; exists {
			return crex.Wrap(ErrInvalidBlueprint, ErrDuplicatePrefix)
		}
		prefixes[bp.Services[i].Prefix] = struct{}{}
	}

	return nil
}
