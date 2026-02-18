package protocol

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedVersion = errors.New("unsupported protocol version")
	ErrUnknownCommand     = errors.New("unknown command")
	ErrMalformedMessage   = errors.New("malformed message")

	ErrMissingCommand   = errors.New("missing command")
	ErrMissingRecipe    = errors.New("missing recipe")
	ErrMissingResource  = errors.New("missing resource name")
	ErrMissingOutput    = errors.New("missing output directory")
	ErrMissingRoot      = errors.New("missing project root")
	ErrMissingMessage   = errors.New("missing error message")
	ErrUnresolvedSource = errors.New("unresolved source")
)

// Combines a sentinel error with a specific error to provide both broad
// categorization and detailed context.
func wrap(sentinel, err error) error {
	return fmt.Errorf("%w: %w", sentinel, err)
}
