package registry

import "fmt"

// Error response from the registry API.
//
// Classifies the error through the Code field, which programs can branch on,
// and describes it through the Message field. The media type is
// [MediaTypeError].
type Error struct {
	Code    ErrorCode `codec:"code"`    // Error code (see [ErrorCode]).
	Message string    `codec:"message"` // Error description.
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
