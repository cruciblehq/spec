package agent

// Requests export of an image as an OCI archive.
type ExportRequest struct {
	Ref ImageRef `codec:"ref"` // Target image reference.
}

// Confirms the export of an image.
//
// The message is the first part of a multipart body, and the archive
// ([MediaTypeImageArchive]) is the second.
type ExportResponse struct {
}
