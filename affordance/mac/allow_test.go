package mac

import (
	"errors"
	"testing"
)

func TestMACAllowValidateEmptyHook(t *testing.T) {
	a := &Allow{Hook: ""}
	if err := a.Validate(); !errors.Is(err, ErrInvalidAllow) {
		t.Fatalf("err = %v, want ErrInvalidAllow", err)
	}
}

func TestMACAllowValidateNilWhere(t *testing.T) {
	a := &Allow{Hook: "file_open"}
	if err := a.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACAllowValidateInvalidWhere(t *testing.T) {
	a := &Allow{
		Hook:  "file_open",
		Where: &Expr{Type: "bogus"},
	}
	if err := a.Validate(); !errors.Is(err, ErrInvalidAllow) {
		t.Fatalf("err = %v, want ErrInvalidAllow", err)
	}
}
