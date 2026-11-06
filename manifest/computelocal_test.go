package manifest

import (
	"errors"
	"testing"
)

func TestComputeLocalValidate(t *testing.T) {
	if err := (&ComputeLocal{Host: "10.0.0.1"}).Validate(); err != nil {
		t.Fatalf("valid local compute: %v", err)
	}
	if err := (&ComputeLocal{}).Validate(); !errors.Is(err, ErrMissingComputeHost) {
		t.Fatalf("missing host = %v, want ErrMissingComputeHost", err)
	}
}
