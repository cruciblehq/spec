package manifest

import (
	"errors"
	"testing"
)

func TestRecipeValidate(t *testing.T) {
	cases := []struct {
		name   string
		recipe Recipe
		want   error
	}{
		{"no stages", Recipe{}, ErrInvalidRecipe},
		{
			"single valid stage",
			Recipe{Stages: []Stage{{From: "crucible/runtime 0.1.0"}}},
			nil,
		},
		{
			"duplicate stage names",
			Recipe{Stages: []Stage{{Name: "build"}, {Name: "build"}}},
			ErrDuplicateStageName,
		},
		{
			"invalid stage propagates",
			Recipe{Stages: []Stage{{Name: "bad name"}}},
			ErrInvalidRecipe,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.recipe.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Recipe.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestRecipeOutputStage(t *testing.T) {
	if (&Recipe{}).OutputStage() != nil {
		t.Fatal("empty recipe should have nil output stage")
	}
	r := &Recipe{Stages: []Stage{{Name: "a"}, {Name: "b"}}}
	out := r.OutputStage()
	if out == nil || out.Name != "b" {
		t.Fatalf("OutputStage() = %+v, want stage b", out)
	}
}
