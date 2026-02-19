package plan

import (
	"encoding/json"

	"github.com/cruciblehq/crex"
)

// Encodes a plan to JSON.
//
// The plan is validated before marshaling. Returns [ErrEncodeFailed]
// if validation or marshaling fails.
func Encode(p *Plan) ([]byte, error) {
	if err := p.Validate(); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}

	data, err := json.Marshal(p)
	if err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	return data, nil
}

// Decodes a JSON document into a plan.
//
// The plan is validated after unmarshaling. Returns [ErrDecodeFailed]
// if unmarshaling or validation fails.
func Decode(data []byte) (*Plan, error) {
	var p Plan
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	if err := p.Validate(); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	return &p, nil
}
