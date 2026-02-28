package manifest

import (
	"github.com/cruciblehq/crex"
	"gopkg.in/yaml.v3"
)

// Encodes a manifest to YAML.
//
// The type-specific configuration in [Manifest.Config] is marshaled into the
// same YAML mapping as the common fields so the output matches the canonical
// manifest format. The manifest is validated before encoding.
func Encode(m *Manifest) ([]byte, error) {
	if err := m.Validate(); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}

	base := struct {
		Version  int      `yaml:"version"`
		Resource Resource `yaml:"resource"`
	}{
		Version:  m.Version,
		Resource: m.Resource,
	}

	// Marshal base fields and config separately, then merge the YAML nodes
	// so type-specific fields appear at the top level.
	var baseNode, cfgNode yaml.Node
	if err := baseNode.Encode(base); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	if err := cfgNode.Encode(m.Config); err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}

	// Both are MappingNodes; append config key-value pairs to the base.
	if cfgNode.Kind == yaml.MappingNode {
		baseNode.Content = append(baseNode.Content, cfgNode.Content...)
	}

	out, err := yaml.Marshal(&baseNode)
	if err != nil {
		return nil, crex.Wrap(ErrEncodeFailed, err)
	}
	return out, nil
}

// Decodes a YAML document into a manifest.
//
// Decoding is performed in two phases. The common fields (version, resource)
// are unmarshaled first to determine the resource type. The raw YAML is then
// unmarshaled again into the concrete configuration type corresponding to
// that resource type. The decoded manifest is validated before returning.
func Decode(data []byte) (*Manifest, error) {
	var m Manifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	configs := map[ResourceType]any{
		TypeRuntime:  &Runtime{},
		TypeService:  &Service{},
		TypeWidget:   &Widget{},
		TypeTemplate: &Template{},
		TypeMachine:  &Machine{},
	}

	target, ok := configs[m.Resource.Type]
	if !ok {
		return nil, crex.Wrap(ErrDecodeFailed, ErrInvalidResourceType)
	}

	if err := yaml.Unmarshal(data, target); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	m.Config = target

	if err := m.Validate(); err != nil {
		return nil, crex.Wrap(ErrDecodeFailed, err)
	}

	return &m, nil
}
