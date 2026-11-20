package manifest

import (
	"io"

	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/crex"
)

// Returns the manifest config as T.
//
// Returns [ErrConfigTypeMismatch] when the config is not a T.
func As[T any](m *Manifest) (T, error) {
	cfg, ok := m.Config.(T)
	if !ok {
		var zero T
		return zero, crex.Newf(ErrConfigTypeMismatch, "expected manifest type %T", *new(T))
	}
	return cfg, nil
}

// Decodes a YAML manifest from r.
//
// Returns [ErrDecodeFailed] when r cannot be read or does not hold a manifest.
func Decode(r io.Reader) (*Manifest, error) {
	var m Manifest
	if err := decode(r, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Encodes m as YAML to w.
//
// Returns [ErrEncodeFailed] when m cannot be encoded. Errors from w are
// returned as they are.
func Encode(w io.Writer, m *Manifest) error {
	return encode(w, m)
}

// Decodes a YAML plan from r.
//
// Returns [ErrDecodeFailed] when r cannot be read or does not hold a plan.
func DecodePlan(r io.Reader) (*Plan, error) {
	var p Plan
	if err := decode(r, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Encodes p as YAML to w.
//
// Returns [ErrEncodeFailed] when p cannot be encoded. Errors from w are
// returned as they are.
func EncodePlan(w io.Writer, p *Plan) error {
	return encode(w, p)
}

// Reads all of r and decodes it as YAML into v.
func decode(r io.Reader, v any) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return crex.Wrap(ErrDecodeFailed, err)
	}
	if err := codec.Decode(data, v, codec.YAML); err != nil {
		return crex.Wrap(ErrDecodeFailed, err)
	}
	return nil
}

// Encodes v as YAML and writes it to w.
func encode(w io.Writer, v any) error {
	data, err := codec.Encode(v, codec.YAML)
	if err != nil {
		return crex.Wrap(ErrEncodeFailed, err)
	}
	_, err = w.Write(data)
	return err
}
