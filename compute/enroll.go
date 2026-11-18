package compute

import "github.com/cruciblehq/spec/trust"

// Requests the machine's key and proof of identity.
type EnrollRequest struct {
}

// Machine's key and proof of identity.
//
// cruxd returns this from [EndpointEnroll] on a machine that has not enrolled
// yet. The machine serves its new key over TLS, and crux authenticates with its
// client certificate. Request is a certificate request signed by that same
// key, and Proof is bound to the TLS session. crux checks both before it signs.
type EnrollResponse struct {
	Request []byte      `codec:"request"` // DER encoded certificate request.
	Proof   trust.Proof `codec:"proof"`   // Evidence that this is the machine crux created.
}
