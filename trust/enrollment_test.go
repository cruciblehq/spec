package trust

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"testing"
)

func certificateRequest(t *testing.T, key ed25519.PrivateKey) []byte {
	t.Helper()
	der, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{DNSNames: []string{testMachine.Name()}}, key)
	if err != nil {
		t.Fatal(err)
	}
	return der
}

func machineProof(t *testing.T, s session, token []byte) Proof {
	t.Helper()
	ekm, err := Exporter(s.server)
	if err != nil {
		t.Fatal(err)
	}
	return TokenProof(token, ekm)
}

func TestVerifyEnrollment(t *testing.T) {
	s, key := enrollSession(t)
	token := NewToken()

	csr, err := VerifyEnrollment(s.client, certificateRequest(t, key), machineProof(t, s, token), token)
	if err != nil {
		t.Fatalf("VerifyEnrollment() = %v, want nil", err)
	}
	if !key.Public().(ed25519.PublicKey).Equal(csr.PublicKey) {
		t.Error("request key does not match the machine key")
	}
}

func TestVerifyEnrollmentWrongToken(t *testing.T) {
	s, key := enrollSession(t)

	if _, err := VerifyEnrollment(s.client, certificateRequest(t, key), machineProof(t, s, NewToken()), NewToken()); !errors.Is(err, ErrInvalidProof) {
		t.Errorf("VerifyEnrollment() = %v, want %v", err, ErrInvalidProof)
	}
}

func TestVerifyEnrollmentOtherKey(t *testing.T) {
	s, _ := enrollSession(t)
	_, other, _ := ed25519.GenerateKey(rand.Reader)
	token := NewToken()

	if _, err := VerifyEnrollment(s.client, certificateRequest(t, other), machineProof(t, s, token), token); !errors.Is(err, ErrInvalidEnrollment) {
		t.Errorf("VerifyEnrollment() = %v, want %v", err, ErrInvalidEnrollment)
	}
}

func TestVerifyEnrollmentOtherSession(t *testing.T) {
	s, key := enrollSession(t)
	other, _ := enrollSession(t)
	token := NewToken()

	if _, err := VerifyEnrollment(s.client, certificateRequest(t, key), machineProof(t, other, token), token); !errors.Is(err, ErrInvalidProof) {
		t.Errorf("VerifyEnrollment() = %v, want %v", err, ErrInvalidProof)
	}
}

func TestVerifyEnrollmentMalformedRequest(t *testing.T) {
	s, _ := enrollSession(t)
	token := NewToken()

	if _, err := VerifyEnrollment(s.client, []byte("not a request"), machineProof(t, s, token), token); !errors.Is(err, ErrInvalidEnrollment) {
		t.Errorf("VerifyEnrollment() = %v, want %v", err, ErrInvalidEnrollment)
	}
}

func TestVerifyEnrollmentNoPeer(t *testing.T) {
	if _, err := VerifyEnrollment(tls.ConnectionState{}, nil, Proof{}, NewToken()); !errors.Is(err, ErrInvalidEnrollment) {
		t.Errorf("VerifyEnrollment() = %v, want %v", err, ErrInvalidEnrollment)
	}
}

func validEnrollment(t *testing.T) (Enrollment, testCA) {
	t.Helper()
	ca := newTestCA(t)
	return Enrollment{
		Machine:     testMachine,
		Root:        ca.root.Raw,
		Revocations: ca.revocations(t, 1),
		Token:       NewToken(),
	}, ca
}

func TestEnrollmentVerify(t *testing.T) {
	e, ca := validEnrollment(t)

	root, crl, err := e.Verify()
	if err != nil {
		t.Fatalf("Verify() = %v, want nil", err)
	}
	if !root.Equal(ca.root) {
		t.Error("Verify() root does not match the CA root")
	}
	if crl.Number.Int64() != 1 {
		t.Errorf("Number = %s, want 1", crl.Number)
	}
}

func TestEnrollmentVerifyBadMachine(t *testing.T) {
	e, _ := validEnrollment(t)
	e.Machine = "bad"

	if _, _, err := e.Verify(); !errors.Is(err, ErrInvalidMachineID) {
		t.Errorf("Verify() = %v, want %v", err, ErrInvalidMachineID)
	}
}

func TestEnrollmentVerifyBadToken(t *testing.T) {
	e, _ := validEnrollment(t)
	e.Token = []byte("short")

	if _, _, err := e.Verify(); !errors.Is(err, ErrInvalidEnrollment) {
		t.Errorf("Verify() = %v, want %v", err, ErrInvalidEnrollment)
	}
}

func TestEnrollmentVerifyBadRoot(t *testing.T) {
	e, _ := validEnrollment(t)
	e.Root = []byte("not a certificate")

	if _, _, err := e.Verify(); !errors.Is(err, ErrInvalidEnrollment) {
		t.Errorf("Verify() = %v, want %v", err, ErrInvalidEnrollment)
	}
}

func TestEnrollmentVerifyRevocationsFromOtherRoot(t *testing.T) {
	e, _ := validEnrollment(t)
	other := newTestCA(t)
	e.Revocations = other.revocations(t, 1)

	if _, _, err := e.Verify(); !errors.Is(err, ErrInvalidRevocations) {
		t.Errorf("Verify() = %v, want %v", err, ErrInvalidRevocations)
	}
}
