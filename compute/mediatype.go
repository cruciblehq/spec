package compute

// String identifier for HTTP Content-Type and Accept headers.
//
// Defines vendor-specific media types for the cruxd compute API following the
// pattern application/vnd.crucible.compute.{name}.v0. JSON message bodies use
// the [MediaType.JSON] form. Binary and streaming bodies use the type as-is.
type MediaType string

// Media types for JSON message bodies.
const (
	MediaTypeHandshakeRequest         MediaType = "application/vnd.crucible.compute.handshake-request.v0"          // [HandshakeRequest].
	MediaTypeHandshakeResponse        MediaType = "application/vnd.crucible.compute.handshake-response.v0"         // [HandshakeResponse].
	MediaTypeLoadRequest              MediaType = "application/vnd.crucible.compute.load-request.v0"               // [LoadRequest].
	MediaTypeLoadResponse             MediaType = "application/vnd.crucible.compute.load-response.v0"              // [LoadResponse].
	MediaTypeExportRequest            MediaType = "application/vnd.crucible.compute.export-request.v0"             // [ExportRequest].
	MediaTypeImportResponse           MediaType = "application/vnd.crucible.compute.import-response.v0"            // [ImportResponse].
	MediaTypeInspectImageRequest      MediaType = "application/vnd.crucible.compute.inspect-image-request.v0"      // [InspectImageRequest].
	MediaTypeInspectImageResponse     MediaType = "application/vnd.crucible.compute.inspect-image-response.v0"     // [InspectImageResponse].
	MediaTypeRemoveRequest            MediaType = "application/vnd.crucible.compute.remove-request.v0"             // [RemoveRequest].
	MediaTypeExtractRequest           MediaType = "application/vnd.crucible.compute.extract-request.v0"            // [ExtractRequest].
	MediaTypeStartRequest             MediaType = "application/vnd.crucible.compute.start-request.v0"              // [StartRequest].
	MediaTypeStopRequest              MediaType = "application/vnd.crucible.compute.stop-request.v0"               // [StopRequest].
	MediaTypeDestroyRequest           MediaType = "application/vnd.crucible.compute.destroy-request.v0"            // [DestroyRequest].
	MediaTypeCopyRequest              MediaType = "application/vnd.crucible.compute.copy-request.v0"               // [CopyRequest].
	MediaTypeInspectContainerRequest  MediaType = "application/vnd.crucible.compute.inspect-container-request.v0"  // [InspectContainerRequest].
	MediaTypeInspectContainerResponse MediaType = "application/vnd.crucible.compute.inspect-container-response.v0" // [InspectContainerResponse].
	MediaTypeExecRequest              MediaType = "application/vnd.crucible.compute.exec-request.v0"               // [ExecRequest].
	MediaTypeCommitRequest            MediaType = "application/vnd.crucible.compute.commit-request.v0"             // [CommitRequest].
	MediaTypeCommitResponse           MediaType = "application/vnd.crucible.compute.commit-response.v0"            // [CommitResponse].
	MediaTypeConfigureRequest         MediaType = "application/vnd.crucible.compute.configure-request.v0"          // [ConfigureRequest].
	MediaTypeError                    MediaType = "application/vnd.crucible.compute.error.v0"                      // [Error] returned on a non-2xx status.
)

// Media types for binary and streaming bodies, carried as-is.
const (
	MediaTypeImageArchive MediaType = "application/vnd.crucible.compute.image-archive.v0" // OCI image tar archive (import, export).
	MediaTypeLayer        MediaType = "application/vnd.crucible.compute.layer.v0"         // Uncompressed tar applied to a container (copy).
	MediaTypeFileTree     MediaType = "application/vnd.crucible.compute.file-tree.v0"     // Tar stream extracted from an image (extract).
	MediaTypeExecResponse MediaType = "application/vnd.crucible.compute.exec-response.v0" // Newline-delimited JSON stream of [ExecResponse] messages (exec).
)

// Returns the media type as a string.
func (m MediaType) String() string {
	return string(m)
}

// Returns the media type with the JSON syntax suffix (+json).
func (m MediaType) JSON() string {
	return string(m) + "+json"
}
