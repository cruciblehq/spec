package state

// Represents the current state of a deployment.
//
// Records what resources have been deployed and their runtime identifiers.
// Used for incremental deployments and resource lifecycle management.
type State struct {
	Version    int        `json:"version"`
	Deployment Deployment `json:"deployment"`
	Services   []Service  `json:"services"`
}
