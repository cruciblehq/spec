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
// Set [Stage.Transient] to true for intermediate stages.
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

	// Specifies the base image source for this stage.
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
		if err := s.Steps[i].Validate(); err != nil {
			return crex.Wrapf(ErrInvalidStage, "step %d: %w", i+1, err)
		}
	}

	return nil
}
