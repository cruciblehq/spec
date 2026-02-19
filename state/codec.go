package state

import (
	"encoding/json"

	"github.com/cruciblehq/crex"
)

// Encodes a state to JSON.
//
// The state is validated before marshaling. Returns [ErrEncodeFailed]
// if validation or marshaling fails.
func Encode(s *State) ([]byte, error) {
	if err := s.Validate(); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	return data, nil
}

// Decodes a JSON document into a state.
//
// The state is validated after unmarshaling. Returns [ErrDecodeFailed]
// if unmarshaling or validation fails.
func Decode(data []byte) (*State, error) {
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	if err := s.Validate(); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	return &s, nil
}
