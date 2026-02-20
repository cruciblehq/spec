package protocol

// Payload for the container-stop command.
type ContainerStopRequest struct {
	ID string `json:"id"` // Container identifier.
}

// Checks that all required fields are present.
func (r *ContainerStopRequest) Validate() error {
	if r.ID == "" {
		return ErrMissingID
	}
	return nil
}

// Payload for the container-destroy command.
type ContainerDestroyRequest struct {
	ID string `json:"id"` // Container identifier.
}

// Checks that all required fields are present.
func (r *ContainerDestroyRequest) Validate() error {
	if r.ID == "" {
		return ErrMissingID
	}
	return nil
}

// Payload for the container-status command.
type ContainerStatusRequest struct {
	ID string `json:"id"` // Container identifier.
}

// Checks that all required fields are present.
func (r *ContainerStatusRequest) Validate() error {
	if r.ID == "" {
		return ErrMissingID
	}
	return nil
}

// Payload for the container-exec command.
//
// The daemon executes the command inside the container and returns
// captured stdout, stderr, and the exit code.
type ContainerExecRequest struct {
	ID      string   `json:"id"`      // Container identifier.
	Command []string `json:"command"` // Command and arguments to execute.
}

// Checks that all required fields are present.
func (r *ContainerExecRequest) Validate() error {
	if r.ID == "" {
		return ErrMissingID
	}
	if len(r.Command) == 0 {
		return ErrMissingExecCommand
	}
	return nil
}

// Payload for the container-update command.
//
// The daemon stops the container, re-imports the image from the new tarball,
// and restarts the container with the same identifier.
type ContainerUpdateRequest struct {
	Ref     string `json:"ref"`     // Resource path.
	Version string `json:"version"` // Image version.
	ID      string `json:"id"`      // Container identifier.
	Path    string `json:"path"`    // Path to the new OCI image tarball.
}

// Checks that all required fields are present.
func (r *ContainerUpdateRequest) Validate() error {
	if r.Ref == "" {
		return ErrMissingRef
	}
	if r.Version == "" {
		return ErrMissingVersion
	}
	if r.ID == "" {
		return ErrMissingID
	}
	if r.Path == "" {
		return ErrMissingPath
	}
	return nil
}
