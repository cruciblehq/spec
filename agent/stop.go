package agent

// Requests the stop of a container's main process.
type StopRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Confirms the stop of a container's main process.
type StopResponse struct {
}
