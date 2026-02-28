package manifest

import (
	"fmt"

	"github.com/cruciblehq/crex"
)

// The OCI image artifact produced by recipe-based builds (runtimes, services,
// and machines).
const ImageFile = "image.tar"

// Describes a build pipeline as one or more stages.
//
// A recipe is the reusable unit shared by resource types that produce OCI
// images. It contains a list of stages, each with its own source image
// and build steps. Exactly one stage must be non-transient, which is the
// stage that is exported as the final build artifact.
type Recipe struct {

	// Build stages.
	//
	// Each stage is an independent build pipeline with its own source image
	// and steps. Stages run in declaration order. Artifacts produced by a
	// stage can be referenced from subsequent stages via the stage name in
	// a copy source (e.g. "builder:/app/bin").
	Stages []Stage `yaml:"stages"`
}

// Validates the recipe.
//
// Checks that at least one stage exists, that named stages have unique names,
// and that exactly one stage is non-transient (the output stage). Each stage
// is validated recursively.
func (r *Recipe) Validate() error {
	if len(r.Stages) == 0 {
		return crex.Wrap(ErrInvalidRecipe, ErrMissingStages)
	}

	seen := make(map[string]bool, len(r.Stages))
	outputStages := 0

	for i := range r.Stages {
		name := r.Stages[i].Name
		if name != "" {
			if seen[name] {
				return crex.Wrapf(ErrInvalidRecipe, "%w: %s", ErrDuplicateStageName, name)
			}
			seen[name] = true
		}

		if err := r.Stages[i].Validate(); err != nil {
			return crex.Wrapf(ErrInvalidRecipe, "stage %s: %w", stageLabel(name, i), err)
		}

		if !r.Stages[i].Transient {
			outputStages++
		}
	}

	if outputStages == 0 {
		return crex.Wrap(ErrInvalidRecipe, ErrNoOutputStage)
	}
	if outputStages > 1 {
		return crex.Wrap(ErrInvalidRecipe, ErrMultipleOutputStages)
	}

	return nil
}

// Returns a label for a stage, preferring the name when available and
// falling back to the 1-based index.
func stageLabel(name string, index int) string {
	if name != "" {
		return fmt.Sprintf("%q", name)
	}
	return fmt.Sprintf("%d", index+1)
}
