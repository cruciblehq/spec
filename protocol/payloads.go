package protocol

import "github.com/cruciblehq/spec/manifest"

// Payload for the build command.
//
// Carries the recipe and metadata needed for the daemon to execute a
// recipe-based build. crux reads the manifest, extracts the recipe, and
// sends it here. The daemon never sees or parses the manifest itself.
type BuildRequest struct {
	Recipe     *manifest.Recipe `json:"recipe"`               // Parsed recipe to execute.
	Resource   string           `json:"resource"`             // Resource name, used as a prefix for container IDs.
	Output     string           `json:"output"`               // Directory for the final build artifact.
	Root       string           `json:"root"`                 // Project root, for resolving copy sources.
	Entrypoint []string         `json:"entrypoint,omitempty"` // OCI entrypoint to set on the output image.
}

// Checks that all required build fields are present and that the recipe
// itself is valid.
func (r *BuildRequest) Validate() error {
	if r.Recipe == nil {
		return ErrMissingRecipe
	}
	if r.Resource == "" {
		return ErrMissingResource
	}
	if r.Output == "" {
		return ErrMissingOutput
	}
	if r.Root == "" {
		return ErrMissingRoot
	}
	return r.Recipe.Validate()
}

// Returned after a successful build.
type BuildResult struct {
	Output string `json:"output"` // Directory containing the built artifact.
}

// Checks that the result contains an output path.
func (r *BuildResult) Validate() error {
	if r.Output == "" {
		return ErrMissingOutput
	}
	return nil
}

// Returned for the status command.
type StatusResult struct {
	Running bool   `json:"running"`          // Always true when the daemon is reachable.
	Version string `json:"version"`          // Daemon version string.
	Pid     int    `json:"pid"`              // Daemon process ID.
	Uptime  string `json:"uptime,omitempty"` // Uptime since start.
	Builds  int    `json:"builds"`           // Total builds processed in this session.
}

// Returned on error.
type ErrorResult struct {
	Message string `json:"message"` // Error description.
}

// Checks that the error carries a message.
func (r *ErrorResult) Validate() error {
	if r.Message == "" {
		return ErrMissingMessage
	}
	return nil
}
