package registry

// Classification of a registry error.
//
// Provides granular error information beyond HTTP status codes, enabling
// clients to implement specific error handling logic. Used in the Code field
// of [Error] responses.
type ErrorCode string

// Known error codes.
const (
	ErrorCodeBadRequest           ErrorCode = "bad_request"                     // Request validation failed (malformed body, invalid fields).
	ErrorCodeUnauthorized         ErrorCode = "unauthorized"                    // Request has no valid credential.
	ErrorCodeNotFound             ErrorCode = "not_found"                       // Requested namespace, resource, or version does not exist.
	ErrorCodeNamespaceExists      ErrorCode = "namespace_exists"                // Cannot create namespace: name already in use.
	ErrorCodeNamespaceNotEmpty    ErrorCode = "namespace_not_empty"             // Cannot delete namespace: contains resources.
	ErrorCodeResourceExists       ErrorCode = "resource_exists"                 // Cannot create resource: name already in use within namespace.
	ErrorCodeResourceHasPublished ErrorCode = "resource_has_published_versions" // Cannot delete resource: contains versions published on the marketplace.
	ErrorCodeVersionExists        ErrorCode = "version_exists"                  // Cannot create version: version string already in use.
	ErrorCodeVersionPublic        ErrorCode = "version_public"                  // Cannot change version: it is public and final.
	ErrorCodeVersionSuperseded    ErrorCode = "version_superseded"              // Cannot push or make public: a public version with the same major and minor is higher.
	ErrorCodeVersionPublished     ErrorCode = "version_published"               // Cannot delete version: published on the marketplace.
	ErrorCodePreconditionFailed   ErrorCode = "precondition_failed"             // Request precondition not met (e.g., If-Match header mismatch).
	ErrorCodeInternalError        ErrorCode = "internal_error"                  // Unexpected server error occurred.
)

// Set of known error codes for validation.
var validErrorCodes = map[ErrorCode]bool{
	ErrorCodeBadRequest:           true,
	ErrorCodeUnauthorized:         true,
	ErrorCodeNotFound:             true,
	ErrorCodeNamespaceExists:      true,
	ErrorCodeNamespaceNotEmpty:    true,
	ErrorCodeResourceExists:       true,
	ErrorCodeResourceHasPublished: true,
	ErrorCodeVersionExists:        true,
	ErrorCodeVersionPublic:        true,
	ErrorCodeVersionSuperseded:    true,
	ErrorCodeVersionPublished:     true,
	ErrorCodePreconditionFailed:   true,
	ErrorCodeInternalError:        true,
}

// Whether the error code is a known value.
func isValidErrorCode(code ErrorCode) bool {
	return validErrorCodes[code]
}
