package manifest

import (
	"errors"
	"testing"
)

func TestWidgetValidate(t *testing.T) {
	cases := []struct {
		name   string
		widget Widget
		want   error
	}{
		{"valid", Widget{Main: "index.js"}, nil},
		{"missing main", Widget{Main: ""}, ErrMissingMain},
	}
	for _, c := range cases {
		if got := c.widget.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
