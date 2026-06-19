package net

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrInvalidGrant       = crex.New("invalid net grant")
	ErrInvalidSpec        = crex.New("invalid network spec")
	ErrInvalidIngressRule = crex.New("invalid network ingress rule")
	ErrInvalidEgressRule  = crex.New("invalid network egress rule")
)
