package protocol

// Returned on error.
type ErrorResult struct {
	Message string `json:"message"` // Error description.
}

// Checks that the error carries a message.
func (r *ErrorResult) Validate() error {
	if r.Message == "" {
		return ErrMissingMessage
	}
	return nil
}
