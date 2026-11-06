package compute

import (
	"context"
	"io"
)

// A container runtime on a single machine.
//
// A Runtime is a handle to the container runtime on one machine, providing
// image operations and the creation of containers. The direct implementation
// inside cruxd talks to containerd; the client implementation inside crux
// forwards each call to cruxd over HTTP/2. The runtime owns its underlying
// connection, so calling any method after [Runtime.Close] results in undefined
// behaviour. Containers are created with [Runtime.Load] and own all operations
// on their writable snapshot.
type Runtime interface {

	// Imports an OCI image archive into the machine's image store.
	//
	// r must be an OCI tar archive as produced by [Runtime.Export]. The
	// returned reference identifies the image and can be passed to Load,
	// Export, Extract, Inspect, or Remove. The image is unpacked so all
	// layer snapshots exist.
	Import(ctx context.Context, r io.Reader) (ImageRef, error)

	// Exports the named image as an OCI tar archive written to w.
	//
	// ref must be a reference returned by [Runtime.Import] or [Container.Commit].
	// The archive includes all layer blobs, manifests, and config descriptors
	// and can be re-imported via [Runtime.Import].
	Export(ctx context.Context, ref ImageRef, w io.Writer) error

	// Extracts a directory tree from the named image as a tar stream.
	//
	// ref must be a reference returned by [Runtime.Import] or [Container.Commit].
	// path is matched against the merged view of all image layers. The returned
	// [io.ReadCloser] must be closed by the caller.
	Extract(ctx context.Context, ref ImageRef, path string) (io.ReadCloser, error)

	// Returns the OCI image config for the named image.
	//
	// The returned [ImageConfig] holds the runtime defaults baked into the
	// image, which are the values in effect before [Options] overrides them
	// at container start. ref must be a reference returned by [Runtime.Import]
	// or [Container.Commit].
	Inspect(ctx context.Context, ref ImageRef) (ImageConfig, error)

	// Removes the named image from the store.
	//
	// ref must be a reference returned by [Runtime.Import] or [Container.Commit].
	// The caller is responsible for ensuring no containers are using the image
	// before removing it.
	Remove(ctx context.Context, ref ImageRef) error

	// Creates a new container from a previously imported image.
	//
	// ref must be a reference returned by [Runtime.Import] or [Container.Commit].
	// Allocates a persistent writable snapshot with the image as parent. The
	// container is not started until [Container.Start] is called and must be
	// released with [Container.Destroy] when no longer needed.
	Load(ctx context.Context, ref ImageRef, opts Options) (Container, error)

	// Closes the runtime and releases its underlying connection.
	//
	// Does not affect images or running containers. After Close, the runtime
	// must not be used again.
	Close() error
}
