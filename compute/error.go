package compute

import "fmt"

// Result of a failed operation.
//
// A failed operation returns a response with [MediaTypeError]. Code classifies
// the failure and Message describes it.
type Error struct {
	Code    ErrorCode `codec:"code"`    // Error code.
	Message string    `codec:"message"` // Error message.
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
