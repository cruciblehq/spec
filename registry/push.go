package registry

import (
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/cruciblehq/spec/reference"
	"github.com/cruciblehq/spec/registry/cache"
	"github.com/cruciblehq/utils-go/crex"
)

// Pushes a resource package to the source's registry.
//
// Resolves the resource identifier, checks that the namespace exists, creates
// the resource and the version if needed, then uploads archive and adds it to
// the source's local cache. Pushing a private version again replaces its
// archive. Returns [ErrVersionPublic] if the version is public, or [ErrVersionSuperseded]
// if a public version with the same major and minor is higher.
func (s Source) Push(ctx context.Context, name, resourceType, version string, archive io.ReadSeeker) error {
	id, err := reference.ParseIdentifier(name, resourceType)
	if err != nil {
		return crex.Wrap(ErrInvalidIdentifier, err)
	}

	client := NewClient(s.Registry, s.Credential, nil)

	if err := verifyNamespace(ctx, client, id.Namespace()); err != nil {
		return err
	}

	if err := ensureResource(ctx, client, id.Namespace(), id.Name(), resourceType); err != nil {
		return err
	}

	if err := ensureVersion(ctx, client, id.Namespace(), id.Name(), version); err != nil {
		return err
	}

	return uploadPackage(ctx, client, s.Cache, id.Namespace(), id.Name(), version, archive)
}

// Checks that the namespace exists in the registry.
//
// Returns [ErrNamespaceNotFound] on a 404 response, or [ErrRegistryOperation]
// for other failures.
func verifyNamespace(ctx context.Context, client *Client, namespace string) error {
	_, err := client.ReadNamespace(ctx, namespace)
	if err == nil {
		return nil
	}

	var regErr *Error
	if errors.As(err, &regErr) && regErr.Code == ErrorCodeNotFound {
		return crex.Wrap(ErrNamespaceNotFound, err)
	}

	return crex.Wrap(ErrRegistryOperation, err)
}

// Ensures the resource exists, creating it if necessary.
//
// Looks up the resource and creates it with the given type if it does not
// yet exist.
func ensureResource(ctx context.Context, client *Client, namespace, resource, resourceType string) error {
	_, err := client.ReadResource(ctx, namespace, resource)
	if err == nil {
		return nil
	}

	var regErr *Error
	if !errors.As(err, &regErr) || regErr.Code != ErrorCodeNotFound {
		return crex.Wrap(ErrRegistryOperation, err)
	}

	resInfo := ResourceInfo{
		Name: resource,
		Type: resourceType,
	}
	_, err = client.CreateResource(ctx, namespace, resInfo)
	if err != nil {
		return crex.Wrap(ErrRegistryOperation, err)
	}

	return nil
}

// Ensures the version exists, creating it if necessary.
//
// Returns [ErrVersionSuperseded] if a public version with the same major and
// minor is higher.
func ensureVersion(ctx context.Context, client *Client, namespace, resource, version string) error {
	versionInfo := VersionInfo{
		String: version,
	}

	_, err := client.CreateVersion(ctx, namespace, resource, versionInfo)
	var regErr *Error
	switch {
	case err == nil:
		return nil
	case errors.As(err, &regErr) && regErr.Code == ErrorCodeVersionExists:
		return nil
	case errors.As(err, &regErr) && regErr.Code == ErrorCodeVersionSuperseded:
		return crex.Wrap(ErrVersionSuperseded, err)
	default:
		return crex.Wrap(ErrRegistryOperation, err)
	}
}

// Uploads the package archive to the registry.
//
// Uploads archive to the specified resource version in the registry. Returns
// [ErrVersionPublic] if the version is public. After a successful upload,
// rewinds archive and adds it to the local cache in cacheDir.
func uploadPackage(ctx context.Context, client *Client, cacheDir, namespace, resource, version string, archive io.ReadSeeker) error {
	uploaded, err := client.UploadArchive(ctx, namespace, resource, version, archive)
	var regErr *Error
	if errors.As(err, &regErr) && regErr.Code == ErrorCodeVersionPublic {
		return crex.Wrap(ErrVersionPublic, err)
	}
	if err != nil {
		return crex.Wrap(ErrRegistryOperation, err)
	}

	if err := updateLocalCache(cacheDir, namespace, resource, version, archive, uploaded.Digest); err != nil {
		// Cache errors are just logged, since the upload succeeded.
		slog.Error("failed to update local cache", "error", err)
	}

	return nil
}

// Adds the pushed package to the local cache in cacheDir.
//
// Reads archive from its start and stores it under the given
// namespace/resource/version with digest verification, keeping the local
// cache in sync with the remote so a subsequent pull does not download it again.
func updateLocalCache(cacheDir, namespace, resource, version string, archive io.ReadSeeker, digest *string) error {
	if _, err := archive.Seek(0, io.SeekStart); err != nil {
		return crex.Wrap(ErrFileSystemOperation, err)
	}

	localCache, err := cache.Open(cacheDir)
	if err != nil {
		return crex.Wrap(ErrCacheOperation, err)
	}
	defer localCache.Close()

	if digest != nil && *digest != "" {
		_, err = localCache.PutWithDigest(namespace, resource, version, *digest, archive)
	} else {
		_, err = localCache.Put(namespace, resource, version, archive)
	}
	if err != nil {
		return crex.Wrap(ErrCacheOperation, err)
	}
	return nil
}
