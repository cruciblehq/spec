package agent

import "net/http"

// Classification of a failed operation.
//
// Carried in the Code field of an [Error] so the client can branch on the kind
// of failure beyond the HTTP status.
type ErrorCode string

// Known error codes.
const (
	ErrorCodeBadRequest         ErrorCode = "bad_request"         // Malformed or invalid request.
	ErrorCodeNotFound           ErrorCode = "not_found"           // Image or container does not exist.
	ErrorCodeConflict           ErrorCode = "conflict"            // Operation conflicts with current state (e.g. container already started).
	ErrorCodePreconditionFailed ErrorCode = "precondition_failed" // Operation invalid in the current state (e.g. container not started).
	ErrorCodeInternalError      ErrorCode = "internal_error"      // Unexpected failure in the agent.
)

// Set of known error codes for validation.
var validErrorCodes = map[ErrorCode]bool{
	ErrorCodeBadRequest:         true,
	ErrorCodeNotFound:           true,
	ErrorCodeConflict:           true,
	ErrorCodePreconditionFailed: true,
	ErrorCodeInternalError:      true,
}

// Whether the error code is a known value.
func isValidErrorCode(code ErrorCode) bool {
	return validErrorCodes[code]
}

// Returns the HTTP status of a response that carries an [Error] with the code.
//
// Unknown codes map to 500.
func (c ErrorCode) HTTPStatus() int {
	switch c {
	case ErrorCodeBadRequest:
		return http.StatusBadRequest
	case ErrorCodeNotFound:
		return http.StatusNotFound
	case ErrorCodeConflict:
		return http.StatusConflict
	case ErrorCodePreconditionFailed:
		return http.StatusPreconditionFailed
	default:
		return http.StatusInternalServerError
	}
}
