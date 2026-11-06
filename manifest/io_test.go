package manifest

import (
	"errors"
	"path/filepath"
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

func TestManifestReadWriteRoundTrip(t *testing.T) {
	m := validManifest()
	path := filepath.Join(t.TempDir(), "crucible.yaml")

	if err := Write(&m, path); err != nil {
		t.Fatalf("Write: %v", err)
	}

	got, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if err := got.Validate(); err != nil {
		t.Fatalf("round-tripped manifest invalid: %v", err)
	}
	if _, err := As[*Service](got); err != nil {
		t.Fatalf("round-tripped config type: %v", err)
	}

	svc, err := ReadAs[*Service](path)
	if err != nil {
		t.Fatalf("ReadAs: %v", err)
	}
	if len(svc.Entrypoint) == 0 {
		t.Fatal("ReadAs service missing entrypoint")
	}
}

func TestManifestReadWriteAt(t *testing.T) {
	m := validManifest()
	dir := t.TempDir()

	if err := WriteAt(&m, dir); err != nil {
		t.Fatalf("WriteAt: %v", err)
	}
	if _, err := ReadAt(dir); err != nil {
		t.Fatalf("ReadAt: %v", err)
	}
	if _, err := ReadAsAt[*Service](dir); err != nil {
		t.Fatalf("ReadAsAt: %v", err)
	}
}

func TestManifestReadErrors(t *testing.T) {
	if _, err := Read(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("Read of missing file should fail")
	}
	if _, err := ReadAs[*Service](filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("ReadAs of missing file should fail")
	}
}

func TestPlanReadWriteRoundTrip(t *testing.T) {
	p := validPlan()
	path := filepath.Join(t.TempDir(), "plan.yaml")

	if err := WritePlan(&p, path); err != nil {
		t.Fatalf("WritePlan: %v", err)
	}
	got, err := ReadPlan(path)
	if err != nil {
		t.Fatalf("ReadPlan: %v", err)
	}
	if got.Version != p.Version {
		t.Fatalf("round-tripped plan version = %d, want %d", got.Version, p.Version)
	}
	if len(got.Deployments) != len(p.Deployments) {
		t.Fatalf("round-tripped deployments = %d, want %d", len(got.Deployments), len(p.Deployments))
	}
}

func TestPlanReadWriteAt(t *testing.T) {
	p := validPlan()
	dir := t.TempDir()

	if err := WritePlanAt(&p, dir); err != nil {
		t.Fatalf("WritePlanAt: %v", err)
	}
	if _, err := ReadPlanAt(dir); err != nil {
		t.Fatalf("ReadPlanAt: %v", err)
	}
}

func TestPlanReadError(t *testing.T) {
	if _, err := ReadPlan(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("ReadPlan of missing file should fail")
	}
}
