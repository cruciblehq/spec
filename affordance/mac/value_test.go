package mac

import (
	"errors"
	"testing"
)

func TestMACValueValidateFieldEmptyName(t *testing.T) {
	v := &Value{IsField: true, Field: ""}
	if err := v.Validate(); !errors.Is(err, ErrInvalidValue) {
		t.Fatalf("err = %v, want ErrInvalidValue", err)
	}
}

func TestMACValueValidateFieldOK(t *testing.T) {
	v := &Value{IsField: true, Field: "task.uid"}
	if err := v.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACValueValidateLiteral(t *testing.T) {
	v := &Value{IsField: false, IntVal: 42}
	if err := v.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
