package manifest

import (
	"errors"
	"testing"
)

func TestStateValidate(t *testing.T) {
	validDeployment := Deployment{
		Service:   "svc",
		Container: "ctr",
		Compute:   "cmp",
		Network:   "net",
	}
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
				Version:     StateVersion,
				Deployments: []Deployment{validDeployment},
				Instances:   []InstanceRecord{{ID: "i", Image: "img", Provider: "local", Compute: "c"}},
				Volumes:     []VolumeRecord{{ID: "v", Service: "s"}},
				Networks:    []NetworkRecord{{ID: "n", Subnet: "10.0.0.0/8"}},
				Attachments: []AttachmentRecord{{Unit: "u", Network: "n"}},
			},
			nil,
		},
		{
			"invalid deployment",
			State{Version: StateVersion, Deployments: []Deployment{{}}},
			ErrInvalidState,
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
		{
			"network missing field",
			State{Version: StateVersion, Networks: []NetworkRecord{{ID: "n"}}},
			ErrMissingRequiredField,
		},
		{
			"attachment missing field",
			State{Version: StateVersion, Attachments: []AttachmentRecord{{Unit: "u"}}},
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
