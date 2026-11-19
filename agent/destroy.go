package agent

// Requests release of a container.
type DestroyRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Confirms the release of a container.
type DestroyResponse struct {
}
