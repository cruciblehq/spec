package mac

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrCompile      = crex.New("MAC compile failed")
	ErrInvalidSpec  = crex.New("invalid MAC spec")
	ErrInvalidAllow = crex.New("invalid MAC allow rule")
	ErrInvalidExpr  = crex.New("invalid MAC expression")
	ErrInvalidValue = crex.New("invalid MAC value")
)
