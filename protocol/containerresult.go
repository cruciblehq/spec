package protocol

// State of a container managed by cruxd.
type ContainerState string

const (
	ContainerRunning    ContainerState = "running"     // Task is active.
	ContainerStopped    ContainerState = "stopped"     // Container exists but has no running task.
	ContainerNotCreated ContainerState = "not created" // Container does not exist.
)

// Returned for the container-status command.
type ContainerStatusResult struct {
	Status ContainerState `json:"status"` // Container state.
}

// Returned for the container-exec command.
type ContainerExecResult struct {
	ExitCode int    `json:"exitCode"`         // Process exit code.
	Stdout   string `json:"stdout,omitempty"` // Captured standard output.
	Stderr   string `json:"stderr,omitempty"` // Captured standard error.
}
