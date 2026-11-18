package compute

// Requests removal of an image.
type RemoveRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Confirms the removal of an image.
type RemoveResponse struct {
}
