package seccomp

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrInvalidGrant   = crex.New("invalid seccomp grant")
	ErrUnknownSyscall = crex.New("unknown syscall")
)
