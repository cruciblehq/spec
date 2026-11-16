package registry

import (
	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/mediatype"
)

// Encoding of message bodies in the registry API.
//
// Bodies carry their media type with this format's suffix (see
// [mediatype.MediaType.With]), except [MediaTypeArchive], which is binary.
const BodyFormat = codec.JSON

// Media types of the Crucible registry API.
//
// Each follows the pattern application/vnd.crucible.{name}.v0. They are used
// in Content-Type headers for request bodies and Accept headers for response
// format negotiation.
const (
	MediaTypeError         mediatype.MediaType = "application/vnd.crucible.error.v0"          // Error responses with codes and messages.
	MediaTypeNamespaceInfo mediatype.MediaType = "application/vnd.crucible.namespace-info.v0" // Namespace create/update requests.
	MediaTypeNamespace     mediatype.MediaType = "application/vnd.crucible.namespace.v0"      // Complete namespace with resource summaries.
	MediaTypeNamespaceList mediatype.MediaType = "application/vnd.crucible.namespace-list.v0" // Collection of namespace summaries.
	MediaTypeResourceInfo  mediatype.MediaType = "application/vnd.crucible.resource-info.v0"  // Resource create/update requests.
	MediaTypeResource      mediatype.MediaType = "application/vnd.crucible.resource.v0"       // Complete resource with version/channel summaries.
	MediaTypeResourceList  mediatype.MediaType = "application/vnd.crucible.resource-list.v0"  // Collection of resource summaries.
	MediaTypeVersionInfo   mediatype.MediaType = "application/vnd.crucible.version-info.v0"   // Version create/update requests.
	MediaTypeVersion       mediatype.MediaType = "application/vnd.crucible.version.v0"        // Complete version with archive details.
	MediaTypeVersionList   mediatype.MediaType = "application/vnd.crucible.version-list.v0"   // Collection of version summaries.
	MediaTypeChannelInfo   mediatype.MediaType = "application/vnd.crucible.channel-info.v0"   // Channel create/update requests.
	MediaTypeChannel       mediatype.MediaType = "application/vnd.crucible.channel.v0"        // Complete channel with full version object.
	MediaTypeChannelList   mediatype.MediaType = "application/vnd.crucible.channel-list.v0"   // Collection of channel summaries.
	MediaTypeArchive       mediatype.MediaType = "application/vnd.crucible.archive.v0"        // Binary archive data (tar.zst format).
)
