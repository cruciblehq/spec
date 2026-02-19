package plan

// Represents a routing rule in the gateway.
//
// Maps request patterns to service instances.
type Route struct {
	Pattern string `json:"pattern"` // URL path prefix to match (e.g. "/api/hub").
	Service string `json:"service"` // Service ID to route matched requests to.
}

// Validates that the route has a pattern and service.
func (r *Route) validate() error {
	if r.Pattern == "" {
		return ErrMissingPattern
	}
	if r.Service == "" {
		return ErrMissingRouteSvc
	}
	return nil
}
