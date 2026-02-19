// Package manifest defines the protocol types for Crucible resource manifests.
//
// A manifest is the top-level document that declares a resource's metadata,
// type, and type-specific configuration. A manifest is represented by the
// [Manifest] struct, which contains a schema version, [Resource] metadata
// (type, name, version), and a type-specific configuration ([Manifest.Config]).
// The config field holds a pointer to one of the concrete types, determined by
// the resource type.
//
// Services and runtimes share a common build pipeline structure called a
// [Recipe]. A recipe consists of one or more [Stage] values, each with its
// own base image and build [Step] values. Stages run in declaration order.
// Artifacts produced by a named stage can be referenced from subsequent stages
// via copy steps (e.g. "builder:/app/bin"). Exactly one stage must be
// non-transient (the output stage), which becomes the final build artifact.
// All other stages must be marked transient.
//
// Build steps within a stage are classified as operations or modifiers:
//
// Operations are the actions that produce side effects in the build container.
// [Step.Run] executes a command through a shell, and [Step.Copy] copies files
// from the host or from another stage into the image. These two are mutually
// exclusive within a single step.
//
// Modifiers adjust the build environment. [Step.Shell] selects the shell for
// run commands, [Step.Env] sets environment variables, [Step.Workdir] sets the
// working directory, and [Step.Platform] restricts the step to a specific
// OS/architecture. Modifiers combine freely with each other. When paired with
// an operation, they apply to that single step. When set alone, they persist
// in the image for subsequent steps.
//
// Some modifier-operation combinations are invalid. [Step.Shell] and [Step.Env]
// cannot be paired with [Step.Copy], since copy operations do not involve shell
// execution or environment variables.
//
// Setting [Step.Platform] together with [Step.Steps] creates a platform group:
// a set of child steps that all execute under the specified platform. Modifiers
// on the group step apply to all children. A platform group cannot also contain
// an operation, and nesting platform groups is not allowed.
//
// Every type in the package exposes a Validate method that checks structural
// correctness. Validation cascades from [Manifest.Validate] down through
// [Resource], the config type, [Recipe], [Stage], and [Step].
//
// Encoding a manifest:
//
//	m := &manifest.Manifest{
//		Resource: manifest.Resource{
//			Type:    manifest.TypeService,
//			Name:    "crucible/hub",
//			Version: "1.0.0",
//		},
//		Config: &manifest.Service{ /* ... */ },
//	}
//	data, err := manifest.Encode(m)
//
// Decoding a manifest:
//
//	m, err := manifest.Decode(data)
//	if err != nil {
//		log.Fatal(err) // malformed YAML or validation failure
//	}
//	switch cfg := m.Config.(type) {
//	case *manifest.Service:
//		fmt.Println("service", cfg)
//	case *manifest.Runtime:
//		fmt.Println("runtime", cfg)
//	}
package manifest
