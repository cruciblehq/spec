package compute

import (
	"fmt"
	"strings"
)

// A bearer token.
//
// The client (crux) presents it in the [HeaderAuthorization] header on every
// request and the server (cruxd) verifies it through an [Authenticator]. The
// contract treats it as a single opaque string.
type Credential string

// Formats the credential as a [HeaderAuthorization] value.
//
// The result is the [AuthScheme] prefix followed by the token, suitable for
// the request header. An empty credential yields an empty string.
func (c Credential) Header() string {
	if c == "" {
		return ""
	}
	return fmt.Sprintf("%s %s", AuthScheme, string(c))
}

// Parses a [HeaderAuthorization] value into a [Credential].
//
// The value must include the [AuthScheme] prefix (the surrounding whitespace
// is trimmed). The second result (bool) indicates whether the value used the
// expected scheme, so a missing or malformed header is distinguishable from
// an empty token.
func ParseCredential(header string) (Credential, bool) {
	token, ok := strings.CutPrefix(header, fmt.Sprintf("%s ", AuthScheme))
	if !ok {
		return "", false
	}
	return Credential(strings.TrimSpace(token)), true
}
