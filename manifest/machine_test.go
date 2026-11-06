package manifest

import "testing"

func TestMachineValidate(t *testing.T) {
	var m Machine
	if err := m.Validate(); err != nil {
		t.Errorf("Validate() = %v, want nil", err)
	}
}
