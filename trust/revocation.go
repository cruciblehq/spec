package trust

import (
	"crypto/x509"

	"github.com/cruciblehq/utils-go/crex"
)

// Parses a certificate revocation list and checks that root signed it.
//
// The list is a standard X.509 CRL (RFC 5280) naming the issuer and agent
// certificates the root has revoked. Its CRL Number orders lists, and an agent
// keeps the list with the highest number it has seen ([Newer]). The root
// certificate must allow CRL signing. Agents do not check NextUpdate, because
// the client sends its newest list on every handshake.
func ParseRevocations(der []byte, root *x509.Certificate) (*x509.RevocationList, error) {
	crl, err := x509.ParseRevocationList(der)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidRevocations, err)
	}
	if err := crl.CheckSignatureFrom(root); err != nil {
		return nil, crex.Wrap(ErrInvalidRevocations, err)
	}
	if crl.Number == nil {
		return nil, crex.Newf(ErrInvalidRevocations, "list from %q has no CRL number", crl.Issuer.CommonName)
	}
	return crl, nil
}

// Whether candidate has a higher CRL number than current.
func Newer(candidate, current *x509.RevocationList) bool {
	return candidate.Number.Cmp(current.Number) > 0
}

// Whether cert is listed in crl.
func revoked(crl *x509.RevocationList, cert *x509.Certificate) bool {
	for _, entry := range crl.RevokedCertificateEntries {
		if entry.SerialNumber.Cmp(cert.SerialNumber) == 0 {
			return true
		}
	}
	return false
}
