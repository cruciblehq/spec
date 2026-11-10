package compute

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrUnauthenticated   = crex.New("unauthenticated")
	ErrErrorCodeInvalid  = crex.New("error code must be a known value")
	ErrErrorMessageEmpty = crex.New("error message must not be empty")
	ErrNotFound          = crex.New("not found")
	ErrAlreadyStarted    = crex.New("already started")
	ErrNotStarted        = crex.New("not started")
)
