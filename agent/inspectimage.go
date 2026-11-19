package agent

// Requests the configuration of an image.
type InspectImageRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Carries the configuration of an image.
type InspectImageResponse struct {
	Config ImageConfig `codec:"config"` // The image's configuration.
}
