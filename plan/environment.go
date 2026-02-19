package plan

// Represents an environment configuration.
//
// Defines a set of environment variables that can be associated with deployments.
type Environment struct {
	ID        string            `json:"id"`        // Stable identifier for this environment set.
	Variables map[string]string `json:"variables"` // Key-value pairs injected as environment variables.
}
