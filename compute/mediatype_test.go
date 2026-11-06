package compute

import "testing"

func TestMediaTypeString(t *testing.T) {
	mt := MediaTypeLoadRequest
	if got := mt.String(); got != string(mt) {
		t.Errorf("String() = %q, want %q", got, string(mt))
	}
}

func TestMediaTypeJSON(t *testing.T) {
	mt := MediaTypeLoadRequest
	want := string(mt) + "+json"
	if got := mt.JSON(); got != want {
		t.Errorf("JSON() = %q, want %q", got, want)
	}
}
