package compute

// Requests the health of cruxd and the runtime behind it.
type HealthRequest struct {
}

// Confirms that cruxd and the runtime behind it can serve requests.
type HealthResponse struct {
}
