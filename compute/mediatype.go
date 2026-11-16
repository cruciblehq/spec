package compute

import (
	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/mediatype"
)

// Encoding of message bodies in the compute API.
const BodyFormat = codec.JSON

// Media types of the cruxd compute API.
//
// Each follows the pattern application/vnd.crucible.compute.{name}.v0. Message
// bodies encoded in [BodyFormat] carry the type with that format's suffix (see
// [mediatype.MediaType.With]). Binary and streaming bodies use the type as is.

// Media types for message bodies encoded in [BodyFormat].
const (
	MediaTypeHandshakeRequest         mediatype.MediaType = "application/vnd.crucible.compute.handshake-request.v0"          // [HandshakeRequest].
	MediaTypeHandshakeResponse        mediatype.MediaType = "application/vnd.crucible.compute.handshake-response.v0"         // [HandshakeResponse].
	MediaTypeEnrollResponse           mediatype.MediaType = "application/vnd.crucible.compute.enroll-response.v0"            // [EnrollResponse].
	MediaTypeCertifyRequest           mediatype.MediaType = "application/vnd.crucible.compute.certify-request.v0"            // [CertifyRequest].
	MediaTypeLoadRequest              mediatype.MediaType = "application/vnd.crucible.compute.load-request.v0"               // [LoadRequest].
	MediaTypeLoadResponse             mediatype.MediaType = "application/vnd.crucible.compute.load-response.v0"              // [LoadResponse].
	MediaTypeExportRequest            mediatype.MediaType = "application/vnd.crucible.compute.export-request.v0"             // [ExportRequest].
	MediaTypeImportResponse           mediatype.MediaType = "application/vnd.crucible.compute.import-response.v0"            // [ImportResponse].
	MediaTypeInspectImageRequest      mediatype.MediaType = "application/vnd.crucible.compute.inspect-image-request.v0"      // [InspectImageRequest].
	MediaTypeInspectImageResponse     mediatype.MediaType = "application/vnd.crucible.compute.inspect-image-response.v0"     // [InspectImageResponse].
	MediaTypeRemoveRequest            mediatype.MediaType = "application/vnd.crucible.compute.remove-request.v0"             // [RemoveRequest].
	MediaTypeExtractRequest           mediatype.MediaType = "application/vnd.crucible.compute.extract-request.v0"            // [ExtractRequest].
	MediaTypeStartRequest             mediatype.MediaType = "application/vnd.crucible.compute.start-request.v0"              // [StartRequest].
	MediaTypeStopRequest              mediatype.MediaType = "application/vnd.crucible.compute.stop-request.v0"               // [StopRequest].
	MediaTypeDestroyRequest           mediatype.MediaType = "application/vnd.crucible.compute.destroy-request.v0"            // [DestroyRequest].
	MediaTypeCopyRequest              mediatype.MediaType = "application/vnd.crucible.compute.copy-request.v0"               // [CopyRequest].
	MediaTypeInspectContainerRequest  mediatype.MediaType = "application/vnd.crucible.compute.inspect-container-request.v0"  // [InspectContainerRequest].
	MediaTypeInspectContainerResponse mediatype.MediaType = "application/vnd.crucible.compute.inspect-container-response.v0" // [InspectContainerResponse].
	MediaTypeExecRequest              mediatype.MediaType = "application/vnd.crucible.compute.exec-request.v0"               // [ExecRequest].
	MediaTypeCommitRequest            mediatype.MediaType = "application/vnd.crucible.compute.commit-request.v0"             // [CommitRequest].
	MediaTypeCommitResponse           mediatype.MediaType = "application/vnd.crucible.compute.commit-response.v0"            // [CommitResponse].
	MediaTypeConfigureRequest         mediatype.MediaType = "application/vnd.crucible.compute.configure-request.v0"          // [ConfigureRequest].
	MediaTypeError                    mediatype.MediaType = "application/vnd.crucible.compute.error.v0"                      // [Error] returned on a non-2xx status.
)

// Media types for binary and streaming bodies, carried as-is.
const (
	MediaTypeImageArchive mediatype.MediaType = "application/vnd.crucible.compute.image-archive.v0" // OCI image tar archive (import, export).
	MediaTypeLayer        mediatype.MediaType = "application/vnd.crucible.compute.layer.v0"         // Uncompressed tar applied to a container (copy).
	MediaTypeFileTree     mediatype.MediaType = "application/vnd.crucible.compute.file-tree.v0"     // Tar stream extracted from an image (extract).
	MediaTypeExecResponse mediatype.MediaType = "application/vnd.crucible.compute.exec-response.v0" // Newline-delimited JSON stream of [ExecResponse] messages (exec).
)
