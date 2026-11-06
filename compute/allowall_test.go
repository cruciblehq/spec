package compute

import (
	"context"
	"testing"
)

func TestAllowAllAuthenticate(t *testing.T) {
	var a AllowAll
	for _, cred := range []Credential{"", "token", "anything"} {
		if err := a.Authenticate(context.Background(), cred); err != nil {
			t.Errorf("Authenticate(%q) = %v, want nil", cred, err)
		}
	}
}
