package protocol

// The current protocol version.
const Version = 0

// Identifies the operation in a protocol envelope.
type Command string

const (
	CmdBuild    Command = "build"    // Execute a recipe-based build.
	CmdStatus   Command = "status"   // Query daemon status.
	CmdShutdown Command = "shutdown" // Shut down the daemon.

	CmdImageImport      Command = "image-import"      // Import an OCI image tarball.
	CmdImageStart       Command = "image-start"       // Start a container from an image.
	CmdImageDestroy     Command = "image-destroy"     // Remove an image and its containers.
	CmdContainerStop    Command = "container-stop"    // Stop a container's task.
	CmdContainerDestroy Command = "container-destroy" // Remove a container.
	CmdContainerStatus  Command = "container-status"  // Query container state.
	CmdContainerExec    Command = "container-exec"    // Execute a command in a container.
	CmdContainerUpdate  Command = "container-update"  // Update a container's image and restart.

	CmdOK    Command = "ok"    // Operation completed successfully.
	CmdError Command = "error" // Operation failed.
)
