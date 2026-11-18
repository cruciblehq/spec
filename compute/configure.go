package compute

// Sets the pending configuration on a container.
type ConfigureRequest struct {
	ID     ContainerID `codec:"id"`     // Target container handle.
	Config ImageConfig `codec:"config"` // Configuration applied at commit.
}

// Confirms the pending configuration of a container.
type ConfigureResponse struct {
}
