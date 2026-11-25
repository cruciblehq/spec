package registry

import (
	"context"
	"io"
)

// Interface for artifact registry operations.
//
// Provides hierarchical storage and retrieval of versioned artifacts organized
// into namespaces and resources. Supports private (changeable) and public
// (final) versions, and compressed archive distribution. Every operation takes
// a context for cancellation and timeouts.
type Registry interface {

	// Creates a new namespace.
	//
	// Namespace names may include lowercase letters (a–z), digits (0–9), and
	// hyphens (-), must start and end with an alphanumeric character, and must
	// not exceed 63 characters. The Description field is optional and may be
	// empty. If a namespace with the given name already exists, the registry
	// returns an error. The response includes the created namespace's metadata
	// with an empty resources list.
	CreateNamespace(ctx context.Context, info NamespaceInfo) (*Namespace, error)

	// Retrieves namespace metadata and resource summaries.
	//
	// Returns namespace information along with lightweight summaries of all
	// contained resources. Summaries include basic metadata (such as latest
	// versions) but exclude full version histories. If the namespace does not
	// exist, an error is returned.
	ReadNamespace(ctx context.Context, namespace string) (*Namespace, error)

	// Updates mutable namespace metadata.
	//
	// Immutable identifiers cannot be changed. Updating metadata does not affect
	// contained resources or their timestamps. If the namespace does not exist,
	// the operation fails.
	UpdateNamespace(ctx context.Context, namespace string, info NamespaceInfo) (*Namespace, error)

	// Permanently deletes a namespace.
	//
	// Namespaces cannot be deleted if they contain any resources. The operation
	// is idempotent, returning success if the namespace does not exist.
	DeleteNamespace(ctx context.Context, namespace string) error

	// Lists all namespaces.
	//
	// Returns a list of all existing namespaces and an empty list if none
	// exist. The list order is up to the implementation.
	ListNamespaces(ctx context.Context) (*NamespaceList, error)

	// Creates a new resource.
	//
	// Resource names follow the same constraints as namespace names. If a
	// resource with the given name already exists in the namespace, an error
	// is returned. The response includes the created resource's metadata with
	// an empty version list.
	CreateResource(ctx context.Context, namespace string, info ResourceInfo) (*Resource, error)

	// Retrieves resource metadata with version summaries.
	//
	// Returns resource information along with lightweight summaries of all
	// versions. Summaries exclude full archive details. If the
	// namespace or resource does not exist, an error is returned.
	ReadResource(ctx context.Context, namespace, resource string) (*Resource, error)

	// Updates mutable resource metadata.
	//
	// Immutable identifiers cannot be changed. If the namespace or resource does
	// not exist, the operation fails.
	UpdateResource(ctx context.Context, namespace, resource string, info ResourceInfo) (*Resource, error)

	// Permanently deletes a resource.
	//
	// Resources cannot be deleted if they contain any published versions. The
	// operation is idempotent, returning success if the resource does not exist.
	DeleteResource(ctx context.Context, namespace, resource string) error

	// Lists all resources in a namespace.
	//
	// Returns a list of resource summaries including statistics and latest
	// versions. The order is up to the implementation, and the list is empty
	// if the namespace contains no resources. If the namespace does not exist,
	// an error is returned.
	ListResources(ctx context.Context, namespace string) (*ResourceList, error)

	// Creates a new version.
	//
	// If a version with the given string already exists, an error is returned.
	// The response includes the created version's metadata. Archive fields
	// remain null until an archive is uploaded. Versions are created private.
	// If a public version with the same major and minor is higher, the
	// operation fails with [ErrorCodeVersionSuperseded].
	CreateVersion(ctx context.Context, namespace, resource string, info VersionInfo) (*Version, error)

	// Retrieves version metadata with archive details.
	//
	// Returns complete version information including archive URL, size, and
	// digest. If the namespace, resource, or version does not exist, an error
	// is returned.
	ReadVersion(ctx context.Context, namespace, resource, version string) (*Version, error)

	// Updates a version.
	//
	// Only private versions can be updated; a public version fails with
	// [ErrorCodeVersionPublic]. Setting Public makes the version public, which
	// requires an uploaded archive, a version that is not a prerelease, and no
	// higher public version with the same major and minor
	// ([ErrorCodeVersionSuperseded]).
	UpdateVersion(ctx context.Context, namespace, resource, version string, info VersionInfo) (*Version, error)

	// Permanently deletes a version.
	//
	// The operation is idempotent, returning success if the version does not
	// exist.
	DeleteVersion(ctx context.Context, namespace, resource, version string) error

	// Lists all versions for a resource.
	//
	// Returns a list of version summaries including timestamps. The list is
	// empty if the resource has no versions. If the namespace or resource does
	// not exist, an error is returned.
	ListVersions(ctx context.Context, namespace, resource string) (*VersionList, error)

	// Uploads a version archive.
	//
	// Uploads the archive data for a version. A private version's archive can
	// be replaced by uploading again; a public version fails with
	// [ErrorCodeVersionPublic]. The digest is calculated from the archive data
	// using SHA-256 for integrity verification. Returns the updated version
	// with populated archive metadata.
	UploadArchive(ctx context.Context, namespace, resource, version string, archive io.Reader) (*Version, error)

	// Downloads a version archive.
	//
	// Returns a reader for the archive data. The caller is responsible for
	// closing the reader. If the namespace, resource, or version does not
	// exist, or if the version has no uploaded archive, an error is returned.
	DownloadArchive(ctx context.Context, namespace, resource, version string) (io.ReadCloser, error)
}
