package manifest

import (
	"errors"
	"testing"
)

func TestDeploymentValidate(t *testing.T) {
	valid := Deployment{
		Service:   "svc",
		Container: "ctr",
		Compute:   "cmp",
		Network:   "net",
	}
	cases := []struct {
		name string
		dep  Deployment
		want error
	}{
		{"valid", valid, nil},
		{"with environment", Deployment{Service: "s", Container: "c", Compute: "m", Network: "n", Environment: "e"}, nil},
		{"missing service", Deployment{Container: "c", Compute: "m", Network: "n"}, ErrMissingDeploymentService},
		{"missing container", Deployment{Service: "s", Compute: "m", Network: "n"}, ErrMissingDeploymentContainer},
		{"missing compute", Deployment{Service: "s", Container: "c", Network: "n"}, ErrMissingDeploymentCompute},
		{"missing network", Deployment{Service: "s", Container: "c", Compute: "m"}, ErrMissingDeploymentNetwork},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.dep.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Deployment.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
