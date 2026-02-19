package state

import "github.com/cruciblehq/crex"

// Current state format version.
const Version = 0

// Represents the current state of a deployment.
//
// Records what resources have been deployed and their runtime identifiers.
// Used for incremental deployments and resource lifecycle management.
type State struct {
	Version    int        `json:"version"`    // Schema version of the state format.
	Deployment Deployment `json:"deployment"` // Metadata about the most recent deployment.
	Services   []Service  `json:"services"`   // Services that were deployed.
}

// Validates the state.
//
// The version must match [Version]. The deployment timestamp must be set.
// Every service must have an ID, a reference, and a resource ID.
func (s *State) Validate() error {
	if s.Version != Version {
		return crex.Wrap(ErrInvalidState, ErrUnsupportedVersion)
	}

	if s.Deployment.DeployedAt.IsZero() {
		return crex.Wrap(ErrInvalidState, ErrMissingDeployedAt)
	}

	for i := range s.Services {
		if err := s.Services[i].validate(); err != nil {
			return crex.Wrap(ErrInvalidState, err)
		}
	}

	return nil
}
