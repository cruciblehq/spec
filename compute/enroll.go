package compute

import "github.com/cruciblehq/spec/trust"

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

// Installs a machine certificate.
//
// crux sends this to [EndpointCertify] to finish enrollment. The root signs the
// certificate for the life of the machine, so it is never renewed. cruxd checks
// that the certificate is signed by the root, names its machine, and holds its
// key. Installing it consumes the enrollment token.
type CertifyRequest struct {
	Certificate []byte `codec:"certificate"` // DER encoded machine certificate.
}
