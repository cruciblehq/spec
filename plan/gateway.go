package plan

// Represents the API gateway configuration.
//
// Defines how external requests are routed to deployed services.
type Gateway struct {
	Routes []Route `json:"routes,omitempty"` // Routing rules that map request patterns to services.
}
