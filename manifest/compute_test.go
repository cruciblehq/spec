package manifest

import (
	"errors"
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func TestComputeValidate(t *testing.T) {
	cases := []struct {
		name    string
		compute Compute
		want    error
	}{
		{"nil config", Compute{Type: ComputeTypeLocal}, ErrMissingComputeProvider},
		{"not validatable", Compute{Type: ComputeTypeLocal, Config: "raw"}, ErrInvalidComputeType},
		{"valid local", Compute{Type: ComputeTypeLocal, Config: &ComputeLocal{Host: "h"}}, nil},
		{"valid aws", Compute{Type: ComputeTypeAWS, Config: &ComputeAWS{InstanceType: "t3.micro"}}, nil},
		{"config error propagates", Compute{Type: ComputeTypeLocal, Config: &ComputeLocal{}}, ErrMissingComputeHost},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.compute.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Compute.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestComputeEncode(t *testing.T) {
	c := &Compute{Type: ComputeTypeLocal, Config: &ComputeLocal{Host: "h"}}
	m, err := c.Encode(&codec.Codec{})
	if err != nil {
		t.Fatalf("Compute.Encode: %v", err)
	}
	out := m.(map[string]any)
	if out["type"] != ComputeTypeLocal {
		t.Fatalf("compute encode type = %v, want local", out["type"])
	}

	c = &Compute{Type: ComputeTypeAWS}
	m, err = c.Encode(&codec.Codec{})
	if err != nil {
		t.Fatalf("Compute.Encode(nil config): %v", err)
	}
	out = m.(map[string]any)
	if out["type"] != ComputeTypeAWS {
		t.Fatalf("compute encode type = %v, want aws", out["type"])
	}
}

func TestComputeDecodeRejectsNonMap(t *testing.T) {
	if err := (&Compute{}).Decode(&codec.Codec{}, "not a map"); !errors.Is(err, ErrDecodeFailed) {
		t.Fatalf("Compute.Decode(non-map) = %v, want ErrDecodeFailed", err)
	}
}

func TestComputeDecodeUnknownType(t *testing.T) {
	err := (&Compute{}).Decode(&codec.Codec{}, map[string]any{"type": "bogus"})
	if !errors.Is(err, ErrInvalidComputeType) {
		t.Fatalf("Compute.Decode(unknown) = %v, want ErrInvalidComputeType", err)
	}
}
