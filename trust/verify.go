package trust

import (
	"crypto/ed25519"
	"crypto/x509"

	"github.com/cruciblehq/utils-go/crex"
)

// Lengths of the certificate chains that TLS verifies.
const (
	clientChainLength = 3 // Client leaf, issuer, and root.
	serverChainLength = 2 // Agent leaf and root.
)

// Checks the chains crypto/tls verified for a client certificate.
//
// TLS has already checked that each chain reaches the root, is in its validity
// period, and is for client authentication. This adds the rules TLS does not
// know: the leaf is an Ed25519 key signed by an issuer, never by the root
// directly, and the issuer is not listed in crl.
func VerifyClient(chains [][]*x509.Certificate, crl *x509.RevocationList) error {
	for _, chain := range chains {
		if len(chain) != clientChainLength {
			continue
		}
		if err := checkLeaf(chain[0]); err != nil {
			return err
		}
		if revoked(crl, chain[1]) {
			return crex.Newf(ErrRevoked, "issuer %q is revoked by CRL number %s", chain[1].Subject.CommonName, crl.Number)
		}
		return nil
	}
	return crex.Newf(ErrInvalidChain, "client certificate is not signed by an issuer")
}

// Checks the chains crypto/tls verified for an agent certificate.
//
// TLS has already checked that each chain reaches the root, names the machine,
// and is for server authentication. This adds the rules TLS does not know: the
// leaf is an Ed25519 key signed by the root directly, and it is not listed in
// crl. Agent certificates last for the life of their machine, so revocation is
// the only way to withdraw one. Call it from [tls.Config.VerifyConnection]
// with the connection's verified chains.
func VerifyServer(chains [][]*x509.Certificate, crl *x509.RevocationList) error {
	for _, chain := range chains {
		if len(chain) != serverChainLength {
			continue
		}
		if err := checkLeaf(chain[0]); err != nil {
			return err
		}
		if revoked(crl, chain[0]) {
			return crex.Newf(ErrRevoked, "machine %q is revoked by CRL number %s", chain[0].Subject.CommonName, crl.Number)
		}
		return nil
	}
	return crex.Newf(ErrInvalidChain, "machine certificate is not signed by the root")
}

// Checks that leaf holds an Ed25519 key.
func checkLeaf(leaf *x509.Certificate) error {
	if _, ok := leaf.PublicKey.(ed25519.PublicKey); !ok {
		return crex.Newf(ErrInvalidChain, "unexpected leaf key type %T", leaf.PublicKey)
	}
	return nil
}
