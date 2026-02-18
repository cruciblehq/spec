package manifest

import "strings"

// Discriminates the origin of a stage's base image.
type SourceType string

const (

	// Local OCI image archive.
	SourceFile SourceType = "file"

	// Crucible runtime reference.
	SourceRef SourceType = "ref"
)

// Describes the origin of a stage's base image.
//
// Type discriminates between a local OCI tarball and a Crucible runtime
// reference. Value holds the raw payload after the type prefix.
type Source struct {

	// Discriminates between file and ref sources.
	Type SourceType

	// Source payload after the type prefix.
	//
	// For file sources this is the local OCI image archive path relative to
	// the manifest file. For ref sources this is the Crucible runtime
	// reference string.
	Value string
}

// Validates the source.
//
// The type must be a known source type and the value must not be empty.
func (s *Source) Validate() error {
	switch s.Type {
	case SourceFile, SourceRef:
	default:
		return wrap(ErrInvalidSource, ErrUnknownSourceType)
	}

	if s.Value == "" {
		return wrap(ErrInvalidSource, ErrMissingValue)
	}

	return nil
}

// Parses the stage's from field into a source.
//
// The string is tokenized on whitespace. A "file" prefix selects a local OCI
// archive. Everything else is treated as a Crucible runtime reference, with
// an optional "ref" prefix stripped first. A runtime literally named "file"
// must use the "ref" prefix to avoid ambiguity.
func (s *Stage) ParseFrom() (Source, error) {
	fields := strings.Fields(s.From)
	if len(fields) == 0 {
		return Source{}, ErrInvalidFromFormat
	}

	typ := SourceType(fields[0])
	switch typ {
	case SourceFile, SourceRef:
		if len(fields) < 2 {
			return Source{}, ErrInvalidFromFormat
		}
		return Source{Type: typ, Value: strings.Join(fields[1:], " ")}, nil

	default:
		return Source{Type: SourceRef, Value: strings.Join(fields, " ")}, nil
	}
}
