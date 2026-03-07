package reference

import "testing"

// Returns Options with default test values.
func testOpts() Options {
	opts, _ := NewOptions("https://registry.test", "official")
	return opts
}

func TestParse(t *testing.T) {
	ref, err := Parse("namespace/name 1.0.0", "template", testOpts())
	if err != nil {
		t.Fatal(err)
	}

	if ref.Type() != "template" {
		t.Errorf("expected type %q, got %q", "template", ref.Type())
	}
	if ref.Namespace() != "namespace" {
		t.Errorf("expected namespace %q, got %q", "namespace", ref.Namespace())
	}
	if ref.Name() != "name" {
		t.Errorf("expected name %q, got %q", "name", ref.Name())
	}
}

func TestParse_WithOptions(t *testing.T) {
	opts, err := NewOptions("https://registry.test", "myteam")
	if err != nil {
		t.Fatal(err)
	}

	ref, err := Parse("widget 1.0.0", "template", opts)
	if err != nil {
		t.Fatal(err)
	}

	if ref.Namespace() != "myteam" {
		t.Errorf("expected namespace %q, got %q", "myteam", ref.Namespace())
	}
}

func TestParse_Error(t *testing.T) {
	_, err := Parse("", "template", testOpts())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMustParse(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	if ref.Name() != "name" {
		t.Errorf("expected name %q, got %q", "name", ref.Name())
	}
}

func TestMustParse_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	MustParse("", "template", testOpts())
}

func TestReference_Version(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	if ref.Version() == nil {
		t.Fatal("expected version, got nil")
	}
}

func TestReference_Channel(t *testing.T) {
	ref := MustParse("namespace/name :stable", "template", testOpts())

	if ref.Channel() == nil {
		t.Fatal("expected channel, got nil")
	}
	if *ref.Channel() != "stable" {
		t.Errorf("expected channel %q, got %q", "stable", *ref.Channel())
	}
}

func TestReference_Digest(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0 sha256:abcd1234", "template", testOpts())

	if ref.Digest() == nil {
		t.Fatal("expected digest, got nil")
	}
}

func TestReference_IsFrozen_True(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0 sha256:abcd1234", "template", testOpts())

	if !ref.IsFrozen() {
		t.Error("expected IsFrozen to be true")
	}
}

func TestReference_IsFrozen_False(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	if ref.IsFrozen() {
		t.Error("expected IsFrozen to be false")
	}
}

func TestReference_IsChannelBased_True(t *testing.T) {
	ref := MustParse("namespace/name :stable", "template", testOpts())

	if !ref.IsChannelBased() {
		t.Error("expected IsChannelBased to be true")
	}
}

func TestReference_IsChannelBased_False(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	if ref.IsChannelBased() {
		t.Error("expected IsChannelBased to be false")
	}
}

func TestReference_IsVersionBased_True(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	if !ref.IsVersionBased() {
		t.Error("expected IsVersionBased to be true")
	}
}

func TestReference_IsVersionBased_False(t *testing.T) {
	ref := MustParse("namespace/name :stable", "template", testOpts())

	if ref.IsVersionBased() {
		t.Error("expected IsVersionBased to be false")
	}
}

func TestReference_String_WithVersion(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0", "template", testOpts())

	s := ref.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}

func TestReference_String_WithChannel(t *testing.T) {
	ref := MustParse("namespace/name :stable", "template", testOpts())

	s := ref.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}

func TestReference_String_WithDigest(t *testing.T) {
	ref := MustParse("namespace/name 1.0.0 sha256:abcd1234", "template", testOpts())

	s := ref.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}
