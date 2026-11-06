package manifest

import (
	"errors"
	"testing"
)

func TestInfrastructureValidate(t *testing.T) {
	cases := []struct {
		name  string
		infra Infrastructure
		want  error
	}{
		{"empty", Infrastructure{}, nil},
		{
			"valid",
			Infrastructure{
				Computes: map[string]Compute{"c": {Type: ComputeTypeLocal, Config: &ComputeLocal{Host: "h"}}},
				Networks: map[string]Network{"n": {}},
			},
			nil,
		},
		{
			"bad compute",
			Infrastructure{Computes: map[string]Compute{"c": {Type: ComputeTypeLocal}}},
			ErrMissingComputeProvider,
		},
		{
			"bad network",
			Infrastructure{Networks: map[string]Network{"n": {Ingress: []IngressRule{{Protocol: "nope"}}}}},
			ErrInvalidNetwork,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.infra.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Infrastructure.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
