package manifest

import "github.com/cruciblehq/crex"

// Defines a Crucible resource.
//
// A manifest specifies metadata about the resource and its type-specific
// configuration. The [Manifest.Config] field is polymorphic, its type being
// determined by [Resource.Type]. Each resource has its own config type.
type Manifest struct {

	// Schema version of the manifest format.
	//
	// Determines how the rest of the manifest is interpreted. Currently
	// the only supported version is 0.
	Version int `yaml:"version"`

	// Common metadata shared across all resource types.
	//
	// Includes the resource type, qualified name, and version. This is
	// required and must be valid for the manifest to be considered valid.
	Resource Resource `yaml:"resource"`

	// Type-specific configuration.
	//
	// The concrete type depends on [Resource.Type]: [Runtime] from runtimes,
	// [Service] for services, [Widget] for widgets, etc.
	Config any `yaml:"-"`
}

// Validates the manifest.
//
// The version must be 0. Resource metadata must be valid. Config must be
// present and match the resource type. The config is validated according
// to its concrete type.
func (m *Manifest) Validate() error {
	if m.Version != 0 {
		return crex.Wrap(ErrInvalidManifest, ErrUnsupportedVersion)
	}

	if err := m.Resource.Validate(); err != nil {
		return crex.Wrap(ErrInvalidManifest, err)
	}

	if m.Config == nil {
		return crex.Wrap(ErrInvalidManifest, ErrMissingConfig)
	}

	if err := m.validateConfig(); err != nil {
		return crex.Wrap(ErrInvalidManifest, err)
	}

	return nil
}

// Validates that Config matches the resource type and is internally valid.
func (m *Manifest) validateConfig() error {
	switch m.Resource.Type {
	case TypeRuntime:
		cfg, ok := m.Config.(*Runtime)
		if !ok {
			return ErrConfigTypeMismatch
		}
		return cfg.Validate()

	case TypeService:
		cfg, ok := m.Config.(*Service)
		if !ok {
			return ErrConfigTypeMismatch
		}
		return cfg.Validate()

	case TypeWidget:
		cfg, ok := m.Config.(*Widget)
		if !ok {
			return ErrConfigTypeMismatch
		}
		return cfg.Validate()

	case TypeTemplate:
		cfg, ok := m.Config.(*Template)
		if !ok {
			return ErrConfigTypeMismatch
		}
		return cfg.Validate()

	default:
		return ErrInvalidResourceType
	}
}
