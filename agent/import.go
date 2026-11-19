package agent

// Requests the import of an OCI archive into the image store.
//
// The message is the first part of a multipart body, and the archive
// ([MediaTypeImageArchive]) is the second.
type ImportRequest struct {
}

// Carries the reference of a newly imported image.
type ImportResponse struct {
	Ref ImageRef `codec:"ref"` // Reference of the imported image.
}
