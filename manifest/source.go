package manifest

import (
	"strings"

	"github.com/cruciblehq/crex"
)

// Discriminates the origin of a stage's base image.
type SourceType string

const (

	// Local OCI image archive.
	SourceFile SourceType = "file"

	// Crucible runtime reference.
	SourceRef SourceType = "ref"

	// Remote OCI image reference pulled from a container registry.
	SourceOCI SourceType = "oci"
)

// Describes the origin of a stage's base image.
//
// Type discriminates between a local OCI tarball, a Crucible runtime
// reference, and a remote OCI image reference. Value holds the raw
// payload after the type prefix.
type Source struct {

	// Discriminates between file, ref, and oci sources.
	Type SourceType

	// Source payload after the type prefix.
	//
	// For file sources this is the local OCI image archive path relative
	// to the manifest file. For ref sources this is the Crucible runtime
	// reference string (name and version separated by a space). For oci
	// sources this is a single-token container image reference such as
	// "alpine:3.21" or "docker.io/library/alpine:3.21".
	Value string
}

// Validates the source.
//
// The type must be a known source type and the value must not be empty.
func (s *Source) Validate() error {
	switch s.Type {
	case SourceFile, SourceRef, SourceOCI:
	default:
		return crex.Wrap(ErrInvalidSource, ErrUnknownSourceType)
	}

	if s.Value == "" {
		return crex.Wrap(ErrInvalidSource, ErrMissingValue)
	}

	return nil
}

// Parses the stage's from field into a source.
//
// The string is tokenized on whitespace. A "file" prefix selects a local
// OCI archive. An "oci" prefix selects a remote container image reference,
// which must be a single token (e.g., "oci alpine:3.21"). Everything else
// is treated as a Crucible runtime reference, where name and version are
// separated by a space (e.g., "crucible/runtime 0.1.0"); the optional
// "ref" prefix is stripped when present. A runtime literally named "file"
// or "oci" must use the "ref" prefix to avoid ambiguity.
func (s *Stage) ParseFrom() (Source, error) {
	fields := strings.Fields(s.From)
	if len(fields) == 0 {
		return Source{}, ErrInvalidFromFormat
	}

	typ := SourceType(fields[0])
	switch typ {
	case SourceOCI:
		if len(fields) != 2 {
			return Source{}, ErrInvalidFromFormat
		}
		return Source{Type: SourceOCI, Value: fields[1]}, nil

	case SourceFile, SourceRef:
		if len(fields) < 2 {
			return Source{}, ErrInvalidFromFormat
		}
		return Source{Type: typ, Value: strings.Join(fields[1:], " ")}, nil

	default:
		return Source{Type: SourceRef, Value: strings.Join(fields, " ")}, nil
	}
}
