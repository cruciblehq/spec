package plan

// Represents a service in the deployment plan.
//
// Contains the resolved reference with exact version and digest.
type Service struct {
	ID        string `json:"id"`
	Reference string `json:"reference"`
}
