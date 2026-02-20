package registry

import (
	"encoding/json"

	"github.com/cruciblehq/crex"
)

// Encodes a registry type to JSON.
//
// If the value implements a Validate method, validation is performed before
// encoding. Returns [ErrEncodeFailed] on failure.
func Encode(v any) ([]byte, error) {
	if val, ok := v.(interface{ Validate() error }); ok {
		if err := val.Validate(); err != nil {
			return nil, crex.Wrap(ErrEncodeFailed, err)
		}
	}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	return data, nil
}

// Decodes JSON into a registry type.
//
// If the decoded value implements a Validate method, validation is performed
// after decoding. Returns [ErrDecodeFailed] on failure.
func Decode[T any](data []byte) (*T, error) {
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	if val, ok := any(&v).(interface{ Validate() error }); ok {
		if err := val.Validate(); err != nil {
			return nil, crex.Wrap(ErrDecodeFailed, err)
		}
	}

	return &v, nil
}
