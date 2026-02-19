package protocol

import (
	"github.com/cruciblehq/crex"
	"github.com/cruciblehq/spec/manifest"
)

// Payload for the build command.
//
// Carries a fully-resolved recipe and metadata needed for the daemon to
// execute a build. All stage sources must be resolved to file paths before
// sending; the daemon does not resolve references.
type BuildRequest struct {
	Recipe     *manifest.Recipe `json:"recipe"`               // Parsed recipe to execute.
	Resource   string           `json:"resource"`             // Resource name, used as a prefix for container IDs.
	Output     string           `json:"output"`               // Directory for the final build artifact.
	Root       string           `json:"root"`                 // Project root, for resolving copy sources.
	Entrypoint []string         `json:"entrypoint,omitempty"` // OCI entrypoint to set on the output image.
	Platforms  []string         `json:"platforms,omitempty"`  // Target platforms (e.g., ["linux/amd64"]). Defaults to host.
}

// Checks that all required build fields are present, validates the recipe,
// and verifies that all stage sources are resolved to file paths.
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

	if err := r.Recipe.Validate(); err != nil {
		return err
	}

	for i := range r.Recipe.Stages {
		src, err := r.Recipe.Stages[i].ParseFrom()
		if err != nil {
			return crex.Wrapf(ErrInvalidBuildRequest, "stage %d: %w", i+1, err)
		}
		if src.Type != manifest.SourceFile {
			return crex.Wrapf(ErrInvalidBuildRequest, "stage %d: %w", i+1, ErrUnresolvedSource)
		}
	}

	return nil
}
