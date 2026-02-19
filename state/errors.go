package state

import "errors"

var (
	ErrInvalidState   = errors.New("invalid state")
	ErrEncodeFailed   = errors.New("failed to encode state")
	ErrDecodeFailed   = errors.New("failed to decode state")

	ErrUnsupportedVersion = errors.New("unsupported state version")
	ErrMissingDeployedAt  = errors.New("missing deployment timestamp")
	ErrMissingServiceID   = errors.New("service missing id")
	ErrMissingReference   = errors.New("service missing reference")
	ErrMissingResourceID  = errors.New("service missing resource id")
)
