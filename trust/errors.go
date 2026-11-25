package trust

import "github.com/cruciblehq/utils-go/crex"

var (
	ErrInvalidMachineID   = crex.New("invalid machine id")
	ErrInvalidRevocations = crex.New("invalid revocation list")
	ErrInvalidProof       = crex.New("invalid enrollment proof")
	ErrInvalidEnrollment  = crex.New("invalid enrollment")
	ErrInvalidChain       = crex.New("invalid certificate chain")
	ErrRevoked            = crex.New("certificate revoked")
	ErrInvalidKey         = crex.New("invalid key")
)
