package agent

import (
	"context"
	"io"
)

// Identifies a container.
//
// The agent returns it from Load, and the client supplies it on every later
// operation on that container. The identifier is opaque to the client.
type ContainerID string

// A container with a persistent writable snapshot.
//
// A Container is created with [Agent.Load] and holds a writable snapshot
// over the base image. [Container.Start] starts the main process, which
// [Container.Stop] terminates; the snapshot persists across Start/Stop cycles.
// [Container.Exec] runs a command and [Container.Copy] applies a tar archive,
// both accumulating changes in the snapshot. [Container.Commit] captures the
// snapshot as a new image. The container and snapshot are released with
// [Container.Destroy].
type Container interface {

	// Starts the container's main process.
	//
	// The command is derived from the Entrypoint and Cmd of the base image. A
	// pending configuration set with [Container.Configure] does not change it.
	// The container must not already be running; a second call returns an error.
	Start(ctx context.Context) error

	// Stops the container's main process.
	//
	// Sends a graceful termination signal and escalates to a kill after a grace
	// period. The writable snapshot is not affected. Safe to call on an
	// already-stopped container.
	Stop(ctx context.Context) error

	// Runs a command in the container and returns its exit code.
	//
	// The container must be started. The command is passed as a single
	// argument to opts.Shell, or to the default shell. Changes to the
	// filesystem accumulate in the writable snapshot and persist for later
	// calls. Process output is written to stdout and stderr. A non-zero exit
	// code is a normal result, not an error. The error is reserved for a
	// command that could not be run or waited on.
	Exec(ctx context.Context, stdout, stderr io.Writer, command string, opts Options) (int, error)

	// Applies an uncompressed tar archive to the container's writable snapshot.
	//
	// r must be an uncompressed tar archive. Entries are applied as absolute
	// paths within the container filesystem, preserving permissions, ownership,
	// and timestamps. OCI whiteout entries follow standard layer semantics.
	// The container's main process is not affected.
	Copy(ctx context.Context, r io.Reader) error

	// Stores cfg as the pending configuration that [Container.Commit] applies.
	//
	// Does not create an image or affect the writable snapshot. Typically
	// called after [Container.Inspect] to modify specific fields while
	// preserving the rest. Returns an error when the configuration cannot be
	// stored.
	Configure(ctx context.Context, cfg ImageConfig) error

	// Returns the container's current OCI image config.
	//
	// If a pending configuration was set via [Container.Configure], that is
	// returned; otherwise the config is read from the base image. The caller
	// owns the returned value and may modify it before passing it back to
	// [Container.Configure].
	Inspect(ctx context.Context) (ImageConfig, error)

	// Creates a new image from the container's current state.
	//
	// Captures the snapshot diff and any pending configuration, recording a
	// history entry with by as the creator. Returns a reference that can be
	// passed to [Agent.Load], Export, Extract, Inspect, or Remove. The live
	// container is not affected.
	Commit(ctx context.Context, by string) (ImageRef, error)

	// Releases the container's writable snapshot.
	//
	// After Destroy returns, no other methods may be called on the container.
	Destroy(ctx context.Context) error
}
