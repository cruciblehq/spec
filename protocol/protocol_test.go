package protocol

import "testing"

func TestContainerID(t *testing.T) {
	tests := []struct {
		ref  string
		want string
	}{
		{"hub", "hub"},
		{"crucible/hub", "crucible-hub"},
		{"a/b/c", "a-b-c"},
		{"no-slashes", "no-slashes"},
	}
	for _, tt := range tests {
		if got := ContainerID(tt.ref); got != tt.want {
			t.Errorf("ContainerID(%q) = %q, want %q", tt.ref, got, tt.want)
		}
	}
}
