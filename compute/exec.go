package compute

// Requests execution of a command in a running container.
type ExecRequest struct {
	ID      ContainerID `codec:"id"`      // Target container handle.
	Command string      `codec:"command"` // Command passed to the shell.
	Options Options     `codec:"options"` // Runtime options for the exec.
}

// The type of an [ExecResponse] message.
//
// Messages of type [ExecOutput] carry process output. A single terminal message
// of type [ExecExit] or [ExecError] ends the response.
type ExecResponseType string

// Exec response message types.
const (
	ExecOutput ExecResponseType = "output" // Carries a chunk of process output.
	ExecExit   ExecResponseType = "exit"   // Terminal message reporting the exit code.
	ExecError  ExecResponseType = "error"  // Terminal message reporting a failure.
)

// A single message in the streamed response to an [ExecRequest], encoded as one
// newline-delimited JSON object.
//
// Output messages carry a chunk of a named [Channel]; the stream ends with
// exactly one terminal message: [ExecExit] with the process exit code, or
// [ExecError] with an [Error] when the command could not complete.
type ExecResponse struct {
	Type    ExecResponseType `codec:"type"`              // What the message carries.
	Channel Channel          `codec:"channel,omitempty"` // Source channel, for [ExecOutput].
	Data    []byte           `codec:"data,omitempty"`    // Output bytes, for [ExecOutput].
	Code    int              `codec:"code,omitempty"`    // Exit code, for [ExecExit].
	Error   *Error           `codec:"error,omitempty"`   // Failure, for [ExecError].
}
