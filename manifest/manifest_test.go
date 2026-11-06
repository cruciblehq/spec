package manifest

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func validManifest() Manifest {
	return Manifest{
		Version:  0,
		Resource: Resource{Type: TypeService, Name: "cruciblehq/my-api", Version: "1.0.0"},
		Config: &Service{
			Recipe:     Recipe{Stages: []Stage{{From: "img"}}},
			Entrypoint: []string{"/bin/app"},
		},
	}
}

func TestManifestValidate(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*Manifest)
		want   error
	}{
		{"valid", func(*Manifest) {}, nil},
		{"bad version", func(m *Manifest) { m.Version = 1 }, ErrUnsupportedVersion},
		{"invalid resource", func(m *Manifest) { m.Resource.Name = "" }, ErrInvalidManifest},
		{"nil config", func(m *Manifest) { m.Config = nil }, ErrMissingConfig},
		{"config type mismatch", func(m *Manifest) { m.Config = &Runtime{} }, ErrConfigTypeMismatch},
		{
			"config validation error",
			func(m *Manifest) { m.Config = &Service{Recipe: Recipe{Stages: []Stage{{From: "img"}}}} },
			ErrMissingEntrypoint,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := validManifest()
			c.mutate(&m)
			if got := m.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Manifest.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestManifestRoundTripByType(t *testing.T) {
	cases := []struct {
		name   string
		typ    ResourceType
		config any
	}{
		{"machine", TypeMachine, &Machine{}},
		{"runtime", TypeRuntime, &Runtime{Recipe: Recipe{Stages: []Stage{{From: "img"}}}}},
		{"widget", TypeWidget, &Widget{Main: "index.js"}},
		{"template", TypeTemplate, &Template{}},
		{"blueprint", TypeBlueprint, &Blueprint{}},
		{"interface", TypeInterface, &Interface{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			checkManifestRoundTrip(t, c.typ, c.config)
		})
	}
}

func checkManifestRoundTrip(t *testing.T, typ ResourceType, config any) {
	t.Helper()
	m := Manifest{
		Version:  0,
		Resource: Resource{Type: typ, Name: "cruciblehq/thing", Version: "1.0.0"},
		Config:   config,
	}
	if err := m.Validate(); err != nil {
		t.Fatalf("validate: %v", err)
	}
	path := filepath.Join(t.TempDir(), "crucible.yaml")
	if err := Write(&m, path); err != nil {
		t.Fatalf("write: %v", err)
	}
	got, err := Read(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if got.Resource.Type != typ {
		t.Fatalf("round-tripped type = %q, want %q", got.Resource.Type, typ)
	}
	if err := got.Validate(); err != nil {
		t.Fatalf("round-tripped invalid: %v", err)
	}
}

func TestManifestDecodeRejectsNonMap(t *testing.T) {
	if err := (&Manifest{}).Decode(&codec.Codec{}, "not a map"); !errors.Is(err, ErrDecodeFailed) {
		t.Fatalf("Manifest.Decode(non-map) = %v, want ErrDecodeFailed", err)
	}
}

func TestManifestDecodeUnknownType(t *testing.T) {
	err := (&Manifest{}).Decode(&codec.Codec{}, map[string]any{"resource": map[string]any{"type": "bogus/thing"}})
	if !errors.Is(err, ErrInvalidResourceType) {
		t.Fatalf("Manifest.Decode(unknown type) = %v, want ErrInvalidResourceType", err)
	}
}
