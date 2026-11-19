package agent

import (
	"context"
	"io"
)

// The agent on one machine, which runs and manages its deployments.
//
// An Agent is reached through one closable handle. It stores images and
// creates containers from them, and its responsibilities grow with the
// protocol. Calling any method after [Agent.Close] results in undefined
// behaviour. Containers are created with [Agent.Load] and own all operations
// on their writable snapshot.
type Agent interface {

	// Imports an OCI image archive into the agent's image store.
	//
	// r must be an OCI tar archive as produced by [Agent.Export]. The
	// returned reference identifies the image and can be passed to Load,
	// Export, Extract, Inspect, or Remove. The image is unpacked so all
	// layer snapshots exist.
	Import(ctx context.Context, r io.Reader) (ImageRef, error)

	// Exports the named image as an OCI tar archive written to w.
	//
	// ref must be a reference returned by [Agent.Import] or [Container.Commit].
	// The archive includes all layer blobs, manifests, and config descriptors
	// and can be re-imported via [Agent.Import].
	Export(ctx context.Context, ref ImageRef, w io.Writer) error

	// Extracts a directory tree from the named image as a tar stream.
	//
	// ref must be a reference returned by [Agent.Import] or [Container.Commit].
	// path is matched against the merged view of all image layers. The returned
	// [io.ReadCloser] must be closed by the caller.
	Extract(ctx context.Context, ref ImageRef, path string) (io.ReadCloser, error)

	// Returns the OCI image config for the named image.
	//
	// The returned [ImageConfig] holds the execution defaults baked into the
	// image, which are the values in effect before [Options] overrides them
	// at container start. ref must be a reference returned by [Agent.Import]
	// or [Container.Commit].
	Inspect(ctx context.Context, ref ImageRef) (ImageConfig, error)

	// Removes the named image from the store.
	//
	// ref must be a reference returned by [Agent.Import] or [Container.Commit].
	// The caller is responsible for ensuring no containers are using the image
	// before removing it.
	Remove(ctx context.Context, ref ImageRef) error

	// Creates a new container from a previously imported image.
	//
	// ref must be a reference returned by [Agent.Import] or [Container.Commit].
	// Allocates a persistent writable snapshot with the image as parent. The
	// container is not started until [Container.Start] is called and must be
	// released with [Container.Destroy] when no longer needed.
	Load(ctx context.Context, ref ImageRef, opts Options) (Container, error)

	// Checks that the agent can serve requests.
	//
	// Returns nil when the agent and the services it depends on are
	// reachable, and an error describing the failure otherwise.
	Health(ctx context.Context) error

	// Closes the handle and releases its underlying connection.
	//
	// Does not affect images or running containers. After Close, the handle
	// must not be used again.
	Close() error
}
