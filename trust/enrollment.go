package trust

import (
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"

	"github.com/cruciblehq/utils-go/crex"
)

// Path on the machine where crux places the enrollment data.
//
// Locally crux writes it through the machine's provisioning config, and in
// the cloud through instance user data. cruxd reads it on first boot.
const EnrollmentPath = "/etc/cruxd/enrollment.json"

// Data crux gives a new machine so it can enroll.
//
// The token is used once. After enrollment, cruxd keeps the root and CRL and
// discards the token, and later trust changes arrive only as newer revocation
// lists over mutual TLS.
type Enrollment struct {
	Machine     MachineID `codec:"machine"`     // Identifier crux assigned to the machine.
	Root        []byte    `codec:"root"`        // DER encoded root certificate.
	Revocations []byte    `codec:"revocations"` // DER encoded CRL signed by the root.
	Token       []byte    `codec:"token"`       // One-time enrollment token.
}

// Validates the enrollment and returns the parsed root and revocation list.
//
// The machine identifier must be valid, the token must have the expected
// size, the root must hold an Ed25519 key, and the revocation list must be
// signed by the root.
func (e Enrollment) Verify() (*x509.Certificate, *x509.RevocationList, error) {
	if err := e.Machine.Validate(); err != nil {
		return nil, nil, err
	}
	if len(e.Token) != tokenSize {
		return nil, nil, crex.Newf(ErrInvalidEnrollment, "token has %d bytes", len(e.Token))
	}
	root, err := x509.ParseCertificate(e.Root)
	if err != nil {
		return nil, nil, crex.Wrap(ErrInvalidEnrollment, err)
	}
	if _, ok := root.PublicKey.(ed25519.PublicKey); !ok {
		return nil, nil, crex.Newf(ErrInvalidEnrollment, "unexpected root key type %T", root.PublicKey)
	}
	crl, err := ParseRevocations(e.Revocations, root)
	if err != nil {
		return nil, nil, err
	}
	return root, crl, nil
}

// Checks a machine's enrollment response on the TLS session it arrived on.
//
// request is the machine's DER encoded certificate request and proof its token
// proof. The request must be signed by the Ed25519 key the machine presented in
// TLS, and the proof must match token on this session. Returns the parsed
// request for crux to sign.
func VerifyEnrollment(state tls.ConnectionState, request []byte, proof Proof, token []byte) (*x509.CertificateRequest, error) {
	if len(state.PeerCertificates) == 0 {
		return nil, crex.Newf(ErrInvalidEnrollment, "machine presented no certificate")
	}
	csr, err := x509.ParseCertificateRequest(request)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidEnrollment, err)
	}
	if err := csr.CheckSignature(); err != nil {
		return nil, crex.Wrap(ErrInvalidEnrollment, err)
	}
	key, ok := csr.PublicKey.(ed25519.PublicKey)
	if !ok {
		return nil, crex.Newf(ErrInvalidEnrollment, "unexpected request key type %T", csr.PublicKey)
	}
	if !key.Equal(state.PeerCertificates[0].PublicKey) {
		return nil, crex.Newf(ErrInvalidEnrollment, "request key does not match the key presented in TLS")
	}
	ekm, err := Exporter(state)
	if err != nil {
		return nil, err
	}
	if err := proof.VerifyToken(token, ekm); err != nil {
		return nil, err
	}
	return csr, nil
}
