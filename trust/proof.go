package trust

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"

	"github.com/cruciblehq/utils-go/crex"
)

// Number of random bytes in an enrollment token.
const tokenSize = 32

// TLS exporter label for enrollment channel binding (RFC 5705).
const exporterLabel = "EXPORTER-crucible-enrollment"

// Number of bytes exported from the TLS session for channel binding.
const exporterSize = 32

// Kind of evidence an agent offers to prove it runs on the machine the client created.
type ProofType string

const (

	// Proof made with the enrollment token the client placed on the machine.
	ProofToken ProofType = "token"
)

// Evidence that an agent runs on the machine the client created.
//
// The proof is bound to the TLS session it is sent over through a TLS exporter
// ([Exporter]), so it cannot be replayed on another connection or moved to
// another agent key.
type Proof struct {
	Type  ProofType `codec:"type"`  // Kind of evidence.
	Value []byte    `codec:"value"` // Evidence, interpreted according to Type.
}

// Returns a new random enrollment token.
func NewToken() []byte {
	token := make([]byte, tokenSize)
	rand.Read(token)
	return token
}

// Returns the channel binding value of a TLS session.
//
// Both ends of the connection compute the same value, and no other session
// produces it.
func Exporter(state tls.ConnectionState) ([]byte, error) {
	ekm, err := state.ExportKeyingMaterial(exporterLabel, nil, exporterSize)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidProof, err)
	}
	return ekm, nil
}

// Returns a token proof for the TLS session with channel binding ekm.
//
// The value is HMAC-SHA256 keyed with the token over ekm.
func TokenProof(token, ekm []byte) Proof {
	return Proof{
		Type:  ProofToken,
		Value: tokenMAC(token, ekm),
	}
}

// Checks that p is a token proof for token on the session with channel binding
// ekm.
func (p Proof) VerifyToken(token, ekm []byte) error {
	if p.Type != ProofToken {
		return crex.Newf(ErrInvalidProof, "proof type %q is not %q", p.Type, ProofToken)
	}
	if !hmac.Equal(p.Value, tokenMAC(token, ekm)) {
		return crex.Newf(ErrInvalidProof, "token proof does not match")
	}
	return nil
}

// Computes HMAC-SHA256 keyed with token over ekm.
func tokenMAC(token, ekm []byte) []byte {
	mac := hmac.New(sha256.New, token)
	mac.Write(ekm)
	return mac.Sum(nil)
}
