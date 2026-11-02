package compute

// Requests creation of a container from an image.
type LoadRequest struct {
	Ref     ImageRef `codec:"ref"`     // Target image reference.
	Options Options  `codec:"options"` // Runtime options applied at start.
}

// Returns the handle of the loaded container.
type LoadResponse struct {
	ID ContainerID `codec:"id"` // Server-side handle for the container.
}

// Requests export of an image as an OCI archive.
//
// The response body is the archive itself ([MediaTypeImageArchive]), so there
// is no response message.
type ExportRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Carries the reference of a newly imported image.
//
// Import's request body is the OCI archive ([MediaTypeImageArchive]); this is
// the operation's only message.
type ImportResponse struct {
	Ref ImageRef `codec:"ref"` // Reference of the imported image.
}

// Requests the configuration of an image.
type InspectImageRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Carries the configuration of an image.
type InspectImageResponse struct {
	Config ImageConfig `codec:"config"` // The image's configuration.
}

// Requests removal of an image.
//
// The response is empty.
type RemoveRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Requests extraction of a path from an image as a tar stream.
//
// The response body is the tar stream ([MediaTypeFileTree]), so there is no
// response message.
type ExtractRequest struct {
	Ref  ImageRef `codec:"ref"`  // Target image reference.
	Path string   `codec:"path"` // Path within the merged image view.
}

// Requests the start of a container's main process.
//
// The response is empty.
type StartRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Requests the stop of a container's main process.
//
// The response is empty.
type StopRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Requests release of a container.
//
// The response is empty.
type DestroyRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Requests application of a tar archive to a container.
//
// The archive travels as the request body ([MediaTypeLayer]); this message
// carries the target. The response is empty.
type CopyRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Requests the configuration of a container.
type InspectContainerRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Carries the configuration of a container.
type InspectContainerResponse struct {
	Config ImageConfig `codec:"config"` // The container's configuration.
}

// Requests a commit of a container's current state.
type CommitRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
	By string      `codec:"by"` // Creator recorded in the image history.
}

// Carries the reference of the image produced by a commit.
type CommitResponse struct {
	Ref ImageRef `codec:"ref"` // Reference of the image produced by the commit.
}

// Sets the pending configuration on a container.
//
// The response is empty.
type ConfigureRequest struct {
	ID     ContainerID `codec:"id"`     // Target container handle.
	Config ImageConfig `codec:"config"` // Configuration applied at commit.
}
