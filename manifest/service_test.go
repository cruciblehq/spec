package manifest

import (
	"errors"
	"testing"
)

func TestServiceValidate(t *testing.T) {
	validRecipe := Recipe{Stages: []Stage{{From: "img"}}}
	cases := []struct {
		name string
		svc  Service
		want error
	}{
		{
			"valid",
			Service{Recipe: validRecipe, Entrypoint: []string{"/bin/app"}},
			nil,
		},
		{
			"missing entrypoint",
			Service{Recipe: validRecipe},
			ErrMissingEntrypoint,
		},
		{
			"invalid schema",
			Service{Recipe: validRecipe, Entrypoint: []string{"/bin/app"}, Schema: &Schema{Params: []Param{{Name: ""}}}},
			ErrMissingParamName,
		},
		{
			"invalid recipe",
			Service{Entrypoint: []string{"/bin/app"}},
			ErrInvalidRecipe,
		},
		{
			"invalid requires ref",
			Service{Recipe: validRecipe, Entrypoint: []string{"/bin/app"}, Requires: []string{"::bad::"}},
			ErrInvalidRef,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.svc.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Service.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
