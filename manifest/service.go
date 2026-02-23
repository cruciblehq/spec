package manifest

import "github.com/cruciblehq/crex"

// Holds configuration specific to service resources.
//
// Service resources are backend components that provide functionality to other
// systems by exposing an API. They build on top of a base image defined by
// the embedded [Recipe], which specifies the source image and build steps.
type Service struct {
	Recipe `yaml:",inline"`

	// Command to run when the container starts.
	//
	// Sets the entrypoint on the output image produced by the recipe.
	Entrypoint []string `yaml:"entrypoint,omitempty"`
}

// Validates the service configuration.
func (s *Service) Validate() error {
	if len(s.Entrypoint) == 0 {
		return crex.Wrap(ErrInvalidService, ErrMissingEntrypoint)
	}

	return s.Recipe.Validate()
}
