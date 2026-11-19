package agent

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrErrorCodeInvalid    = crex.New("error code must be a known value")
	ErrErrorMessageEmpty   = crex.New("error message must not be empty")
	ErrNotFound            = crex.New("not found")
	ErrAlreadyStarted      = crex.New("already started")
	ErrNotStarted          = crex.New("not started")
	ErrExecMessageTooLarge = crex.New("exec message too large")
)
