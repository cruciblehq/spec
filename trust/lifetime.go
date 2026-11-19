package trust

import "time"

// Certificate lifetimes.
//
// The root rarely signs and lasts for years. Issuer certificates expire after
// a month, so a leaked issuer key stops working on its own. Client certificates
// are minted each time the client runs. Agent certificates are signed by the
// root at enrollment and last as long as their machine, so the client never has
// to come back to renew them.
const (
	RootLifetime   = 20 * 365 * 24 * time.Hour // Root certificate validity.
	IssuerLifetime = 30 * 24 * time.Hour       // Issuer certificate validity.
	ClientLifetime = 24 * time.Hour            // Client certificate validity.
)
