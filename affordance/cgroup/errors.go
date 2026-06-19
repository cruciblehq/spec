package cgroup

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrInvalidGrant = crex.New("invalid cgroup grant")
	ErrConflict     = crex.New("cgroup conflict")
	ErrUnknownKnob  = crex.New("unknown cgroup knob")
)
