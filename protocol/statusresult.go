package protocol

// Returned for the status command.
type StatusResult struct {
	Running bool   `json:"running"`          // Always true when the daemon is reachable.
	Version string `json:"version"`          // Daemon version string.
	Pid     int    `json:"pid"`              // Daemon process ID.
	Uptime  string `json:"uptime,omitempty"` // Uptime since start.
	Builds  int    `json:"builds"`           // Total builds processed in this session.
}
