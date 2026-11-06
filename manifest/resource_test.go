package manifest

import (
	"errors"
	"testing"
)

func TestResourceValidate(t *testing.T) {
	cases := []struct {
		name string
		res  Resource
		want error
	}{
		{
			"valid",
			Resource{Type: TypeService, Name: "cruciblehq/my-api", Version: "1.0.0"},
			nil,
		},
		{
			"invalid type",
			Resource{Type: "", Name: "cruciblehq/my-api", Version: "1.0.0"},
			ErrInvalidResource,
		},
		{
			"missing name",
			Resource{Type: TypeService, Version: "1.0.0"},
			ErrMissingName,
		},
		{
			"missing version",
			Resource{Type: TypeService, Name: "cruciblehq/my-api"},
			ErrMissingVersion,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.res.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Resource.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
