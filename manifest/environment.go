package manifest

// A named set of environment variables.
//
// Environments provide concrete values for parameters declared by services.
// The specific environment to use is selected at build time, impacting the
// resulting deployment plan.
type Environment struct {
	ID        string            `codec:"id"`        // Unique identifier for this environment.
	Variables map[string]string `codec:"variables"` // Variable names and their values.
}

// Validates the environment entry.
func (e *Environment) Validate() error {
	if e.ID == "" {
		return ErrMissingEnvironmentID
	}
	if !isValidName(e.ID) {
		return ErrInvalidEnvironmentID
	}
	return nil
}
