package manifest

import (
	"github.com/cruciblehq/utils-go/crex"
)

// Identifies a resource with a persistent and parameterized reference string.
//
// A reference names the resource in one of two ways, never both. [Ref.Ref]
// is used to look up and fetch the resource from the registry. [Ref.Path]
// points to a local directory that holds the resource's manifest. [Ref.ID]
// assigns a stable identifier to the reference when it is used to name a
// specific instance of a resource (e.g. a service in a blueprint). The ID
// persists from composition through plan into state, allowing the instance
// to be tracked and referenced across the deployment lifecycle. The reference
// may carry [Ref.Args] to parameterize the referenced resource.
type Ref struct {

	// Stable identifier for this instance.
	//
	// Set when the reference names a specific instance (e.g. a service in a
	// blueprint). The ID persists from composition through plan into state.
	ID string `codec:"id,omitempty"`

	// Crucible reference string identifying the resource.
	//
	// The format is "namespace/name version", where namespace/name is the
	// qualified resource name and version is a version constraint. The
	// registry resolves this string to a specific artifact. This field is
	// required unless Path is set.
	Ref string `codec:"ref,omitempty"`

	// Local directory that holds the resource's manifest.
	//
	// A relative path is resolved against the directory the tool runs in.
	// This field is required unless Ref is set.
	Path string `codec:"path,omitempty"`

	// Named arguments for the referenced resource.
	//
	// Maps parameter names to their string values. Each key must match a
	// parameter declared in the resource's schema.
	Args Args `codec:"args,omitempty"`
}

// Validates the ref.
func (r *Ref) Validate() error {
	if r.ID != "" && !isValidName(r.ID) {
		return crex.Wrap(ErrInvalidRef, ErrInvalidRefID)
	}
	if r.Ref == "" && r.Path == "" {
		return crex.Wrap(ErrInvalidRef, ErrMissingRefTarget)
	}
	if r.Ref != "" && r.Path != "" {
		return crex.Wrap(ErrInvalidRef, ErrRefAndPath)
	}
	if err := r.Args.Validate(); err != nil {
		return crex.Wrap(ErrInvalidRef, err)
	}
	return nil
}
