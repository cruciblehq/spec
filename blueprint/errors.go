package blueprint

import "errors"

var (
	ErrInvalidBlueprint = errors.New("invalid blueprint")
	ErrEncodeFailed     = errors.New("failed to encode blueprint")
	ErrDecodeFailed     = errors.New("failed to decode blueprint")

	ErrUnsupportedVersion = errors.New("unsupported blueprint version")
	ErrMissingServices    = errors.New("blueprint must have at least one service")
	ErrMissingServiceID   = errors.New("service missing id")
	ErrMissingReference   = errors.New("service missing reference")
	ErrMissingPrefix      = errors.New("service missing prefix")
	ErrDuplicateServiceID = errors.New("duplicate service id")
	ErrDuplicatePrefix    = errors.New("duplicate service prefix")
)
