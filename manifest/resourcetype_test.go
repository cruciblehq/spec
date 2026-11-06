package manifest

import (
	"errors"
	"testing"
)

func TestParseResourceType(t *testing.T) {
	valid := []string{"machine", "runtime", "service", "interface", "template", "widget", "blueprint"}
	for _, s := range valid {
		rt, err := ParseResourceType(s)
		if err != nil || string(rt) != s {
			t.Errorf("ParseResourceType(%q) = (%q, %v), want (%q, nil)", s, rt, err, s)
		}
	}
	for _, s := range []string{"", "unknown", "Machine"} {
		if _, err := ParseResourceType(s); !errors.Is(err, ErrInvalidResourceType) {
			t.Errorf("ParseResourceType(%q) err = %v, want ErrInvalidResourceType", s, err)
		}
	}
}
