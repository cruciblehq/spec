package compute

// Requests extraction of a path from an image as a tar stream.
type ExtractRequest struct {
	Ref  ImageRef `codec:"ref"`  // Target image reference.
	Path string   `codec:"path"` // Path within the merged image view.
}

// Confirms the extraction of a path from an image.
//
// The message is the first part of a multipart body, and the tar stream
// ([MediaTypeFileTree]) is the second.
type ExtractResponse struct {
}
