package agent

// Requests the start of a container's main process.
type StartRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Confirms the start of a container's main process.
type StartResponse struct {
}
