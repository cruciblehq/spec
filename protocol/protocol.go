package protocol

// The current protocol version.
const Version = 0

// Identifies the operation in a protocol envelope.
type Command string

const (
	CmdBuild    Command = "build"    // Execute a recipe-based build.
	CmdStatus   Command = "status"   // Query daemon status.
	CmdShutdown Command = "shutdown" // Shut down the daemon.
	CmdOK       Command = "ok"       // Operation completed successfully.
	CmdError    Command = "error"    // Operation failed.
)
