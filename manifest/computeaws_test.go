package manifest

import (
	"errors"
	"testing"
)

func TestComputeAWSValidate(t *testing.T) {
	cases := []struct {
		name string
		cfg  ComputeAWS
		want error
	}{
		{"valid default tenancy", ComputeAWS{InstanceType: "t3.micro"}, nil},
		{"valid dedicated", ComputeAWS{InstanceType: "t3.micro", Tenancy: "dedicated"}, nil},
		{"valid host", ComputeAWS{InstanceType: "t3.micro", Tenancy: "host"}, nil},
		{"missing instance type", ComputeAWS{}, ErrMissingComputeInstanceType},
		{"invalid tenancy", ComputeAWS{InstanceType: "t3.micro", Tenancy: "weird"}, ErrInvalidTenancy},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.cfg.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("ComputeAWS.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
