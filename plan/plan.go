package plan

// Current plan format version.
const Version = 0

// Represents a deployment plan.
//
// Specifies what resources will be deployed and the infrastructure configuration
// required to run them. Generated during the planning phase by resolving
// references, allocating infrastructure, and determining routing.
type Plan struct {
	Version      int           `json:"version"`
	Services     []Service     `json:"services"`
	Compute      []Compute     `json:"compute"`
	Environments []Environment `json:"environments,omitempty"`
	Bindings     []Binding     `json:"bindings"`
	Gateway      Gateway       `json:"gateway"`
}
