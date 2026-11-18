package compute

// Requests application of a tar archive to a container.
//
// The message is the first part of a multipart body, and the archive
// ([MediaTypeLayer]) is the second.
type CopyRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
}

// Confirms that a tar archive was applied to a container.
type CopyResponse struct {
}
