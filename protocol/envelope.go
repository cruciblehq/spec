package protocol

// Request/response wrapper for all socket messages.
//
// The Version field is set on outgoing messages and checked on incoming
// ones to detect incompatible clients.
type Envelope struct {
	Version int     `json:"version"`           // Protocol version for compatibility checks.
	Command Command `json:"command"`           // Command name for routing to handlers.
	Payload any     `json:"payload,omitempty"` // Command-specific data.
}

// Checks that required fields are present.
func (e *Envelope) Validate() error {
	if e.Command == "" {
		return ErrMissingCommand
	}
	return nil
}
