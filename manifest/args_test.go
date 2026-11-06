package manifest

import (
	"errors"
	"testing"
)

func TestArgsValidate(t *testing.T) {
	if err := (Args{"foo": "1", "bar_baz": "2"}).Validate(); err != nil {
		t.Errorf("valid args: %v", err)
	}
	if err := Args(nil).Validate(); err != nil {
		t.Errorf("nil args: %v", err)
	}
	if err := (Args{"Bad-Key": "x"}).Validate(); !errors.Is(err, ErrInvalidArgKey) {
		t.Errorf("bad key: got %v, want ErrInvalidArgKey", err)
	}
}
