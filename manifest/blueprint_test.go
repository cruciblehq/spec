package manifest

import (
	"errors"
	"testing"
)

func TestBlueprintServiceAccessors(t *testing.T) {
	b := &Blueprint{}

	if err := b.AddService(Ref{ID: "a", Ref: "ns/name 1.0.0"}); err != nil {
		t.Fatalf("AddService: %v", err)
	}
	if got := b.Service("a"); got == nil || got.ID != "a" {
		t.Fatalf("Service(a) = %v, want ref a", got)
	}
	if b.Service("missing") != nil {
		t.Fatal("Service(missing) should be nil")
	}
	if err := b.AddService(Ref{ID: "a", Ref: "ns/name 1.0.0"}); !errors.Is(err, ErrServiceExists) {
		t.Fatalf("duplicate AddService = %v, want ErrServiceExists", err)
	}

	if err := b.RemoveService("a"); err != nil {
		t.Fatalf("RemoveService: %v", err)
	}
	if b.Service("a") != nil {
		t.Fatal("service a should be removed")
	}
	if err := b.RemoveService("a"); !errors.Is(err, ErrServiceNotFound) {
		t.Fatalf("RemoveService(missing) = %v, want ErrServiceNotFound", err)
	}
}

func TestBlueprintEnvironmentAccessors(t *testing.T) {
	b := &Blueprint{}

	if err := b.AddEnvironment(Environment{ID: "prod"}); err != nil {
		t.Fatalf("AddEnvironment: %v", err)
	}
	if got := b.Environment("prod"); got == nil || got.ID != "prod" {
		t.Fatalf("Environment(prod) = %v, want prod", got)
	}
	if b.Environment("missing") != nil {
		t.Fatal("Environment(missing) should be nil")
	}
	if err := b.AddEnvironment(Environment{ID: "prod"}); !errors.Is(err, ErrEnvironmentExists) {
		t.Fatalf("duplicate AddEnvironment = %v, want ErrEnvironmentExists", err)
	}

	if err := b.RemoveEnvironment("prod"); err != nil {
		t.Fatalf("RemoveEnvironment: %v", err)
	}
	if b.Environment("prod") != nil {
		t.Fatal("environment prod should be removed")
	}
	if err := b.RemoveEnvironment("prod"); !errors.Is(err, ErrEnvironmentNotFound) {
		t.Fatalf("RemoveEnvironment(missing) = %v, want ErrEnvironmentNotFound", err)
	}
}

func TestBlueprintValidate(t *testing.T) {
	cases := []struct {
		name string
		bp   Blueprint
		want error
	}{
		{"empty", Blueprint{}, nil},
		{
			"valid",
			Blueprint{
				Services:     []Ref{{ID: "api", Ref: "ns/api 1.0.0"}},
				Gateway:      Gateway{Routes: []Route{{Pattern: "/api", Service: "api"}}},
				Environments: []Environment{{ID: "prod"}},
			},
			nil,
		},
		{
			"invalid service ref",
			Blueprint{Services: []Ref{{ID: "api", Ref: ""}}},
			ErrInvalidBlueprint,
		},
		{
			"missing service id",
			Blueprint{Services: []Ref{{Ref: "ns/api 1.0.0"}}},
			ErrMissingServiceID,
		},
		{
			"duplicate service id",
			Blueprint{Services: []Ref{{ID: "api", Ref: "ns/api 1.0.0"}, {ID: "api", Ref: "ns/api 2.0.0"}}},
			ErrDuplicateServiceID,
		},
		{
			"route references unknown service",
			Blueprint{Gateway: Gateway{Routes: []Route{{Pattern: "/x", Service: "ghost"}}}},
			ErrRouteServiceNotFound,
		},
		{
			"invalid gateway route",
			Blueprint{
				Services: []Ref{{ID: "api", Ref: "ns/api 1.0.0"}},
				Gateway:  Gateway{Routes: []Route{{Pattern: "bad", Service: "api"}}},
			},
			ErrInvalidBlueprint,
		},
		{
			"invalid environment",
			Blueprint{Environments: []Environment{{ID: "Bad"}}},
			ErrInvalidBlueprint,
		},
		{
			"duplicate environment id",
			Blueprint{Environments: []Environment{{ID: "prod"}, {ID: "prod"}}},
			ErrDuplicateEnvironmentID,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.bp.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Blueprint.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
