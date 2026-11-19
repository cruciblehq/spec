package registry

import (
	"regexp"
	"strings"

	"github.com/cruciblehq/spec/reference"
)

// Maximum length of a name, in bytes.
const maxNameLength = 63

var (

	// Valid name pattern.
	namePattern = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)

	// Valid digest pattern (algorithm:hex).
	digestPattern = regexp.MustCompile(`^[a-z0-9]+:[a-f0-9]+$`)
)

// Checks that a name is valid for a namespace, resource, or channel.
//
// Names may include lowercase letters, digits, and hyphens, must start and end
// with an alphanumeric character, and must not exceed 63 characters.
func ValidateName(name string) error {
	if name == "" {
		return ErrNameEmpty
	}
	if len(name) > maxNameLength {
		return ErrNameTooLong
	}
	if !namePattern.MatchString(name) {
		return ErrNameInvalid
	}
	return nil
}

// Checks that a version string is valid.
//
// The version must parse with [reference.ParseVersion].
func ValidateVersionString(version string) error {
	if _, err := reference.ParseVersion(version); err != nil {
		return ErrVersionInvalid
	}
	return nil
}

// Checks that timestamps are valid.
//
// Both createdAt and updatedAt must be positive unix epochs, and updatedAt
// must not precede createdAt.
func ValidateTimestamps(createdAt, updatedAt int64) error {
	if createdAt <= 0 {
		return ErrTimestampInvalid
	}
	if updatedAt <= 0 {
		return ErrTimestampInvalid
	}
	if updatedAt < createdAt {
		return ErrTimestampOrder
	}
	return nil
}

// Checks that a resource type string is valid.
//
// The type must not be empty or only whitespace. It is not checked against a
// specific pattern, to allow flexibility, but it must be set to ensure
// meaningful categorization. Client implementations may choose to enforce
// additional constraints on the type string as needed.
func ValidateResourceType(t string) error {
	if strings.TrimSpace(t) == "" {
		return ErrTypeEmpty
	}
	return nil
}

// Checks that a digest string is valid.
//
// Must follow the algorithm:hex format (e.g., "sha256:abcdef...").
func ValidateDigest(digest string) error {
	if !digestPattern.MatchString(digest) {
		return ErrDigestInvalid
	}
	return nil
}

// Checks that archive fields are valid.
//
// Archive, size, and digest must either all be nil (not uploaded) or all be
// set with valid values: archive non-empty, size positive, digest in
// algorithm:hex format.
func ValidateArchiveFields(archive *string, size *int64, digest *string) error {
	set := 0
	if archive != nil {
		set++
	}
	if size != nil {
		set++
	}
	if digest != nil {
		set++
	}
	if set != 0 && set != 3 {
		return ErrArchiveIncomplete
	}
	if set == 0 {
		return nil
	}
	if *archive == "" {
		return ErrArchiveEmpty
	}
	if *size <= 0 {
		return ErrSizeInvalid
	}
	return ValidateDigest(*digest)
}

// Checks that a count is valid.
//
// Counts must not be negative.
func ValidateCount(n int) error {
	if n < 0 {
		return ErrCountNegative
	}
	return nil
}

// Validates a namespace identifier.
//
// Ensures the namespace name follows naming conventions.
func ValidateNamespace(namespace string) error {
	return ValidateName(namespace)
}

// Validates a resource identifier (namespace + resource).
//
// Ensures both namespace and resource names follow naming conventions.
func ValidateIdentifier(namespace, resource string) error {
	if err := ValidateName(namespace); err != nil {
		return err
	}
	return ValidateName(resource)
}

// Validates a version reference (namespace + resource + version).
//
// Ensures namespace and resource names follow naming conventions and
// the version string is a valid semantic version.
func ValidateReference(namespace, resource, version string) error {
	if err := ValidateName(namespace); err != nil {
		return err
	}
	if err := ValidateName(resource); err != nil {
		return err
	}
	return ValidateVersionString(version)
}

// Validates a channel reference (namespace + resource + channel).
//
// Ensures namespace, resource, and channel names all follow naming conventions.
func ValidateChannelReference(namespace, resource, channel string) error {
	if err := ValidateName(namespace); err != nil {
		return err
	}
	if err := ValidateName(resource); err != nil {
		return err
	}
	return ValidateName(channel)
}

// Validates channel info (namespace + resource + channel name + version).
//
// Ensures namespace, resource, and channel names follow naming conventions,
// and the target version string is a valid semantic version.
func ValidateChannelInfo(namespace, resource string, info ChannelInfo) error {
	if err := ValidateName(namespace); err != nil {
		return err
	}
	if err := ValidateName(resource); err != nil {
		return err
	}
	if err := ValidateName(info.Name); err != nil {
		return err
	}
	return ValidateVersionString(info.Version)
}
