package blueprint

// A service instance within a blueprint.
//
// Multiple instances of the same underlying service can appear in a single
// blueprint as long as they have distinct IDs and non-overlapping prefixes.
type Service struct {

	// Stable identifier for this service instance.
	//
	// The ID is carried through plan generation and into the deployment
	// state, so it should not change between blueprint revisions unless
	// the intent is to replace the service.
	ID string `json:"id"`

	// Crucible resource reference for the service.
	//
	// Follows the standard reference format: "namespace/name constraint",
	// for example "cruciblehq/hub ^1.0.0". The reference is resolved
	// against the registry during [Blueprint.Execute].
	Reference string `json:"reference"`

	// HTTP path prefix for the service in the gateway.
	//
	// All of the service's endpoints are exposed under this prefix.
	// Prefixes must not overlap — "/api/hub" and "/api/hub/users" would
	// conflict because one nests inside the other.
	Prefix string `json:"prefix"`
}

// Validates a service entry.
func (s *Service) validate() error {
	if s.ID == "" {
		return ErrMissingServiceID
	}
	if s.Reference == "" {
		return ErrMissingReference
	}
	if s.Prefix == "" {
		return ErrMissingPrefix
	}
	return nil
}
