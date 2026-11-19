package trust

import (
	"crypto/tls"
	"crypto/x509"
)

// Returns the TLS settings shared by every connection between a client and an agent.
//
// Only TLS 1.3 is allowed, and the key exchange is restricted to the hybrid of
// X25519 and ML-KEM-768, so recorded traffic stays protected against a future
// quantum computer.
func baseConfig() *tls.Config {
	return &tls.Config{
		MinVersion:       tls.VersionTLS13,
		CurvePreferences: []tls.CurveID{tls.X25519MLKEM768},
	}
}

// Returns the TLS settings for an agent.
//
// The agent presents the certificate cert returns, which is its own key signed
// by itself before enrollment and its agent certificate after. Every client
// must present a certificate that TLS verifies against root and that passes
// [VerifyClient] with the list crl returns. Both functions are called on every
// handshake, so the agent can install a new certificate or list while serving.
func ServerConfig(cert func() *tls.Certificate, root *x509.Certificate, crl func() *x509.RevocationList) *tls.Config {
	roots := x509.NewCertPool()
	roots.AddCert(root)

	cfg := baseConfig()
	cfg.ClientAuth = tls.RequireAndVerifyClientCert
	cfg.ClientCAs = roots
	cfg.GetCertificate = func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
		return cert(), nil
	}
	cfg.VerifyConnection = func(cs tls.ConnectionState) error {
		return VerifyClient(cs.VerifiedChains, crl())
	}
	return cfg
}

// Returns the TLS settings for a client connecting to the agent on machine id.
//
// The client presents cert, its client certificate followed by its issuer. The
// agent must present a certificate that TLS verifies against root and the name
// of id, and that passes [VerifyServer] with crl. The name is checked instead
// of the address the client dials.
func ClientConfig(cert tls.Certificate, root *x509.Certificate, crl *x509.RevocationList, id MachineID) *tls.Config {
	roots := x509.NewCertPool()
	roots.AddCert(root)

	cfg := baseConfig()
	cfg.Certificates = []tls.Certificate{cert}
	cfg.RootCAs = roots
	cfg.ServerName = id.Name()
	cfg.VerifyConnection = func(cs tls.ConnectionState) error {
		return VerifyServer(cs.VerifiedChains, crl)
	}
	return cfg
}

// Returns the TLS settings for a client enrolling the agent on a new machine.
//
// An agent that has not enrolled presents a certificate it signed itself, so
// TLS cannot verify it. The client instead trusts the agent through the token
// proof, which is bound to this TLS session, and checks that the key in the
// certificate request is the key the agent presented ([VerifyEnrollment]).
func EnrollConfig(cert tls.Certificate) *tls.Config {
	cfg := baseConfig()
	cfg.Certificates = []tls.Certificate{cert}
	cfg.InsecureSkipVerify = true // NOSONAR: the agent is authenticated by the token proof.
	return cfg
}
