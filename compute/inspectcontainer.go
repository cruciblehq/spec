package compute

// Requests the configuration of a container.
type InspectContainerRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Carries the configuration of a container.
type InspectContainerResponse struct {
	Config ImageConfig `codec:"config"` // The container's configuration.
}
