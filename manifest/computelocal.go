package manifest

// Local machine configuration.
//
// The identity of a machine that is a directly reachable host. Unlike cloud
// providers, no provisioning API is involved. The executor connects to the
// host as it is and schedules services onto it.
type ComputeLocal struct {
	Host string `codec:"host,omitempty"` // Hostname or IP address of the target machine.
}

// Validates the local compute configuration.
func (c *ComputeLocal) Validate() error {
	if c.Host == "" {
		return ErrMissingComputeHost
	}
	return nil
}
