package registry

import (
	"context"

	"github.com/cruciblehq/spec/reference"
	"github.com/cruciblehq/spec/registry/cache"
	"github.com/cruciblehq/utils-go/crex"
)

// Configures the default registry and namespace used when resolving references.
//
// When a reference string omits the registry host or the namespace component,
// the corresponding default is applied by [Source.Parse] before the reference
// is passed to any pull or push operation.
type Source struct {
	Registry   string // Default registry URL.
	Namespace  string // Default namespace.
	Cache      string // Directory of the local archive cache.
	Credential string // Credential for Registry; empty for none.
}

// Creates a new [Source] with the given registry, namespace, and cache directory.
//
// All parameters are required. Returns an error wrapping [ErrMissingOption]
// if any is empty.
func NewSource(registry, namespace, cacheDir string) (Source, error) {
	if registry == "" {
		return Source{}, crex.ProgrammingError("missing registry", "a registry URL is required").
			Cause(crex.Wrap(ErrMissingOption, ErrMissingRegistry)).
			Err()
	}
	if namespace == "" {
		return Source{}, crex.ProgrammingError("missing namespace", "a namespace is required").
			Cause(crex.Wrap(ErrMissingOption, ErrMissingNamespace)).
			Err()
	}
	if cacheDir == "" {
		return Source{}, crex.ProgrammingError("missing cache directory", "a cache directory is required").
			Cause(crex.Wrap(ErrMissingOption, ErrMissingCache)).
			Err()
	}
	return Source{Registry: registry, Namespace: namespace, Cache: cacheDir}, nil
}

// Parses a resource reference string for the given type.
//
// Applies the default registry and namespace from [Source] when those
// components are absent from the reference string.
func (s Source) Parse(resourceType, ref string) (*reference.Reference, error) {
	parsed, err := reference.Parse(ref, resourceType)
	if err != nil {
		return nil, err
	}
	return parsed.WithDefaults(s.Registry, s.Namespace), nil
}

// Pulls a resource from the registry and extracts it locally.
//
// Checks the local cache first. Only downloads on cache miss or if the cached
// version is invalid or its digest does not match. Returns a [PullResult]
// containing the local extraction directory, digest, and version metadata.
// The credential is sent only when ref points at the source's registry.
func (s Source) Pull(ctx context.Context, ref *reference.Reference) (*PullResult, error) {
	credential := ""
	if ref.Registry() == s.Registry {
		credential = s.Credential
	}

	localCache, err := cache.Open(s.Cache)
	if err != nil {
		return nil, crex.Wrap(ErrCacheOperation, err)
	}
	defer localCache.Close()

	client := NewClient(ref.Registry(), credential, nil)
	return Pull(ctx, client, localCache, ref)
}
