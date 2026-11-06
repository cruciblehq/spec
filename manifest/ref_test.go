package manifest

import (
	"errors"
	"testing"
)

func TestRefValidate(t *testing.T) {
	cases := []struct {
		name string
		ref  Ref
		want error
	}{
		{"valid no id", Ref{Ref: "ns/name 1.0.0"}, nil},
		{"valid with id", Ref{ID: "svc", Ref: "ns/name 1.0.0"}, nil},
		{"bad id", Ref{ID: "Bad", Ref: "ns/name 1.0.0"}, ErrInvalidRefID},
		{"missing target", Ref{Ref: ""}, ErrMissingRefTarget},
		{"bad arg key", Ref{Ref: "ns/name 1.0.0", Args: Args{"Bad-Key": "v"}}, ErrInvalidArgKey},
	}
	for _, c := range cases {
		if got := c.ref.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
