package agent

// Requests the health of the agent and the services it depends on.
type HealthRequest struct {
}

// Confirms that the agent and the services it depends on can serve requests.
type HealthResponse struct {
}
