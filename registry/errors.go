package registry

import "github.com/cruciblehq/utils-go/crex"

var (

	// Validation errors.

	ErrNameEmpty         = crex.New("name cannot be empty")
	ErrNameTooLong       = crex.New("name cannot exceed 63 characters")
	ErrNameInvalid       = crex.New("name must contain only lowercase letters, numbers, and hyphens, and must start and end with an alphanumeric character")
	ErrVersionInvalid    = crex.New("version format must be semantic version")
	ErrTimestampInvalid  = crex.New("timestamp must be a positive unix epoch")
	ErrTimestampOrder    = crex.New("updatedAt must not be before createdAt")
	ErrTypeEmpty         = crex.New("resource type cannot be empty")
	ErrArchiveEmpty      = crex.New("archive URL cannot be empty when set")
	ErrSizeInvalid       = crex.New("archive size must be positive")
	ErrDigestInvalid     = crex.New("digest must be in algorithm:hex format")
	ErrArchiveIncomplete = crex.New("archive, size, and digest must all be set or all be null")
	ErrCountNegative     = crex.New("count must not be negative")
	ErrErrorCodeInvalid  = crex.New("error code must be a known value")
	ErrErrorMessageEmpty = crex.New("error message cannot be empty")

	// Type validation errors.

	ErrInvalidNamespace = crex.New("invalid namespace")
	ErrInvalidResource  = crex.New("invalid resource")
	ErrInvalidVersion   = crex.New("invalid version")
	ErrInvalidChannel   = crex.New("invalid channel")
)
