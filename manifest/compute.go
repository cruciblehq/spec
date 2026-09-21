package manifest

import (
	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/crex"
)

// Identifies a compute provider.
//
// Each value names a provider whose backend the compute package resolves to
// provision and connect to a compute unit.
type ComputeType string

// Supported compute provider types.
const (
	ComputeTypeAWS   ComputeType = "aws"   // AWS EC2 provider.
	ComputeTypeLocal ComputeType = "local" // Local machine provider.
)

// Represents one compute unit in the infrastructure plan.
//
// A compute unit is an allocation of capacity for hosting services and can be
// backed by a variety of providers (AWS EC2, local machine, etc). The provider
// type determines the configuration fields required to provision and connect
// to the compute unit. The executor provisions the compute units and schedules
// services onto them based on the deployment specifications in the plan.
type Compute struct {

	// Provider type.
	//
	// Determines the type of [Compute.Config], which holds the configuration
	// fields specific to that provider. Supported providers are determined by
	// the compute package. Currently supported providers are [ComputeTypeAWS]
	// for AWS EC2 instances and [ComputeTypeLocal].
	Type ComputeType `codec:"type"`

	// Provider-specific configuration.
	//
	// The concrete type depends on [Compute.Type]: [*ComputeAWS] for "aws",
	// [*ComputeLocal] for "local".
	Config any `codec:"-"`

}

// Validates the compute entry.
//
// The provider type is not validated here since other packages determined the
// supported types. The provider-specific configuration (Compute.Config) must
// implement the Validatable interface and is validated by calling its Validate
// method. The configuration must also be present (non-nil), since the provider
// type implies required configuration fields. If kernel requirements are
// present they are also validated.
func (c *Compute) Validate() error {
	if c.Config == nil {
		return ErrMissingComputeProvider
	}
	v, ok := c.Config.(codec.Validatable)
	if !ok {
		return crex.Newf(ErrInvalidComputeType, "config type %T is not validatable", c.Config)
	}
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}

// Decodes a raw parsed map into the compute entry.
//
// Reads the provider type first, then decodes the remaining fields into the
// concrete configuration type for that provider.
func (c *Compute) Decode(cd *codec.Codec, raw any) error {
	src, ok := raw.(map[string]any)
	if !ok {
		return crex.Newf(ErrDecodeFailed, "unexpected compute type %T", raw)
	}
	if err := cd.Field(src, c, "Type"); err != nil {
		return crex.Wrap(ErrDecodeFailed, err)
	}
	var target any
	switch c.Type {
	case ComputeTypeAWS:
		target = &ComputeAWS{}
	case ComputeTypeLocal:
		target = &ComputeLocal{}
	default:
		return crex.Newf(ErrInvalidComputeType, "unknown provider %q", c.Type)
	}
	if err := cd.Decode(src, target); err != nil {
		return crex.Wrap(ErrDecodeFailed, err)
	}
	c.Config = target
	return nil
}

// Encodes the compute entry to a serializable map.
//
// Merges the provider configuration fields with the provider type key into a
// single flat map.
func (c *Compute) Encode(cd *codec.Codec) (any, error) {
	var cfg map[string]any
	if c.Config == nil {
		cfg = map[string]any{"type": c.Type}
	} else {
		var err error
		cfg, err = encodeToMap(cd, c.Config)
		if err != nil {
			return nil, err
		}
		cfg["type"] = c.Type
	}
	return cfg, nil
}
