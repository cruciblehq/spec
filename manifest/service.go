package manifest

import (
	"github.com/cruciblehq/spec/reference"
	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/crex"
)

// Defines a service resource manifest.
//
// Service resources are backend components that provide functionality to other
// systems by exposing an API. They build on top of a base image defined by
// the embedded [Recipe], which specifies the source image and build steps.
type Service struct {
	Recipe

	// Declared parameters for this service.
	//
	// Lists configuration values the service accepts at runtime. Values are
	// bound through environment declarations.
	Schema *Schema `codec:"schema,omitempty"`

	// Interfaces asserted by this service.
	//
	// Each item names an interface the service claims to satisfy. The platform
	// treats these as symbolic assertions and does not attempt to validate
	// actual conformance.
	Implements []string `codec:"implements,omitempty"`

	// Interface contracts required by this service.
	//
	// Each item names an interface the service expects to bind. Resolution
	// identifies one or more implementations of the matching interface
	// identity and version.
	Requires []string `codec:"requires,omitempty"`

	// Interface contracts attached to an existing binding.
	//
	// Each item names an interface the service expects to bind to a logical
	// dependency established elsewhere through a [Service.Requires] entry.
	// The reference syntax matches [Service.Requires] exactly.
	Attaches []string `codec:"attaches,omitempty"`

	// Command to run when the container starts.
	//
	// Sets the entrypoint on the output image produced by the recipe. The
	// entrypoint is required for service resources since they are expected
	// to run as containers.
	Entrypoint []string `codec:"entrypoint,omitempty"`
}

// Encodes the service configuration to a format-independent map.
//
// The schema, entrypoint, and recipe stages are encoded in the canonical
// serialized format.
func (s *Service) Encode(c *codec.Codec) (any, error) {
	m := make(map[string]any)
	stages, err := s.Recipe.encodeStages(c)
	if err != nil {
		return nil, err
	}
	if stages != nil {
		m["stages"] = stages
	}
	if s.Schema != nil {
		schema, err := c.ToMap(s.Schema)
		if err != nil {
			return nil, err
		}
		if len(schema) > 0 {
			m["schema"] = schema
		}
	}
	if len(s.Entrypoint) > 0 {
		m["entrypoint"] = s.Entrypoint
	}
	if len(s.Implements) > 0 {
		m["implements"] = s.Implements
	}
	if len(s.Requires) > 0 {
		m["requires"] = s.Requires
	}
	if len(s.Attaches) > 0 {
		m["attaches"] = s.Attaches
	}
	return m, nil
}

// Validates the service configuration.
func (s *Service) Validate() error {
	if len(s.Entrypoint) == 0 {
		return crex.Wrap(ErrInvalidService, ErrMissingEntrypoint)
	}

	if err := validateInterfaceReferences(s.Implements, true); err != nil {
		return crex.Wrap(ErrInvalidService, err)
	}
	if err := validateInterfaceReferences(s.Requires, false); err != nil {
		return crex.Wrap(ErrInvalidService, err)
	}
	if err := validateInterfaceReferences(s.Attaches, false); err != nil {
		return crex.Wrap(ErrInvalidService, err)
	}

	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return crex.Wrap(ErrInvalidService, err)
		}
	}

	return s.Recipe.Validate()
}

// Validates each interface reference used by the service.
func validateInterfaceReferences(items []string, exact bool) error {
	for _, item := range items {
		ref, err := reference.Parse(item, string(TypeInterface))
		if err != nil {
			return crex.Wrap(ErrInvalidRef, err)
		}
		if ref.Version() == nil {
			return crex.Wrap(ErrInvalidRef, ErrMissingRefTarget)
		}
		if exact && !ref.Version().IsExact() {
			return crex.Wrap(ErrInvalidRef, ErrMissingRefTarget)
		}
	}
	return nil
}
