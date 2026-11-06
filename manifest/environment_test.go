package manifest

import (
	"errors"
	"testing"
)

func TestEnvironmentValidate(t *testing.T) {
	cases := []struct {
		name string
		env  Environment
		want error
	}{
		{"valid", Environment{ID: "prod"}, nil},
		{"missing id", Environment{ID: ""}, ErrMissingEnvironmentID},
		{"bad id", Environment{ID: "Bad"}, ErrInvalidEnvironmentID},
	}
	for _, c := range cases {
		if got := c.env.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
