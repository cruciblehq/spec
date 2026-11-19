package trust

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cruciblehq/utils-go/crex"
)

// Number of random bytes in a machine identifier.
const machineIDSize = 16

// DNS suffix of every machine name.
const machineDomain = "machine.crucible"

// Identifier the client assigns to a machine when it creates it.
//
// The identifier is 16 random bytes in lowercase hexadecimal. It
// never changes for the life of the machine and does not depend on the
// machine's address.
type MachineID string

// Returns a new random machine identifier.
func NewMachineID() MachineID {
	return MachineID(randHex(machineIDSize))
}

// Returns the DNS name that identifies the machine in its agent's certificate.
//
// The client checks this name on every connection instead of the address it
// dialed, so port forwards and address changes do not affect identity.
func (id MachineID) Name() string {
	return fmt.Sprintf("%s.%s", id, machineDomain)
}

// Checks that the identifier is 32 lowercase hexadecimal characters.
func (id MachineID) Validate() error {
	s := string(id)
	if len(s) != hex.EncodedLen(machineIDSize) || strings.ToLower(s) != s {
		return crex.Newf(ErrInvalidMachineID, "%q is not %d lowercase hex characters", s, hex.EncodedLen(machineIDSize))
	}
	if _, err := hex.DecodeString(s); err != nil {
		return crex.Newf(ErrInvalidMachineID, "%q is not hexadecimal", s)
	}
	return nil
}

// Returns n random bytes as a lowercase hex string.
//
// crypto/rand.Read never fails, so there is no error to return.
func randHex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}
