package compute

import (
	"context"
)

// HTTP header with the client credential (on every request).
const HeaderAuthorization = "Authorization"

// Scheme prefix for the [HeaderAuthorization] value.
const AuthScheme = "Bearer"

// Verifies client credentials.
//
// cruxd calls Authenticate for every request before dispatching it and rejects
// the request when a non-nil error is returned. crux presents its credentials
// in the [HeaderAuthorization] header. The implementation lives in cruxd, but
// the interface is shared.
type Authenticator interface {

	// Authenticates a credential.
	//
	// A nil error admits the request. [ErrUnauthenticated] means the request
	// is rejected.
	Authenticate(ctx context.Context, cred Credential) error
}
