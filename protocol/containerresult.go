package protocol

// Returned for the container-status command.
type ContainerStatusResult struct {
	Status string `json:"status"` // Container state ("running", "stopped", "not created").
}

// Returned for the container-exec command.
type ContainerExecResult struct {
	ExitCode int    `json:"exitCode"`         // Process exit code.
	Stdout   string `json:"stdout,omitempty"` // Captured standard output.
	Stderr   string `json:"stderr,omitempty"` // Captured standard error.
}
