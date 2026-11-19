package agent

import ocispec "github.com/opencontainers/image-spec/specs-go/v1"

// OCI image configuration.
//
// Alias of the OCI image-spec config type. This carries the execution defaults
// for an image (entrypoint, cmd, env, working directory, user, exposed ports,
// and so on).
type ImageConfig = ocispec.ImageConfig

// References an image in the agent's store.
//
// A name or digest understood by the image store. Produced by Import and Commit
// and accepted by every operation that acts on a stored image.
type ImageRef string
