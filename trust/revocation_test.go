package trust

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	"testing"
)

func (ca testCA) revocations(t *testing.T, number int64, revoked ...*x509.Certificate) []byte {
	t.Helper()
	var entries []x509.RevocationListEntry
	for _, cert := range revoked {
		entries = append(entries, x509.RevocationListEntry{SerialNumber: cert.SerialNumber, RevocationTime: testNow})
	}
	crl, err := SignRevocationList(number, Validity{NotBefore: testNow, NotAfter: testNow.Add(IssuerLifetime)}, entries, ca.root, ca.rootKey)
	if err != nil {
		t.Fatal(err)
	}
	return crl.Raw
}

func (ca testCA) crl(t *testing.T, number int64, revoked ...*x509.Certificate) *x509.RevocationList {
	t.Helper()
	crl, err := ParseRevocations(ca.revocations(t, number, revoked...), ca.root)
	if err != nil {
		t.Fatal(err)
	}
	return crl
}

func TestParseRevocations(t *testing.T) {
	ca := newTestCA(t)

	crl, err := ParseRevocations(ca.revocations(t, 4, ca.issuer), ca.root)
	if err != nil {
		t.Fatalf("ParseRevocations() = %v, want nil", err)
	}
	if crl.Number.Int64() != 4 {
		t.Errorf("Number = %s, want 4", crl.Number)
	}
	if !revoked(crl, ca.issuer) {
		t.Error("revoked(issuer) = false, want true")
	}
}

func TestParseRevocationsOtherRoot(t *testing.T) {
	ca := newTestCA(t)
	other := newTestCA(t)

	if _, err := ParseRevocations(other.revocations(t, 1), ca.root); !errors.Is(err, ErrInvalidRevocations) {
		t.Errorf("ParseRevocations() = %v, want %v", err, ErrInvalidRevocations)
	}
}

func TestParseRevocationsMalformed(t *testing.T) {
	ca := newTestCA(t)

	if _, err := ParseRevocations([]byte("not a list"), ca.root); !errors.Is(err, ErrInvalidRevocations) {
		t.Errorf("ParseRevocations() = %v, want %v", err, ErrInvalidRevocations)
	}
}

func TestNewer(t *testing.T) {
	ca := newTestCA(t)
	older := ca.crl(t, 1)
	newer := ca.crl(t, 2)

	if !Newer(newer, older) {
		t.Error("Newer(2, 1) = false, want true")
	}
	if Newer(older, newer) {
		t.Error("Newer(1, 2) = true, want false")
	}
	if Newer(older, older) {
		t.Error("Newer(1, 1) = true, want false")
	}
}

func TestRevokedOtherSerial(t *testing.T) {
	ca := newTestCA(t)
	crl := ca.crl(t, 1, &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "other"}})

	if revoked(crl, ca.issuer) {
		t.Error("revoked(issuer) = true, want false")
	}
}
