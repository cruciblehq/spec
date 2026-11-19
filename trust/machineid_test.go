package trust

import (
	"encoding/hex"
	"errors"
	"testing"
)

func TestNewMachineID(t *testing.T) {
	id := NewMachineID()
	if err := id.Validate(); err != nil {
		t.Errorf("NewMachineID().Validate() = %v, want nil", err)
	}
	if NewMachineID() == id {
		t.Error("NewMachineID() returned the same id twice, want different")
	}
}

func TestMachineIDName(t *testing.T) {
	id := MachineID("0123456789abcdef0123456789abcdef")
	if got := id.Name(); got != "0123456789abcdef0123456789abcdef.machine.crucible" {
		t.Errorf("Name() = %q, want %q", got, "0123456789abcdef0123456789abcdef.machine.crucible")
	}
}

func TestMachineIDValidate(t *testing.T) {
	if err := MachineID("0123456789abcdef0123456789abcdef").Validate(); err != nil {
		t.Errorf("Validate() = %v, want nil", err)
	}
	if err := MachineID("").Validate(); !errors.Is(err, ErrInvalidMachineID) {
		t.Errorf("empty Validate() = %v, want %v", err, ErrInvalidMachineID)
	}
	if err := MachineID("0123456789ABCDEF0123456789ABCDEF").Validate(); !errors.Is(err, ErrInvalidMachineID) {
		t.Errorf("uppercase Validate() = %v, want %v", err, ErrInvalidMachineID)
	}
	if err := MachineID("0123456789abcdef0123456789abcdeg").Validate(); !errors.Is(err, ErrInvalidMachineID) {
		t.Errorf("non hex Validate() = %v, want %v", err, ErrInvalidMachineID)
	}
	if err := MachineID("0123456789abcdef").Validate(); !errors.Is(err, ErrInvalidMachineID) {
		t.Errorf("short Validate() = %v, want %v", err, ErrInvalidMachineID)
	}
}

func TestRandHex(t *testing.T) {
	s := randHex(8)
	if len(s) != 16 {
		t.Fatalf("len(randHex(8)) = %d, want 16", len(s))
	}
	if _, err := hex.DecodeString(s); err != nil {
		t.Fatalf("randHex(8) = %q, want hex", s)
	}
	if s == randHex(8) {
		t.Fatalf("randHex(8) returned %q twice", s)
	}
}
