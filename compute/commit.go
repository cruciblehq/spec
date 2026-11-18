package compute

// Requests a commit of a container's current state.
type CommitRequest struct {
	ID ContainerID `codec:"id"` // Target container handle.
	By string      `codec:"by"` // Creator recorded in the image history.
}

// Carries the reference of the image produced by a commit.
type CommitResponse struct {
	Ref ImageRef `codec:"ref"` // Reference of the image produced by the commit.
}
