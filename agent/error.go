package agent

import "fmt"

// Error response for a failed operation.
//
// Returned with [MediaTypeError] and the status given by [ErrorCode.HTTPStatus]
// when a unary operation cannot complete. Code classifies the failure and
// Message describes it. A failure of a streaming operation during the stream
// is reported instead by a terminal [ExecError] message, because the response
// status is already sent once streaming has begun.
type Error struct {
	Code    ErrorCode `codec:"code"`    // Classification that programs can branch on.
	Message string    `codec:"message"` // Description for people.
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
