package plan

import "errors"

var (
	ErrInvalidPlan    = errors.New("invalid plan")
	ErrEncodeFailed   = errors.New("failed to encode plan")
	ErrDecodeFailed   = errors.New("failed to decode plan")

	ErrUnsupportedVersion = errors.New("unsupported plan version")
	ErrMissingServices    = errors.New("plan must have at least one service")
	ErrMissingCompute     = errors.New("plan must have at least one compute")
	ErrMissingServiceID   = errors.New("service missing id")
	ErrMissingReference   = errors.New("service missing reference")
	ErrMissingComputeID   = errors.New("compute missing id")
	ErrMissingProvider    = errors.New("compute missing provider")
	ErrMissingBindSvc     = errors.New("binding missing service")
	ErrMissingBindCompute = errors.New("binding missing compute")
	ErrMissingPattern     = errors.New("route missing pattern")
	ErrMissingRouteSvc    = errors.New("route missing service")
)
