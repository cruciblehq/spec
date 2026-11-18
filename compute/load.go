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
