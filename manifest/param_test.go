package manifest

import (
	"errors"
	"testing"
)

func TestParamValidate(t *testing.T) {
	cases := []struct {
		name  string
		param Param
		want  error
	}{
		{"required", Param{Name: "count"}, nil},
		{"string default", Param{Name: "shell", Default: "/bin/sh"}, nil},
		{"int default", Param{Name: "n", Default: 3}, nil},
		{"bool default", Param{Name: "flag", Default: true}, nil},
		{"missing name", Param{Name: ""}, ErrMissingParamName},
		{"bad name", Param{Name: "Bad"}, ErrInvalidParamName},
		{"bad default type", Param{Name: "x", Default: []string{"a"}}, ErrInvalidParamDefault},
	}
	for _, c := range cases {
		if got := c.param.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
