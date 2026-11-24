package manifest

import (
	"errors"
	"testing"
)

func TestStateValidate(t *testing.T) {
	cases := []struct {
		name  string
		state State
		want  error
	}{
		{"empty valid", State{Version: StateVersion}, nil},
		{"bad version", State{Version: StateVersion + 1}, ErrInvalidState},
		{
			"valid records",
			State{
				Version:   StateVersion,
				Instances: []InstanceRecord{{ID: "i", Image: "img", Provider: "local", Compute: "c"}},
				Volumes:   []VolumeRecord{{ID: "v", Service: "s"}},
			},
			nil,
		},
		{
			"instance missing field",
			State{Version: StateVersion, Instances: []InstanceRecord{{ID: "i"}}},
			ErrMissingRequiredField,
		},
		{
			"volume missing field",
			State{Version: StateVersion, Volumes: []VolumeRecord{{ID: "v"}}},
			ErrMissingRequiredField,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.state.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("State.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
