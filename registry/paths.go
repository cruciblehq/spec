package registry

import (
	"net/url"
	"strings"
)

// URL path patterns of the registry API.
//
// Each pattern names its variable segments in braces, in the form
// [net/http.ServeMux] accepts. Servers route on them, and clients build
// concrete paths from them with the path functions below.
const (
	PatternNamespaceList = "/namespaces"                                                             // Namespace collection.
	PatternNamespace     = "/namespaces/{namespace}"                                                 // One namespace.
	PatternResourceList  = "/namespaces/{namespace}/resources"                                       // Resources of a namespace.
	PatternResource      = "/namespaces/{namespace}/resources/{resource}"                            // One resource.
	PatternVersionList   = "/namespaces/{namespace}/resources/{resource}/versions"                   // Versions of a resource.
	PatternVersion       = "/namespaces/{namespace}/resources/{resource}/versions/{version}"         // One version.
	PatternArchive       = "/namespaces/{namespace}/resources/{resource}/versions/{version}/archive" // Archive of a version.
)

// Names of the variable segments in the path patterns.
const (
	ParamNamespace = "namespace" // Namespace name.
	ParamResource  = "resource"  // Resource name.
	ParamVersion   = "version"   // Version string.
)

// Returns the path of the namespace collection.
func NamespaceListPath() string {
	return PatternNamespaceList
}

// Returns the path of namespace ns.
func NamespacePath(ns string) string {
	return expand(PatternNamespace, ns)
}

// Returns the path of the resources in namespace ns.
func ResourceListPath(ns string) string {
	return expand(PatternResourceList, ns)
}

// Returns the path of resource res in namespace ns.
func ResourcePath(ns, res string) string {
	return expand(PatternResource, ns, res)
}

// Returns the path of the versions of resource res.
func VersionListPath(ns, res string) string {
	return expand(PatternVersionList, ns, res)
}

// Returns the path of version ver of resource res.
func VersionPath(ns, res, ver string) string {
	return expand(PatternVersion, ns, res, ver)
}

// Returns the path of the archive of version ver.
func ArchivePath(ns, res, ver string) string {
	return expand(PatternArchive, ns, res, ver)
}

// Replaces the variable segments of pattern with values, in order.
//
// Each value is escaped as a single path segment.
func expand(pattern string, values ...string) string {
	segments := strings.Split(pattern, "/")
	next := 0
	for i, s := range segments {
		if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
			segments[i] = url.PathEscape(values[next])
			next++
		}
	}
	return strings.Join(segments, "/")
}
