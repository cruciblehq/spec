package manifest

// Represents the type of a Crucible resource.
type ResourceType string

const (

	// Runtime resource type.
	TypeRuntime ResourceType = "runtime"

	// Service resource type.
	TypeService ResourceType = "service"

	// Template resource type.
	TypeTemplate ResourceType = "template"

	// Widget resource type.
	TypeWidget ResourceType = "widget"
)

// Converts a string to a resource type, returning an error if invalid.
func ParseResourceType(s string) (ResourceType, error) {
	switch ResourceType(s) {
	case TypeRuntime, TypeService, TypeTemplate, TypeWidget:
		return ResourceType(s), nil
	default:
		return "", ErrInvalidResourceType
	}
}
