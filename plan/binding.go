package plan

// Represents a binding of a service to compute infrastructure.
//
// Associates a service with a compute instance and optional environment
// configuration.
type Binding struct {
	Service     string `json:"service"`               // Service ID to bind.
	Compute     string `json:"compute"`               // Compute resource ID to run the service on.
	Environment string `json:"environment,omitempty"` // Environment set ID to inject (optional).
}

// Validates that the binding references a service and compute resource.
func (b *Binding) validate() error {
	if b.Service == "" {
		return ErrMissingBindSvc
	}
	if b.Compute == "" {
		return ErrMissingBindCompute
	}
	return nil
}
