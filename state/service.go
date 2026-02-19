package state

// Represents a service that has been deployed.
type Service struct {
	ID         string `json:"id"`          // Stable identifier assigned at composition time.
	Reference  string `json:"reference"`   // Frozen resource reference with exact version and digest.
	ResourceID string `json:"resource_id"` // Runtime identifier assigned during deployment.
}

// Validates that the service has an ID, reference, and resource ID.
func (svc *Service) validate() error {
	if svc.ID == "" {
		return ErrMissingServiceID
	}
	if svc.Reference == "" {
		return ErrMissingReference
	}
	if svc.ResourceID == "" {
		return ErrMissingResourceID
	}
	return nil
}
