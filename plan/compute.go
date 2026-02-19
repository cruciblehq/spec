package plan

// Represents a compute resource in the deployment plan.
//
// Defines the compute instance to provision. The Config field contains
// provider-specific configuration based on the Provider value.
type Compute struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Config   any    `json:"config,omitempty"`
}

// AWS compute configuration.
//
// Specifies EC2 instance settings for AWS deployments.
type ComputeAWS struct {
	InstanceType string `json:"instance_type"`
	Region       string `json:"region,omitempty"`
}

// Local compute configuration.
//
// No additional configuration needed for local deployments.
type ComputeLocal struct{}
