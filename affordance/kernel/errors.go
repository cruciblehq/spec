package kernel

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrInvalidGrant = crex.New("invalid kernel grant")
	ErrInvalidSpec  = crex.New("invalid kernel spec")
)
