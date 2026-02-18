package reference

import "testing"

func TestParseIdentifier(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.test",
		DefaultNamespace: "official",
	}
	id, err := ParseIdentifier("namespace/name", "template", opts)
	if err != nil {
		t.Fatal(err)
	}

	if id.Type() != "template" {
		t.Errorf("expected type %q, got %q", "template", id.Type())
	}
	if id.Namespace() != "namespace" {
		t.Errorf("expected namespace %q, got %q", "namespace", id.Namespace())
	}
	if id.Name() != "name" {
		t.Errorf("expected name %q, got %q", "name", id.Name())
	}
}

func TestParseIdentifier_WithOptions(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://custom.registry.io",
		DefaultNamespace: "crucible",
	}

	id, err := ParseIdentifier("widget", "template", opts)
	if err != nil {
		t.Fatal(err)
	}

	reg := id.Registry()
	if reg.String() != "https://custom.registry.io" {
		t.Errorf("expected registry %q, got %q", "https://custom.registry.io", reg.String())
	}
	if id.Namespace() != "crucible" {
		t.Errorf("expected namespace %q, got %q", "crucible", id.Namespace())
	}
}

func TestParseIdentifier_Error(t *testing.T) {
	_, err := ParseIdentifier("", "template", IdentifierOptions{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMustParseIdentifier(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.test",
		DefaultNamespace: "official",
	}
	id := MustParseIdentifier("namespace/name", "template", opts)

	if id.Name() != "name" {
		t.Errorf("expected name %q, got %q", "name", id.Name())
	}
}

func TestMustParseIdentifier_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	MustParseIdentifier("", "template", IdentifierOptions{})
}

func TestIdentifier_Path_DefaultRegistry(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.test",
		DefaultNamespace: "official",
	}
	id := MustParseIdentifier("namespace/name", "template", opts)

	if id.Path() != "namespace/name" {
		t.Errorf("expected path %q, got %q", "namespace/name", id.Path())
	}
}

func TestIdentifier_Path_CustomRegistry(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.test",
		DefaultNamespace: "official",
	}
	id := MustParseIdentifier("myregistry.com/path/to/resource", "template", opts)

	if id.Path() != "path/to/resource" {
		t.Errorf("expected path %q, got %q", "path/to/resource", id.Path())
	}
}

func TestIdentifier_URI(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.crucible.net",
		DefaultNamespace: "official",
	}
	id := MustParseIdentifier("namespace/name", "template", opts)

	expected := "https://registry.crucible.net/namespace/name"
	if id.URI() != expected {
		t.Errorf("expected URI %q, got %q", expected, id.URI())
	}
}

func TestIdentifier_String(t *testing.T) {
	opts := IdentifierOptions{
		DefaultRegistry:  "https://registry.test",
		DefaultNamespace: "official",
	}
	id := MustParseIdentifier("namespace/name", "template", opts)

	expected := "template https://registry.test/namespace/name"
	if id.String() != expected {
		t.Errorf("expected string %q, got %q", expected, id.String())
	}
}
