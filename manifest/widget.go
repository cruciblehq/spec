package manifest

// Holds configuration specific to widget resources.
//
// Widget resources are frontend components that can be embedded into apps.
// This structure defines configurations that are unique to widget resource
// manifests, such as build settings and requested affordances. It is used as
// [Manifest.Config] when the resource type is [ResourceType.Widget].
type Widget struct {

	// Build entry point.
	Main string `yaml:"main"`
}

// Validates the widget configuration.
//
// The main entry point is required.
func (w *Widget) Validate() error {
	if w.Main == "" {
		return wrap(ErrInvalidWidget, ErrMissingMain)
	}
	return nil
}
