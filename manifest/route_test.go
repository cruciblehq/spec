package manifest

import (
	"errors"
	"testing"
)

func TestRouteValidate(t *testing.T) {
	cases := []struct {
		name  string
		route Route
		want  error
	}{
		{"valid", Route{Pattern: "/api", Service: "api"}, nil},
		{"bad pattern", Route{Pattern: "api", Service: "api"}, ErrInvalidRoutePattern},
		{"bad service", Route{Pattern: "/api", Service: "Bad"}, ErrInvalidRouteService},
	}
	for _, c := range cases {
		if got := c.route.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
