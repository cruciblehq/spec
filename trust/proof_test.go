package trust

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"testing"
)

func enrollSession(t *testing.T) (session, ed25519.PrivateKey) {
	t.Helper()
	ca := newTestCA(t)
	client, clientKey := ca.client(t)
	machine := selfSigned(t)
	s := handshake(t, serverConfig(ca, machine, ca.crl(t, 1)), EnrollConfig(tlsCert(clientKey, client, ca.issuer)))
	if s.serverErr != nil || s.clientErr != nil {
		t.Fatalf("handshake = (%v, %v), want (nil, nil)", s.serverErr, s.clientErr)
	}
	return s, machine.PrivateKey.(ed25519.PrivateKey)
}

func TestExporter(t *testing.T) {
	s, _ := enrollSession(t)
	other, _ := enrollSession(t)
	client, server, otherClient := s.client, s.server, other.client

	clientEKM, err := Exporter(client)
	if err != nil {
		t.Fatal(err)
	}
	serverEKM, err := Exporter(server)
	if err != nil {
		t.Fatal(err)
	}
	otherEKM, err := Exporter(otherClient)
	if err != nil {
		t.Fatal(err)
	}

	if len(clientEKM) != exporterSize {
		t.Errorf("len(Exporter()) = %d, want %d", len(clientEKM), exporterSize)
	}
	if !bytes.Equal(clientEKM, serverEKM) {
		t.Error("client and server Exporter() differ, want equal")
	}
	if bytes.Equal(clientEKM, otherEKM) {
		t.Error("Exporter() equal across sessions, want different")
	}
}

func TestNewToken(t *testing.T) {
	a := NewToken()
	b := NewToken()
	if len(a) != tokenSize {
		t.Errorf("len(NewToken()) = %d, want %d", len(a), tokenSize)
	}
	if bytes.Equal(a, b) {
		t.Error("NewToken() returned the same token twice, want different")
	}
}

func TestTokenProof(t *testing.T) {
	token := NewToken()
	ekm := make([]byte, exporterSize)
	rand.Read(ekm)
	proof := TokenProof(token, ekm)

	if proof.Type != ProofToken {
		t.Errorf("Type = %q, want %q", proof.Type, ProofToken)
	}
	if err := proof.VerifyToken(token, ekm); err != nil {
		t.Errorf("VerifyToken() = %v, want nil", err)
	}
}

func TestTokenProofMismatch(t *testing.T) {
	token := NewToken()
	other := NewToken()
	ekm := make([]byte, exporterSize)
	rand.Read(ekm)
	proof := TokenProof(token, ekm)

	if err := proof.VerifyToken(other, ekm); !errors.Is(err, ErrInvalidProof) {
		t.Errorf("other token VerifyToken() = %v, want %v", err, ErrInvalidProof)
	}
	if err := proof.VerifyToken(token, other); !errors.Is(err, ErrInvalidProof) {
		t.Errorf("other session VerifyToken() = %v, want %v", err, ErrInvalidProof)
	}
}

func TestTokenProofWrongType(t *testing.T) {
	token := NewToken()
	ekm := make([]byte, exporterSize)
	proof := TokenProof(token, ekm)
	proof.Type = "aws"

	if err := proof.VerifyToken(token, ekm); !errors.Is(err, ErrInvalidProof) {
		t.Errorf("VerifyToken() = %v, want %v", err, ErrInvalidProof)
	}
}
