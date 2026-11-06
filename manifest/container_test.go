package manifest

import (
	"errors"
	"testing"
)

func TestContainerValidate(t *testing.T) {
	cases := []struct {
		name string
		ctr  Container
		want error
	}{
		{"empty network", Container{}, nil},
		{
			"valid network",
			Container{Network: Network{Ingress: []IngressRule{{Protocol: "tcp"}}}},
			nil,
		},
		{
			"invalid network",
			Container{Network: Network{Ingress: []IngressRule{{Protocol: "nope"}}}},
			ErrInvalidContainer,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.ctr.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Container.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
