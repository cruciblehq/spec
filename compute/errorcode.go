package compute

// Error code.
//
// This is included in the Code field of an [Error] to classify the failure.
// These are operation outcomes, independent of the HTTP status, which reflects
// only the transport.
type ErrorCode string

const (
	ErrorCodeImageNotFound           ErrorCode = "image_not_found"           // The target image does not exist.
	ErrorCodeContainerNotFound       ErrorCode = "container_not_found"       // The target container does not exist.
	ErrorCodeContainerAlreadyStarted ErrorCode = "container_already_started" // Start on a container whose main process is already running.
	ErrorCodeContainerNotStarted     ErrorCode = "container_not_started"     // Exec on a container whose main process is not running.
)

// Set of known error codes for validation.
var validErrorCodes = map[ErrorCode]bool{
	ErrorCodeImageNotFound:           true,
	ErrorCodeContainerNotFound:       true,
	ErrorCodeContainerAlreadyStarted: true,
	ErrorCodeContainerNotStarted:     true,
}

// Whether the error code is a known value.
func isValidErrorCode(code ErrorCode) bool {
	return validErrorCodes[code]
}
