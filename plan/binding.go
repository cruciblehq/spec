package plan

// Represents a binding of a service to compute infrastructure.
//
// Associates a service with a compute instance and optional environment
// configuration.
type Binding struct {
	Service     string `json:"service"`
	Compute     string `json:"compute"`
	Environment string `json:"environment,omitempty"`
}
