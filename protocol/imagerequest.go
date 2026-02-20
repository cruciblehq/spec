package protocol

// Payload for the image-import command.
//
// The daemon imports the OCI archive at Path into the containerd image
// store, tags it as Ref:Version, and unpacks it for the host platform.
type ImageImportRequest struct {
	Ref     string `json:"ref"`     // Resource path (e.g., "my-namespace/my-service").
	Version string `json:"version"` // Image version (e.g., "1.0.0").
	Path    string `json:"path"`    // Path to the OCI image tarball.
}

// Checks that all required fields are present.
func (r *ImageImportRequest) Validate() error {
	if r.Ref == "" {
		return ErrMissingRef
	}
	if r.Version == "" {
		return ErrMissingVersion
	}
	if r.Path == "" {
		return ErrMissingPath
	}
	return nil
}

// Payload for the image-start command.
//
// The daemon creates a container from the image identified by Ref:Version
// and starts a long-running task so exec can attach later. If ID is empty,
// the daemon defaults to the resource name.
type ImageStartRequest struct {
	Ref     string `json:"ref"`          // Resource path.
	Version string `json:"version"`      // Image version.
	ID      string `json:"id,omitempty"` // Container identifier.
}

// Checks that all required fields are present.
func (r *ImageStartRequest) Validate() error {
	if r.Ref == "" {
		return ErrMissingRef
	}
	if r.Version == "" {
		return ErrMissingVersion
	}
	return nil
}

// Payload for the image-destroy command.
//
// The daemon removes the image identified by Ref:Version and all containers
// created from it.
type ImageDestroyRequest struct {
	Ref     string `json:"ref"`     // Resource path.
	Version string `json:"version"` // Image version.
}

// Checks that all required fields are present.
func (r *ImageDestroyRequest) Validate() error {
	if r.Ref == "" {
		return ErrMissingRef
	}
	if r.Version == "" {
		return ErrMissingVersion
	}
	return nil
}
