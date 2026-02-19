package protocol

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
