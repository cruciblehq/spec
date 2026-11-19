package manifest

import (
	"github.com/cruciblehq/spec/reference"
	"github.com/cruciblehq/utils-go/crex"
)

// Common identity envelope shared by every resource manifest.
//
// Decoding reads Type first and uses it to determine which concrete struct the
// rest of the document maps to. Keeping these fields in a shared type means
// any program can inspect a manifest for its type and metadata without knowing
// the specific resource types in advance.
type Resource struct {

	// Discriminator that controls how the rest of the manifest is decoded.
	//
	// Must be one of the [ResourceType] values. It is read before any other
	// field and selects the config type the rest of the manifest decodes into.
	Type ResourceType `codec:"type"`

	// Identifier for the resource.
	//
	// Uses the format "namespace/name" (e.g. "cruciblehq/my-api"). The
	// registry host is intentionally absent so the same manifest can be
	// published to multiple registries without modification.
	Name string `codec:"name"`

	// Summary of the resource's purpose, for people.
	//
	// Not interpreted by Crucible programs. Present for documentation and
	// registry engines that surface resource metadata.
	Description string `codec:"description,omitempty"`

	// Semantic version declared by this resource.
	//
	// Set by the resource author at publish time and immutable thereafter.
	// The registry files the artifact under this version so consumers can
	// reference an exact, stable artifact.
	Version string `codec:"version"`
}

// Validates the resource metadata.
//
// The type must be known, the name must parse as "namespace/name" with a
// namespace, and the version must be set.
func (r *Resource) Validate() error {
	if _, err := ParseResourceType(string(r.Type)); err != nil {
		return crex.Wrap(ErrInvalidResource, err)
	}
	if r.Name == "" {
		return crex.Wrap(ErrInvalidResource, ErrMissingName)
	}
	id, err := reference.ParseIdentifier(r.Name, string(r.Type))
	if err != nil {
		return crex.Wrap(ErrInvalidResource, err)
	}
	if id.Namespace() == "" {
		return crex.Wrap(ErrInvalidResource, ErrMissingNamespace)
	}
	if r.Version == "" {
		return crex.Wrap(ErrInvalidResource, ErrMissingVersion)
	}
	return nil
}
