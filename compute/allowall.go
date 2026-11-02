package compute

import "context"

// Placeholder [Authenticator].
//
// This implementation admits every credential without verification, until a
// real authenticator is implemented. For now, it's just a placeholder.
type AllowAll struct{}

// Admits every credential.
func (AllowAll) Authenticate(ctx context.Context, cred Credential) error {
	return nil
}
