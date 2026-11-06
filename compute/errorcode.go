package compute

// Machine-readable classification of a failed operation.
//
// Carried in the Code field of an [Error] so crux can branch on the kind of
// failure beyond the HTTP status.
type ErrorCode string

const (
	ErrorCodeBadRequest         ErrorCode = "bad_request"         // Malformed or invalid request.
	ErrorCodeUnauthenticated    ErrorCode = "unauthenticated"     // Missing or rejected credential.
	ErrorCodeNotFound           ErrorCode = "not_found"           // Image or container does not exist.
	ErrorCodeConflict           ErrorCode = "conflict"            // Operation conflicts with current state (e.g. container already started).
	ErrorCodePreconditionFailed ErrorCode = "precondition_failed" // Operation invalid in the current state (e.g. container not started).
	ErrorCodeInternalError      ErrorCode = "internal_error"      // Unexpected server-side failure.
)

// Set of known error codes for validation.
var validErrorCodes = map[ErrorCode]bool{
	ErrorCodeBadRequest:         true,
	ErrorCodeUnauthenticated:    true,
	ErrorCodeNotFound:           true,
	ErrorCodeConflict:           true,
	ErrorCodePreconditionFailed: true,
	ErrorCodeInternalError:      true,
}

// Whether the error code is a known value.
func isValidErrorCode(code ErrorCode) bool {
	return validErrorCodes[code]
}
