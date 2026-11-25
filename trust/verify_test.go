package trust

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"testing"
	"time"
)

var testNow = time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)

const testMachine = MachineID("0123456789abcdef0123456789abcdef")

type testCA struct {
	root      *x509.Certificate
	rootKey   ed25519.PrivateKey
	issuer    *x509.Certificate
	issuerKey ed25519.PrivateKey
}

func validFor(d time.Duration) Validity {
	return Validity{NotBefore: testNow.Add(-time.Hour), NotAfter: testNow.Add(d)}
}

func createCert(t *testing.T, tmpl, parent *x509.Certificate, pub crypto.PublicKey, signer crypto.Signer) *x509.Certificate {
	t.Helper()
	cert, err := Sign(tmpl, parent, pub, signer)
	if err != nil {
		t.Fatal(err)
	}
	return cert
}

func newTestCA(t *testing.T) testCA {
	t.Helper()
	rootPub, rootKey, _ := ed25519.GenerateKey(rand.Reader)
	root := createCert(t, AuthorityTemplate("root", validFor(RootLifetime), false), nil, rootPub, rootKey)
	issuerPub, issuerKey, _ := ed25519.GenerateKey(rand.Reader)
	issuer := createCert(t, AuthorityTemplate("issuer", validFor(IssuerLifetime), true), root, issuerPub, rootKey)
	return testCA{root: root, rootKey: rootKey, issuer: issuer, issuerKey: issuerKey}
}

func leafTemplate(usage x509.ExtKeyUsage, names []string) *x509.Certificate {
	return LeafTemplate("leaf", validFor(ClientLifetime), usage, names...)
}

func (ca testCA) client(t *testing.T) (*x509.Certificate, ed25519.PrivateKey) {
	t.Helper()
	pub, key, _ := ed25519.GenerateKey(rand.Reader)
	return createCert(t, leafTemplate(x509.ExtKeyUsageClientAuth, nil), ca.issuer, pub, ca.issuerKey), key
}

func (ca testCA) machine(t *testing.T) (*x509.Certificate, ed25519.PrivateKey) {
	t.Helper()
	pub, key, _ := ed25519.GenerateKey(rand.Reader)
	tmpl := leafTemplate(x509.ExtKeyUsageServerAuth, []string{testMachine.Name()})
	tmpl.NotAfter = testNow.Add(RootLifetime)
	return createCert(t, tmpl, ca.root, pub, ca.rootKey), key
}

func TestVerifyClient(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.client(t)

	if err := VerifyClient([][]*x509.Certificate{{leaf, ca.issuer, ca.root}}, ca.crl(t, 1)); err != nil {
		t.Errorf("VerifyClient() = %v, want nil", err)
	}
}

func TestVerifyClientSignedByRoot(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.machine(t)

	if err := VerifyClient([][]*x509.Certificate{{leaf, ca.root}}, ca.crl(t, 1)); !errors.Is(err, ErrInvalidChain) {
		t.Errorf("VerifyClient() = %v, want %v", err, ErrInvalidChain)
	}
}

func TestVerifyClientIssuerRevoked(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.client(t)

	if err := VerifyClient([][]*x509.Certificate{{leaf, ca.issuer, ca.root}}, ca.crl(t, 2, ca.issuer)); !errors.Is(err, ErrRevoked) {
		t.Errorf("VerifyClient() = %v, want %v", err, ErrRevoked)
	}
}

func TestVerifyClientNonEd25519(t *testing.T) {
	ca := newTestCA(t)
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	leaf := createCert(t, leafTemplate(x509.ExtKeyUsageClientAuth, nil), ca.issuer, &key.PublicKey, ca.issuerKey)

	if err := VerifyClient([][]*x509.Certificate{{leaf, ca.issuer, ca.root}}, ca.crl(t, 1)); !errors.Is(err, ErrInvalidChain) {
		t.Errorf("VerifyClient() = %v, want %v", err, ErrInvalidChain)
	}
}

func TestVerifyClientNoChains(t *testing.T) {
	ca := newTestCA(t)

	if err := VerifyClient(nil, ca.crl(t, 1)); !errors.Is(err, ErrInvalidChain) {
		t.Errorf("VerifyClient() = %v, want %v", err, ErrInvalidChain)
	}
}

func TestVerifyServer(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.machine(t)

	if err := VerifyServer([][]*x509.Certificate{{leaf, ca.root}}, ca.crl(t, 1)); err != nil {
		t.Errorf("VerifyServer() = %v, want nil", err)
	}
}

func TestVerifyServerSignedByIssuer(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.client(t)

	if err := VerifyServer([][]*x509.Certificate{{leaf, ca.issuer, ca.root}}, ca.crl(t, 1)); !errors.Is(err, ErrInvalidChain) {
		t.Errorf("VerifyServer() = %v, want %v", err, ErrInvalidChain)
	}
}

func TestVerifyServerRevoked(t *testing.T) {
	ca := newTestCA(t)
	leaf, _ := ca.machine(t)

	if err := VerifyServer([][]*x509.Certificate{{leaf, ca.root}}, ca.crl(t, 2, leaf)); !errors.Is(err, ErrRevoked) {
		t.Errorf("VerifyServer() = %v, want %v", err, ErrRevoked)
	}
}

func TestVerifyServerNoChains(t *testing.T) {
	ca := newTestCA(t)

	if err := VerifyServer(nil, ca.crl(t, 1)); !errors.Is(err, ErrInvalidChain) {
		t.Errorf("VerifyServer() = %v, want %v", err, ErrInvalidChain)
	}
}
