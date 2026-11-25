// Package reference defines the structure and parsing of Crucible resource
// references.
//
// A resource reference is a symbolic string that encapsulates the information
// needed to fetch and verify a resource. Its general shape is an optional
// resource type, a location, a version constraint, and an optional digest,
// written for example as "widget official/my-widget ^1.2.0
// sha256:3a7bd3e2360a3d80c1...". The location is one token of up to three
// segments, registry/namespace/name. The namespace groups related resources
// and the name identifies one within that namespace. The registry may be
// omitted, and the namespace may be omitted only when the registry is also
// omitted. Omitted segments stay empty after parsing, and callers fill them
// with [Reference.WithDefaults]. As a result, the references "my-widget",
// "official/my-widget", and "registry.crucible.net/official/my-widget" all
// denote the same resource when the defaults match.
//
// Versioning follows semantic versioning with a small number of deliberate
// deviations. The version segment accepts the operators that semver libraries
// usually support: coercion of partial versions, OR with || and implicit AND
// by whitespace, the comparison operators =, !=, >, <, >=, and <=, hyphen
// ranges, x-style wildcards, tilde patch ranges, and caret minor ranges.
// Crucible diverges from semver by requiring every version range to have an
// explicit upper bound, since future major versions introduce breaking changes
// that cannot be anticipated; ">=2.0.0 <3.0.0" and ">1.5.0 <2.0.0" are valid,
// while ">=2.0.0" and ">1.5.0" are not. Operators that imply an upper bound
// such as ^1.2.3, ~1.2.3, 1.2.x, and 1.2.3 - 2.0.0 are allowed because the
// bound is unambiguous. Crucible also rejects the bare asterisk wildcard
// because it undermines the stability guarantees that versioning is meant to
// provide.
//
// Pre-releases receive special treatment. Standard semver compares pre-release
// identifiers lexically in ASCII order, which makes "BETA" sort below "alpha"
// and produces ordering that is unhelpful for Crucible's usage. More
// importantly, pre-releases are inherently unstable and unsuitable for
// Crucible's dynamic composition model, so ranges and other operators never
// match them. A pre-release is selected only by its exact version, for example
// "my-widget 1.2.3-alpha.1", which matches that pre-release and nothing else.
//
// When a resource is prepared for deployment its dependencies are resolved to
// concrete versions and the references are frozen by appending a digest. The
// original constraints are preserved alongside the digest for auditing, but
// the digest is what downstream tooling uses to fetch and verify content. A
// digest is a cryptographic hash of the resource content, written for example
// as "sha256:3a7bd3e2360a3d80c1...", and turns the reference into a pointer
// by content that always denotes the same bytes regardless of any changes to
// the symbolic components.
//
// References also carry a resource type. The type is usually not written in
// the reference string; it is supplied by the context when a reference is
// parsed in a position where a particular type is expected. When the expected
// type is ambiguous, callers may write it explicitly as the leading token,
// for example "widget my-widget ^1.2.0 sha256:3a7bd3e2360a3d80c1...". It must
// then match the type the context expects.
//
// Parsing a reference and inspecting its components:
//
//	ref, err := reference.Parse("official/my-widget ^1.2.0", "widget")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(ref.Path(), ref.Version())
//	fmt.Println(ref.String())
package reference
