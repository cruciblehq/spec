package manifest

import (
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func TestManifestValidate_InterfaceResource(t *testing.T) {
	m := &Manifest{
		Version: 0,
		Resource: Resource{
			Type:    TypeInterface,
			Name:    "platform/mail",
			Version: "3.0.0",
		},
		Config: &Interface{},
	}

	if err := m.Validate(); err != nil {
		t.Fatalf("validate interface manifest: %v", err)
	}
}

func TestInterfaceReferenceMatches(t *testing.T) {
	ok, err := InterfaceReferenceMatches("platform/mail 3.0.0", "platform/mail 3.0.0")
	if err != nil {
		t.Fatalf("matching same version: %v", err)
	}
	if !ok {
		t.Fatal("same interface and version should match")
	}

	ok, err = InterfaceReferenceMatches("platform/mail >=3.0.0 <4.0.0", "platform/mail 3.1.0")
	if err != nil {
		t.Fatalf("matching range requirement: %v", err)
	}
	if !ok {
		t.Fatal("implementation should satisfy the range requirement")
	}

	ok, err = InterfaceReferenceMatches("platform/mail 3.0.0", "platform/mail 3.1.0")
	if err != nil {
		t.Fatalf("matching different version: %v", err)
	}
	if ok {
		t.Fatal("different interface version must not match")
	}

	ok, err = InterfaceReferenceMatches("platform/mail 3.0.0", "platform/sms 3.0.0")
	if err != nil {
		t.Fatalf("matching different interface name: %v", err)
	}
	if ok {
		t.Fatal("different interface name must not match")
	}
}

func TestServiceValidate_InterfaceDeclarations(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Implements: []string{"platform/mail 3.0.0", "platform/notifications 2.1.0"},
		Requires:   []string{"platform/mail 3.0.0"},
	}

	if err := svc.Validate(); err != nil {
		t.Fatalf("validate service interface declarations: %v", err)
	}
}

func TestServiceValidate_AttachesDeclaration(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Implements: []string{"platform/dashboard 1.0.0"},
		Attaches:   []string{"platform/mail >=3.0.0 <4.0.0"},
	}

	if err := svc.Validate(); err != nil {
		t.Fatalf("validate service attaches declarations: %v", err)
	}
}

func TestServiceValidate_RejectsInvalidInterfaceDeclaration(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Implements: []string{"platform/mail"},
	}

	if err := svc.Validate(); err == nil {
		t.Fatal("invalid interface declaration should fail validation")
	}
}

func TestServiceValidate_RejectsRangeInExactImplementation(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Implements: []string{"platform/mail >=3.0.0 <4.0.0"},
	}

	if err := svc.Validate(); err == nil {
		t.Fatal("range-based implementation declaration should fail validation")
	}
}

func TestServiceValidate_AllowsRangeRequirement(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Requires:   []string{"platform/mail >=3.0.0 <4.0.0"},
	}

	if err := svc.Validate(); err != nil {
		t.Fatalf("range-based requirement declaration should validate: %v", err)
	}
}

func TestServiceEncode_IncludesInterfaceLists(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{}}},
		Entrypoint: []string{"./app"},
		Implements: []string{"platform/mail 3.0.0"},
		Requires:   []string{"platform/cache >=2.0.0 <3.0.0"},
		Attaches:   []string{"platform/logging >=1.0.0 <2.0.0"},
	}

	m, err := svc.Encode(&codec.Codec{})
	if err != nil {
		t.Fatalf("encode service: %v", err)
	}
	data, ok := m.(map[string]any)
	if !ok {
		t.Fatalf("encoded service = %T, want map", m)
	}
	if _, ok := data["implements"]; !ok {
		t.Fatal("encoded service missing implements")
	}
	if _, ok := data["requires"]; !ok {
		t.Fatal("encoded service missing requires")
	}
	if _, ok := data["attaches"]; !ok {
		t.Fatal("encoded service missing attaches")
	}
}
