package compute

import (
	"errors"
	"testing"
)

func TestErrorError(t *testing.T) {
	e := &Error{Code: ErrorCodeNotFound, Message: "missing"}
	if got := e.Error(); got != "not_found: missing" {
		t.Errorf("Error() = %q, want %q", got, "not_found: missing")
	}
}

func TestErrorValidate(t *testing.T) {
	cases := []struct {
		name string
		err  *Error
		want error
	}{
		{"valid", &Error{Code: ErrorCodeNotFound, Message: "x"}, nil},
		{"unknown code", &Error{Code: "nonsense", Message: "x"}, ErrErrorCodeInvalid},
		{"empty message", &Error{Code: ErrorCodeNotFound, Message: ""}, ErrErrorMessageEmpty},
	}
	for _, c := range cases {
		if got := c.err.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
