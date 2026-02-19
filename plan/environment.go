package plan

// Represents an environment configuration.
//
// Defines a set of environment variables that can be associated with deployments.
type Environment struct {
	ID        string            `json:"id"`
	Variables map[string]string `json:"variables"`
}
