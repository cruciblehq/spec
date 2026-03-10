package manifest

import (
	"strconv"

	"github.com/cruciblehq/crex"
)

// A build stage in a recipe.
//
// Each stage is an independent build pipeline with its own source image
// and steps. Named stages can be referenced from subsequent stages
// (e.g. "builder:/app/bin" in a copy step). Stages are non-transient by
// default, meaning their image is exported as the final build artifact.
// Set [Stage.Transient] to true for intermediate stages. When [Stage.Platform]
// is set, the stage only runs for the matching target platform. Steps within a
// platform-scoped stage cannot use [Step.Platform] since the platform is
// already fixed for the entire stage.
type Stage struct {

	// Identifies the stage for cross-stage references.
	//
	// When set, must be unique across all stages in the recipe. Used as
	// the prefix in copy source paths (e.g. "builder:/path"). Stages that
	// do not need to be referenced by other stages can omit the name.
	Name string `yaml:"name,omitempty"`

	// Marks this stage as an intermediate build helper.
	//
	// Transient stages are not exported as the final build artifact. They
	// exist only to produce artifacts that are copied into later stages.
	// In a single-stage recipe this field can be omitted (defaults to
	// false). In a multi-stage recipe every stage except the output stage
	// must be marked transient.
	Transient bool `yaml:"transient,omitempty"`

	// Restricts this stage to a specific target platform.
	//
	// When set, the stage is only built when the target platform matches.
	// The format is "os/arch" (e.g. "linux/arm64"). Steps within a
	// platform-scoped stage cannot use [Step.Platform]; the platform is
	// fixed for the entire stage.
	Platform string `yaml:"platform,omitempty"`

	// Specifies the base image source for this stage.
	//
	// A "file" prefix selects a local OCI image archive on disk, such as
	// "file ./alpine.tar". An "oci" prefix pulls a remote container image
	// from a registry; the reference is a single colon-delimited token
	// like "oci alpine:3.21" or "oci docker.io/library/alpine:3.21".
	// Anything else is treated as a Crucible runtime reference where the
	// name and version constraint are separated by a space, such as
	// "crucible/runtime 0.1.0". An explicit "ref" prefix may be used for
	// clarity or to disambiguate a resource named "file" or "oci".
	From string `yaml:"from"`

	// Ordered build steps for this stage.
	Steps []Step `yaml:"steps"`
}

// Validates the stage.
//
// The base image source is parsed and validated. Each step is validated
// recursively with positional context.
func (s *Stage) Validate() error {
	if s.Name != "" {
		if _, err := strconv.Atoi(s.Name); err == nil {
			return crex.Wrap(ErrInvalidStage, ErrNumericStageName)
		}
	}

	src, err := s.ParseFrom()
	if err != nil {
		return crex.Wrap(ErrInvalidStage, err)
	}

	if err := src.Validate(); err != nil {
		return crex.Wrap(ErrInvalidStage, err)
	}

	for i := range s.Steps {
		if s.Platform != "" && stepUsesPlatform(&s.Steps[i]) {
			return crex.Wrapf(ErrInvalidStage, "step %d: %w", i+1, ErrPlatformInPlatformStage)
		}
		if err := s.Steps[i].Validate(); err != nil {
			return crex.Wrapf(ErrInvalidStage, "step %d: %w", i+1, err)
		}
	}

	return nil
}

// Reports whether a step or any of its children use the platform field.
func stepUsesPlatform(s *Step) bool {
	if s.Platform != "" {
		return true
	}
	for i := range s.Steps {
		if stepUsesPlatform(&s.Steps[i]) {
			return true
		}
	}
	return false
}
