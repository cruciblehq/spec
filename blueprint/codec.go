package blueprint

import (
	"encoding/json"

	"github.com/cruciblehq/crex"
)

// Encodes a blueprint to JSON.
//
// The blueprint is validated before marshaling. Returns [ErrEncodeFailed]
// if validation or marshaling fails.
func Encode(bp *Blueprint) ([]byte, error) {
	if err := bp.Validate(); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}

	data, err := json.Marshal(bp)
	if err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	return data, nil
}

// Decodes a JSON document into a blueprint.
//
// The blueprint is validated after unmarshaling. Returns [ErrDecodeFailed]
// if unmarshaling or validation fails.
func Decode(data []byte) (*Blueprint, error) {
	var bp Blueprint
	if err := json.Unmarshal(data, &bp); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	if err := bp.Validate(); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	return &bp, nil
}