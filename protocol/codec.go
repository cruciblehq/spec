package protocol

import (
	"encoding/json"

	"github.com/cruciblehq/crex"
)

// Encodes a command and payload into a JSON envelope.
func Encode(cmd Command, payload any) ([]byte, error) {
	env := Envelope{
		Version: Version,
		Command: cmd,
		Payload: payload,
	}
	return json.Marshal(env)
}

// Decodes a JSON envelope, returning the raw payload for typed decoding.
func Decode(data []byte) (*Envelope, json.RawMessage, error) {
	var raw struct {
		Version int             `json:"version"`
		Command Command         `json:"command"`
		Payload json.RawMessage `json:"payload,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, nil, crex.Wrap(ErrMalformedMessage, err)
	}

	if raw.Version != Version {
		return nil, nil, ErrUnsupportedVersion
	}

	env := &Envelope{
		Version: raw.Version,
		Command: raw.Command,
	}

	return env, raw.Payload, nil
}

// Decodes the payload bytes into the target type.
func DecodePayload[T any](payload json.RawMessage) (*T, error) {
	if len(payload) == 0 {
		return nil, nil
	}
	var t T
	if err := json.Unmarshal(payload, &t); err != nil {
		return nil, crex.Wrap(ErrMalformedMessage, err)
	}
	return &t, nil
}
