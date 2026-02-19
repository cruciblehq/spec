package state

import "time"

// Represents deployment metadata.
type Deployment struct {
	DeployedAt time.Time `json:"deployed_at"`
}
