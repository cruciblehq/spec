package agent

// Version of the wire contract.
//
// Used for compatibility between a client and the agent. The client sends it
// during handshake, and the agent rejects clients with incompatible versions.
const Version = "0"

// Port the agent listens on.
const DefaultPort = 2789

// Endpoint paths the agent serves.
const (
	EndpointHealth           = "/health"             // Health of the agent.
	EndpointHandshake        = "/handshake"          // Compatibility check and agent identity.
	EndpointEnroll           = "/enroll"             // Agent key and proof of identity.
	EndpointCertify          = "/certify"            // Install the agent certificate.
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
