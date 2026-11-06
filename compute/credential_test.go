package compute

import "testing"

func TestCredentialHeader(t *testing.T) {
	if got := Credential("").Header(); got != "" {
		t.Errorf("empty Header() = %q, want empty", got)
	}
	if got := Credential("tok123").Header(); got != "Bearer tok123" {
		t.Errorf("Header() = %q, want %q", got, "Bearer tok123")
	}
}

func TestParseCredential(t *testing.T) {
	cred := Credential("tok123")
	parsed, ok := ParseCredential(cred.Header())
	if !ok || parsed != cred {
		t.Errorf("ParseCredential(%q) = (%q, %v), want (%q, true)", cred.Header(), parsed, ok, cred)
	}

	if _, ok := ParseCredential("tok123"); ok {
		t.Error("ParseCredential without scheme returned ok")
	}
	if _, ok := ParseCredential(""); ok {
		t.Error("ParseCredential of empty value returned ok")
	}
}
