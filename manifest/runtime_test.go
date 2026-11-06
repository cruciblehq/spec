package manifest

import (
	"errors"
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func TestRuntimeValidate(t *testing.T) {
	cases := []struct {
		name string
		rt   Runtime
		want error
	}{
		{
			"valid",
			Runtime{Recipe: Recipe{Stages: []Stage{{From: "img"}}}},
			nil,
		},
		{
			"valid with schema",
			Runtime{
				Recipe: Recipe{Stages: []Stage{{From: "img"}}},
				Schema: &Schema{Params: []Param{{Name: "ok"}}},
			},
			nil,
		},
		{
			"invalid schema",
			Runtime{
				Recipe: Recipe{Stages: []Stage{{From: "img"}}},
				Schema: &Schema{Params: []Param{{Name: ""}}},
			},
			ErrInvalidRecipe,
		},
		{"empty recipe", Runtime{}, ErrInvalidRecipe},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.rt.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Runtime.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestRuntimeEncode(t *testing.T) {
	rt := &Runtime{
		Recipe: Recipe{Stages: []Stage{{From: "img"}}},
		Schema: &Schema{Params: []Param{{Name: "p"}}},
	}
	m, err := rt.Encode(&codec.Codec{})
	if err != nil {
		t.Fatalf("Runtime.Encode: %v", err)
	}
	out := m.(map[string]any)
	if _, ok := out["stages"]; !ok {
		t.Fatalf("runtime encode missing stages: %v", out)
	}
	if _, ok := out["schema"]; !ok {
		t.Fatalf("runtime encode missing schema: %v", out)
	}
}
