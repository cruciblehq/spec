package compute

import "fmt"

// Error response for a failed operation.
//
// Returned with a non-2xx status and [MediaTypeError] when a unary operation
// cannot complete. Code classifies the failure for programmatic handling and
// Message describes it. A mid-stream failure of a streaming operation is
// reported instead by a terminal [ExecError] message, because the response
// status is already committed once streaming has begun.
type Error struct {
	Code    ErrorCode `codec:"code"`    // Machine-readable classification.
	Message string    `codec:"message"` // Human-readable description.
}

// Implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Validates the error response.
//
// The code must be a known [ErrorCode] and the message must not be empty.
func (e *Error) Validate() error {
	if !isValidErrorCode(e.Code) {
		return ErrErrorCodeInvalid
	}
	if e.Message == "" {
		return ErrErrorMessageEmpty
	}
	return nil
}
