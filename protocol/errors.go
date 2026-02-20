package protocol

import "errors"

var (
	ErrUnsupportedVersion  = errors.New("unsupported protocol version")
	ErrUnknownCommand      = errors.New("unknown command")
	ErrMalformedMessage    = errors.New("malformed message")
	ErrInvalidBuildRequest = errors.New("invalid build request")

	ErrMissingCommand   = errors.New("missing command")
	ErrMissingRecipe    = errors.New("missing recipe")
	ErrMissingResource  = errors.New("missing resource name")
	ErrMissingOutput    = errors.New("missing output directory")
	ErrMissingRoot      = errors.New("missing project root")
	ErrMissingMessage   = errors.New("missing error message")
	ErrUnresolvedSource = errors.New("unresolved source")

	ErrMissingRef         = errors.New("missing resource reference")
	ErrMissingVersion     = errors.New("missing version")
	ErrMissingPath        = errors.New("missing path")
	ErrMissingID          = errors.New("missing container identifier")
	ErrMissingExecCommand = errors.New("missing exec command")
)
