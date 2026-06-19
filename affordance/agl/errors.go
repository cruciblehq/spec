package agl

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrParse = crex.New("agl parse failed")
	ErrLex   = crex.New("agl lex failed")
)
