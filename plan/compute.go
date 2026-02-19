package plan

// Represents a compute resource in the deployment plan.
//
// Defines the compute instance to provision. The Config field contains
// provider-specific configuration based on the Provider value.
type Compute struct {
	ID       string `json:"id"`               // Stable identifier for this compute resource.
	Provider string `json:"provider"`         // Infrastructure provider (e.g. "aws", "local").
	Config   any    `json:"config,omitempty"` // Provider-specific configuration ([ComputeAWS] or [ComputeLocal]).
}

// AWS compute configuration.
//
// Specifies EC2 instance settings for AWS deployments.
type ComputeAWS struct {
	InstanceType string `json:"instance_type"`    // EC2 instance type (e.g. "t3.micro").
	Region       string `json:"region,omitempty"` // AWS region for the instance.
}

// Local compute configuration.
//
// No additional configuration needed for local deployments.
type ComputeLocal struct{}

// Validates that the compute resource has an ID and provider.
func (c *Compute) validate() error {
	if c.ID == "" {
		return ErrMissingComputeID
	}
	if c.Provider == "" {
		return ErrMissingProvider
	}
	return nil
}
