package blueprint

// Defines a system composition.
//
// A Blueprint declares a service composition and configuration, being the main
// input to the deployment process. Blueprints prescribe which resources should
// be deployed together and how they are composed and exposed.
type Blueprint struct {

	// Schema version of the blueprint format.
	//
	// Must be the first field in  document. This value determines how the rest
	// of the blueprint is parsed. Currently only version 0 is supported.
	Version int `json:"version"`

	// Services to deploy.
	//
	// Each entry becomes a [plan.Service] after execution. Every service
	// is exposed through the gateway at its declared prefix.
	Services []Service `json:"services"`
}
