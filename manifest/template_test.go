package manifest

import (
	"errors"
	"testing"
)

func TestTemplateValidate(t *testing.T) {
	if err := (&Template{}).Validate(); err != nil {
		t.Fatalf("nil schema template: %v", err)
	}
	if err := (&Template{Schema: &Schema{Params: []Param{{Name: "ok"}}}}).Validate(); err != nil {
		t.Fatalf("valid schema template: %v", err)
	}
	if err := (&Template{Schema: &Schema{Params: []Param{{Name: ""}}}}).Validate(); !errors.Is(err, ErrMissingParamName) {
		t.Fatalf("invalid schema = %v, want ErrMissingParamName", err)
	}
}
