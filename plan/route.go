package plan

// Represents a routing rule in the gateway.
//
// Maps request patterns to service instances.
type Route struct {
	Pattern string `json:"pattern"`
	Service string `json:"service"`
}
