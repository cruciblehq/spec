package agent

import (
	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/mediatype"
)

// Encoding of message bodies in the agent protocol.
const BodyFormat = codec.JSON

// Media types of the agent protocol.
//
// Each follows the pattern application/vnd.crucible.agent.{name}.v0. Message
// bodies encoded in [BodyFormat] carry the type with that format's suffix (see
// [mediatype.MediaType.With]). Binary data travels in a multipart/related body
// (RFC 2387). Its first part is the operation's message, whose media type is
// also the type parameter of the body, and its second part is the data, which
// uses its type as is.

// Media types for message bodies encoded in [BodyFormat].
const (
	MediaTypeHandshakeRequest         mediatype.MediaType = "application/vnd.crucible.agent.handshake-request.v0"          // [HandshakeRequest].
	MediaTypeHandshakeResponse        mediatype.MediaType = "application/vnd.crucible.agent.handshake-response.v0"         // [HandshakeResponse].
	MediaTypeHealthRequest            mediatype.MediaType = "application/vnd.crucible.agent.health-request.v0"             // [HealthRequest].
	MediaTypeHealthResponse           mediatype.MediaType = "application/vnd.crucible.agent.health-response.v0"            // [HealthResponse].
	MediaTypeEnrollRequest            mediatype.MediaType = "application/vnd.crucible.agent.enroll-request.v0"             // [EnrollRequest].
	MediaTypeEnrollResponse           mediatype.MediaType = "application/vnd.crucible.agent.enroll-response.v0"            // [EnrollResponse].
	MediaTypeCertifyRequest           mediatype.MediaType = "application/vnd.crucible.agent.certify-request.v0"            // [CertifyRequest].
	MediaTypeCertifyResponse          mediatype.MediaType = "application/vnd.crucible.agent.certify-response.v0"           // [CertifyResponse].
	MediaTypeLoadRequest              mediatype.MediaType = "application/vnd.crucible.agent.load-request.v0"               // [LoadRequest].
	MediaTypeLoadResponse             mediatype.MediaType = "application/vnd.crucible.agent.load-response.v0"              // [LoadResponse].
	MediaTypeExportRequest            mediatype.MediaType = "application/vnd.crucible.agent.export-request.v0"             // [ExportRequest].
	MediaTypeExportResponse           mediatype.MediaType = "application/vnd.crucible.agent.export-response.v0"            // [ExportResponse].
	MediaTypeImportRequest            mediatype.MediaType = "application/vnd.crucible.agent.import-request.v0"             // [ImportRequest].
	MediaTypeImportResponse           mediatype.MediaType = "application/vnd.crucible.agent.import-response.v0"            // [ImportResponse].
	MediaTypeInspectImageRequest      mediatype.MediaType = "application/vnd.crucible.agent.inspect-image-request.v0"      // [InspectImageRequest].
	MediaTypeInspectImageResponse     mediatype.MediaType = "application/vnd.crucible.agent.inspect-image-response.v0"     // [InspectImageResponse].
	MediaTypeRemoveRequest            mediatype.MediaType = "application/vnd.crucible.agent.remove-request.v0"             // [RemoveRequest].
	MediaTypeRemoveResponse           mediatype.MediaType = "application/vnd.crucible.agent.remove-response.v0"            // [RemoveResponse].
	MediaTypeExtractRequest           mediatype.MediaType = "application/vnd.crucible.agent.extract-request.v0"            // [ExtractRequest].
	MediaTypeExtractResponse          mediatype.MediaType = "application/vnd.crucible.agent.extract-response.v0"           // [ExtractResponse].
	MediaTypeStartRequest             mediatype.MediaType = "application/vnd.crucible.agent.start-request.v0"              // [StartRequest].
	MediaTypeStartResponse            mediatype.MediaType = "application/vnd.crucible.agent.start-response.v0"             // [StartResponse].
	MediaTypeStopRequest              mediatype.MediaType = "application/vnd.crucible.agent.stop-request.v0"               // [StopRequest].
	MediaTypeStopResponse             mediatype.MediaType = "application/vnd.crucible.agent.stop-response.v0"              // [StopResponse].
	MediaTypeDestroyRequest           mediatype.MediaType = "application/vnd.crucible.agent.destroy-request.v0"            // [DestroyRequest].
	MediaTypeDestroyResponse          mediatype.MediaType = "application/vnd.crucible.agent.destroy-response.v0"           // [DestroyResponse].
	MediaTypeCopyRequest              mediatype.MediaType = "application/vnd.crucible.agent.copy-request.v0"               // [CopyRequest].
	MediaTypeCopyResponse             mediatype.MediaType = "application/vnd.crucible.agent.copy-response.v0"              // [CopyResponse].
	MediaTypeInspectContainerRequest  mediatype.MediaType = "application/vnd.crucible.agent.inspect-container-request.v0"  // [InspectContainerRequest].
	MediaTypeInspectContainerResponse mediatype.MediaType = "application/vnd.crucible.agent.inspect-container-response.v0" // [InspectContainerResponse].
	MediaTypeExecRequest              mediatype.MediaType = "application/vnd.crucible.agent.exec-request.v0"               // [ExecRequest].
	MediaTypeCommitRequest            mediatype.MediaType = "application/vnd.crucible.agent.commit-request.v0"             // [CommitRequest].
	MediaTypeCommitResponse           mediatype.MediaType = "application/vnd.crucible.agent.commit-response.v0"            // [CommitResponse].
	MediaTypeConfigureRequest         mediatype.MediaType = "application/vnd.crucible.agent.configure-request.v0"          // [ConfigureRequest].
	MediaTypeConfigureResponse        mediatype.MediaType = "application/vnd.crucible.agent.configure-response.v0"         // [ConfigureResponse].
	MediaTypeError                    mediatype.MediaType = "application/vnd.crucible.agent.error.v0"                      // [Error] returned on a non-2xx status.
)

// Media types for data parts and the exec stream, used as is.
const (
	MediaTypeImageArchive mediatype.MediaType = "application/vnd.crucible.agent.image-archive.v0" // OCI image tar archive (import, export).
	MediaTypeLayer        mediatype.MediaType = "application/vnd.crucible.agent.layer.v0"         // Uncompressed tar applied to a container (copy).
	MediaTypeFileTree     mediatype.MediaType = "application/vnd.crucible.agent.file-tree.v0"     // Tar stream extracted from an image (extract).
	MediaTypeExecResponse mediatype.MediaType = "application/vnd.crucible.agent.exec-response.v0" // Stream of [ExecResponse] messages, one JSON object per line (exec).
)
