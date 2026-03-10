package manifest

import (
	"fmt"

	"github.com/cruciblehq/crex"
)

// The OCI image artifact produced by recipe-based builds (runtimes and services).
const ImageFile = "image.tar"

// The disk image artifact produced by machine builds.
const MachineImageFile = "disk.qcow2"

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
	outputPlatforms := make(map[string]bool)
	unplatformedOutputs := 0

	for i := range r.Stages {
		if err := r.validateStage(i, seen, outputPlatforms, &unplatformedOutputs); err != nil {
			return err
		}
	}

	if unplatformedOutputs == 0 && len(outputPlatforms) == 0 {
		return crex.Wrap(ErrInvalidRecipe, ErrNoOutputStage)
	}
	if unplatformedOutputs > 1 {
		return crex.Wrap(ErrInvalidRecipe, ErrMultipleOutputStages)
	}

	return nil
}

// Validates a single stage within the recipe, checking for duplicate names
// and tracking output stage counts.
func (r *Recipe) validateStage(i int, seen map[string]bool, outputPlatforms map[string]bool, unplatformedOutputs *int) error {
	stage := &r.Stages[i]
	name := stage.Name

	if name != "" {
		if seen[name] {
			return crex.Wrapf(ErrInvalidRecipe, "%w: %s", ErrDuplicateStageName, name)
		}
		seen[name] = true
	}

	if err := stage.Validate(); err != nil {
		return crex.Wrapf(ErrInvalidRecipe, "stage %s: %w", stageLabel(name, i), err)
	}

	if !stage.Transient {
		if err := r.trackOutputStage(stage.Platform, outputPlatforms, unplatformedOutputs); err != nil {
			return err
		}
	}

	return nil
}

// Tracks output stage platform uniqueness.
func (r *Recipe) trackOutputStage(platform string, outputPlatforms map[string]bool, unplatformedOutputs *int) error {
	if platform == "" {
		*unplatformedOutputs++
		return nil
	}
	if outputPlatforms[platform] {
		return crex.Wrapf(ErrInvalidRecipe, "%w: %s", ErrDuplicateOutputPlatform, platform)
	}
	outputPlatforms[platform] = true
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
