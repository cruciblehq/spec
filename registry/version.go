package registry

import "github.com/cruciblehq/utils-go/crex"

// Properties of a version for creation or update.
//
// Used as the request body for version creation and update operations. For
// update requests, the version field must match the URL path parameter or
// context. Setting Public makes a private version public, which cannot be
// undone. The media type is [MediaTypeVersionInfo].
type VersionInfo struct {
	String string `codec:"string"`           // Version string (e.g., "1.0.0").
	Public bool   `codec:"public,omitempty"` // Whether the version is public.
}

// Validates the version info.
//
// The version string must be a valid semantic version (see
// [ValidateVersionString]).
func (info *VersionInfo) Validate() error {
	if err := ValidateVersionString(info.String); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	return nil
}

// Lightweight version representation for listings.
//
// Provides version metadata without full archive details. Used in resource
// listings and version lists to keep payloads compact. Includes read-only
// fields like timestamps.
type VersionSummary struct {
	String    string `codec:"string"`    // Version string (e.g., "1.0.0").
	Public    bool   `codec:"public"`    // Whether the version is public.
	CreatedAt int64  `codec:"createdAt"` // When the version was created.
	UpdatedAt int64  `codec:"updatedAt"` // When the version was last updated.
}

// Validates the version summary.
func (s *VersionSummary) Validate() error {
	if err := ValidateVersionString(s.String); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	if err := ValidateTimestamps(s.CreatedAt, s.UpdatedAt); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	return nil
}

// Complete version with archive details.
//
// The archive, size, and digest fields are null before archive upload and
// populated afterward. A private version can have its archive replaced for
// iterative development. A public version is final, so a dependency that
// names it always gets the same archive. Includes scoping information to
// identify the version's location. The media type is [MediaTypeVersion].
type Version struct {
	Namespace string  `codec:"namespace"` // Namespace this version belongs to.
	Resource  string  `codec:"resource"`  // Resource this version belongs to.
	String    string  `codec:"string"`    // Version string (e.g., "1.0.0").
	Public    bool    `codec:"public"`    // Whether the version is public.
	Archive   *string `codec:"archive"`   // Download URL or null if not uploaded.
	Size      *int64  `codec:"size"`      // Archive size in bytes (null if not uploaded).
	Digest    *string `codec:"digest"`    // Archive digest (e.g., "sha256:abc...", null if not uploaded).
	CreatedAt int64   `codec:"createdAt"` // When the version was created.
	UpdatedAt int64   `codec:"updatedAt"` // When the version was last updated.
}

// Validates the version.
func (v *Version) Validate() error {
	if err := ValidateName(v.Namespace); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	if err := ValidateName(v.Resource); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	if err := ValidateVersionString(v.String); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	if err := ValidateArchiveFields(v.Archive, v.Size, v.Digest); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	if err := ValidateTimestamps(v.CreatedAt, v.UpdatedAt); err != nil {
		return crex.Wrap(ErrInvalidVersion, err)
	}
	return nil
}

// Collection of versions for a resource.
//
// The media type is [MediaTypeVersionList].
type VersionList struct {
	Versions []VersionSummary `codec:"versions"` // List of versions.
}

// Validates the version list.
func (l *VersionList) Validate() error {
	for i := range l.Versions {
		if err := l.Versions[i].Validate(); err != nil {
			return crex.Wrap(ErrInvalidVersion, err)
		}
	}
	return nil
}
