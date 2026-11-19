package agent

import "github.com/cruciblehq/spec/trust"

// Requests the agent's key and proof that it runs on the machine the client
// created.
type EnrollRequest struct {
}

// Carries the agent's key and its proof of identity.
//
// The agent returns this from [EndpointEnroll] while its machine has not
// enrolled. The agent serves its new key over TLS, and the client authenticates
// with its client certificate. Request is a certificate request signed by that
// same key, and Proof is bound to the TLS session. The client checks both
// before it signs.
type EnrollResponse struct {
	Request []byte      `codec:"request"` // DER encoded certificate request.
	Proof   trust.Proof `codec:"proof"`   // Evidence that the agent runs on the machine the client created.
}
