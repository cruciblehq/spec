package manifest

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestAs(t *testing.T) {
	m := validManifest()
	svc, err := As[*Service](&m)
	if err != nil {
		t.Fatalf("As[*Service]: %v", err)
	}
	if len(svc.Entrypoint) == 0 {
		t.Fatal("asserted service missing entrypoint")
	}
	if _, err := As[*Runtime](&m); !errors.Is(err, ErrConfigTypeMismatch) {
		t.Fatalf("As[*Runtime] = %v, want ErrConfigTypeMismatch", err)
	}
}

func TestManifestEncodeDecodeRoundTrip(t *testing.T) {
	m := validManifest()
	var buf bytes.Buffer
	if err := Encode(&buf, &m); err != nil {
		t.Fatalf("Encode: %v", err)
	}
	got, err := Decode(&buf)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if err := got.Validate(); err != nil {
		t.Fatalf("round-tripped manifest invalid: %v", err)
	}
	if _, err := As[*Service](got); err != nil {
		t.Fatalf("round-tripped config type: %v", err)
	}
}

func TestManifestDecodeError(t *testing.T) {
	if _, err := Decode(strings.NewReader("version: [")); !errors.Is(err, ErrDecodeFailed) {
		t.Fatalf("Decode(malformed) = %v, want ErrDecodeFailed", err)
	}
}

func TestPlanEncodeDecodeRoundTrip(t *testing.T) {
	p := validPlan()
	var buf bytes.Buffer
	if err := EncodePlan(&buf, &p); err != nil {
		t.Fatalf("EncodePlan: %v", err)
	}
	got, err := DecodePlan(&buf)
	if err != nil {
		t.Fatalf("DecodePlan: %v", err)
	}
	if got.Version != p.Version {
		t.Fatalf("round-tripped plan version = %d, want %d", got.Version, p.Version)
	}
	if len(got.Deployments) != len(p.Deployments) {
		t.Fatalf("round-tripped deployments = %d, want %d", len(got.Deployments), len(p.Deployments))
	}
}

func TestPlanDecodeError(t *testing.T) {
	if _, err := DecodePlan(strings.NewReader("version: [")); !errors.Is(err, ErrDecodeFailed) {
		t.Fatalf("DecodePlan(malformed) = %v, want ErrDecodeFailed", err)
	}
}
