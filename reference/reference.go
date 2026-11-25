package reference

import (
	"strings"

	"github.com/cruciblehq/utils-go/crex"
)

// Resource reference.
//
// A reference encapsulates all information needed to locate, identify, and
// verify a Crucible resource. It combines an [Identifier] with version
// information. References are immutable once created. Use [Parse] to
// construct valid references.
type Reference struct {
	Identifier                    // Resource identifier (registry, namespace, name).
	version    *VersionConstraint // Semantic version constraint.
	digest     *Digest            // Content digest for pinning to an exact version.
}

// Parses a reference string.
//
// The context type is required, and used to set the type when the reference
// string does not include one, or to validate the type when it does. When
// the reference string includes a type, it must match the context type. The
// expected string format is [<type>] [[[registry/]namespace/]name]
// <version-constraint> [<digest>]. The type is optional and must be lowercase
// alphabetic. When omitted, the context type is used. When present, it must
// match the context type exactly. The location is a single token with up to
// three slash-separated segments (registry/namespace/name). When segments are
// omitted, the corresponding fields are empty. Callers apply defaults where
// needed. A version constraint is required and may span multiple tokens
// (e.g., ">=1.0.0 <2.0.0"). The digest is optional and follows the format
// algorithm:hash (e.g., "sha256:abcd1234"). When present, it freezes the
// reference to a specific content version.
func Parse(s, contextType string) (*Reference, error) {
	p := &referenceParser{
		tokens: strings.Fields(s),
	}
	return p.parse(contextType)
}

// Creates a reference from an identifier, version constraint, and optional
// digest.
//
// Useful for programmatically building references when you already have the
// parsed components, avoiding the overhead of formatting and re-parsing a
// reference string. The version parameter is a version constraint string
// (e.g., "1.2.3", ">=1.0.0 <2.0.0"). The digest parameter is optional. When
// provided, the reference becomes frozen, pointing to an exact immutable
// resource version.
func New(id *Identifier, version string, digest *Digest) (*Reference, error) {
	if id == nil {
		return nil, crex.Wrap(ErrInvalidReference, ErrEmptyReference)
	}

	vc, err := ParseVersionConstraint(version)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidReference, err)
	}

	return &Reference{
		Identifier: *id,
		version:    vc,
		digest:     digest,
	}, nil
}

// Returns a copy of this reference with defaults applied for any empty fields.
//
// If the registry is empty and defaultRegistry is non-empty, the registry is
// set. If the namespace is empty and defaultNamespace is non-empty, the
// namespace is set. Fields that are already populated are never overwritten.
func (r *Reference) WithDefaults(defaultRegistry, defaultNamespace string) *Reference {
	clone := *r
	clone.Identifier = *clone.Identifier.WithDefaults(defaultRegistry, defaultNamespace)
	return &clone
}

// Semantic version constraint.
func (r *Reference) Version() *VersionConstraint {
	return r.version
}

// Cryptographic hash for content verification. Nil if not frozen.
func (r *Reference) Digest() *Digest {
	return r.digest
}

// Whether the reference includes a digest.
//
// A frozen reference refers to an exact, immutable resource version.
func (r *Reference) IsFrozen() bool {
	return r.digest != nil
}

// Returns a string representation of the reference.
//
// Includes only the fields that are set. Registry and namespace appear only
// when present on the underlying identifier. The version constraint is included
// when set, and digest is appended if present. This is not necessarily a
// canonical form that parses back to the same reference. A reference parsed
// without applying defaults may omit the registry and namespace.
func (r *Reference) String() string {
	if r == nil {
		return ""
	}
	var sb strings.Builder

	sb.WriteString(r.Identifier.String())

	if r.version != nil {
		sb.WriteByte(' ')
		sb.WriteString(r.version.String())
	}

	if r.IsFrozen() {
		sb.WriteByte(' ')
		sb.WriteString(r.digest.String())
	}

	return sb.String()
}
