package manifest

import (
	"errors"
	"testing"
)

func TestGatewayValidate(t *testing.T) {
	cases := []struct {
		name    string
		gateway Gateway
		want    error
	}{
		{"empty", Gateway{}, nil},
		{"valid routes", Gateway{Routes: []Route{{Pattern: "/a", Service: "a"}, {Pattern: "/b", Service: "b"}}}, nil},
		{"invalid route", Gateway{Routes: []Route{{Pattern: "bad", Service: "a"}}}, ErrInvalidRoutePattern},
		{"duplicate pattern", Gateway{Routes: []Route{{Pattern: "/a", Service: "a"}, {Pattern: "/a", Service: "b"}}}, ErrDuplicateRoutePattern},
	}
	for _, c := range cases {
		if got := c.gateway.Validate(); !errors.Is(got, c.want) {
			t.Errorf("%s: Validate() = %v, want %v", c.name, got, c.want)
		}
	}
}
