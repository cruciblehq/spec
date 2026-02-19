package state

// Represents a service that has been deployed.
type Service struct {
	ID         string `json:"id"`          // Service identifier, assigned at composition time.
	Reference  string `json:"reference"`   // Frozen service reference with exact version and digest.
	ResourceID string `json:"resource_id"` // Runtime resource identifier assigned during deployment.
}
