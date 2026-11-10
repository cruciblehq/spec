package compute

// Version of the wire contract.
//
// Used for compatibility between crux and cruxd. The version is sent by crux
// during handshake. cruxd rejects clients with incompatible protocol versions.
const Version = "0"

// Endpoint paths exposed by cruxd.
const (
	EndpointHealth           = "/health"             // Liveness probe.
	EndpointHandshake        = "/handshake"          // Compatibility check and cruxd identity.
	EndpointImport           = "/images/import"      // Import an OCI archive.
	EndpointExport           = "/images/export"      // Export an image as an OCI archive.
	EndpointExtract          = "/images/extract"     // Extract a path from an image.
	EndpointInspect          = "/images/inspect"     // Read an image config.
	EndpointRemove           = "/images/remove"      // Remove an image.
	EndpointLoad             = "/containers/load"    // Create a container from an image.
	EndpointStart            = "/containers/start"   // Start a container's main process.
	EndpointStop             = "/containers/stop"    // Stop a container's main process.
	EndpointExec             = "/containers/exec"    // Run a command (streaming).
	EndpointCopy             = "/containers/copy"    // Apply a tar archive to a container.
	EndpointConfig           = "/containers/config"  // Set a container's pending config.
	EndpointInspectContainer = "/containers/inspect" // Read a container's config.
	EndpointCommit           = "/containers/commit"  // Commit a container to an image.
	EndpointDestroy          = "/containers/destroy" // Release a container.
)

// HTTP headers.
const (
	HeaderContentType = "Content-Type" // HTTP header carrying the message media type.
)

// Query parameters.
const (
	QueryParamID = "id" // Container handle on [EndpointCopy].
)
