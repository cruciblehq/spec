package plan

// Identifies the infrastructure provider that the deployment plan targets.
//
// Controls how compute resources are configured in the generated [plan.Plan].
// Each provider has its own configuration schema for compute resources, and
// the provider type determines which schema is used. For example, the "aws"
// provider type indicates that compute resources should be configured with
// AWS-specific settings, while the "local" provider type indicates that no
// additional configuration is needed for local deployments.
type ProviderType string

const (

	// Targets Amazon Web Services. Compute entries in the plan will
	// contain [plan.ComputeAWS] configuration.
	ProviderTypeAWS ProviderType = "aws"

	// Targets the local machine. No additional compute configuration
	// is generated.
	ProviderTypeLocal ProviderType = "local"
)
