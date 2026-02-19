package manifest

import "errors"

var (

	// General errors
	ErrInvalidResourceType = errors.New("invalid resource type")
	ErrInvalidManifest     = errors.New("invalid manifest")
	ErrInvalidResource     = errors.New("invalid resource")
	ErrInvalidRecipe       = errors.New("invalid recipe")
	ErrInvalidStage        = errors.New("invalid stage")
	ErrInvalidSource       = errors.New("invalid source")
	ErrInvalidStep         = errors.New("invalid step")
	ErrInvalidService      = errors.New("invalid service")
	ErrInvalidWidget       = errors.New("invalid widget")

	// Codec
	ErrEncodeFailed = errors.New("failed to encode manifest")
	ErrDecodeFailed = errors.New("failed to decode manifest")

	// Manifest
	ErrUnsupportedVersion = errors.New("unsupported manifest version")
	ErrMissingConfig      = errors.New("missing config")
	ErrConfigTypeMismatch = errors.New("config type does not match resource type")

	// Resource
	ErrMissingName    = errors.New("missing resource name")
	ErrMissingVersion = errors.New("missing resource version")

	// Recipe
	ErrMissingStages        = errors.New("recipe must have at least one stage")
	ErrMultipleOutputStages = errors.New("recipe must have exactly one non-transient stage")
	ErrNoOutputStage        = errors.New("recipe must have at least one non-transient stage")
	ErrDuplicateStageName   = errors.New("duplicate stage name")

	// Stage
	ErrMissingFrom       = errors.New("stage missing base image")
	ErrInvalidFromFormat = errors.New("invalid from format")
	ErrNumericStageName  = errors.New("stage name must not be numeric")

	// Source
	ErrUnknownSourceType = errors.New("unknown source type")
	ErrMissingValue      = errors.New("missing source value")

	// Step
	ErrMutuallyExclusiveOps  = errors.New("run and copy are mutually exclusive")
	ErrEmptyStep             = errors.New("step has no fields set")
	ErrShellWithCopy         = errors.New("shell cannot be used with copy")
	ErrEnvWithCopy           = errors.New("env cannot be used with copy")
	ErrStepsWithoutPlatform  = errors.New("child steps require platform")
	ErrPlatformWithOperation = errors.New("platform group cannot have operations")
	ErrNestedPlatformGroup   = errors.New("platform groups cannot be nested")

	// Service
	ErrMissingEntrypoint = errors.New("service missing entrypoint")

	// Widget
	ErrMissingMain = errors.New("widget missing main entry point")
)
