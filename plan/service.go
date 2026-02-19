package plan

// Represents a service in the deployment plan.
//
// Contains the resolved reference with exact version and digest.
type Service struct {
	ID        string `json:"id"`        // Stable identifier for this service instance.
	Reference string `json:"reference"` // Resolved resource reference with exact version and digest.
}

// Validates that the service has an ID and reference.
func (s *Service) validate() error {
	if s.ID == "" {
		return ErrMissingServiceID
	}
	if s.Reference == "" {
		return ErrMissingReference
	}
	return nil
}
