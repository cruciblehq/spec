package manifest

import (
	"errors"
	"testing"
)

func TestSchemaValidate(t *testing.T) {
	cases := []struct {
		name   string
		schema Schema
		want   error
	}{
		{"empty", Schema{}, nil},
		{"valid with default", Schema{Default: "foo", Params: []Param{{Name: "foo"}}}, nil},
		{"invalid param", Schema{Params: []Param{{Name: ""}}}, ErrMissingParamName},
		{
			"duplicate param",
			Schema{Params: []Param{{Name: "foo"}, {Name: "foo"}}},
			ErrDuplicateParamName,
		},
		{"default invalid name", Schema{Default: "Bad Name"}, ErrInvalidParamName},
		{"default not in schema", Schema{Default: "missing"}, ErrDefaultNotInSchema},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.schema.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Schema.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
