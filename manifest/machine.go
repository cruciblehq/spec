package manifest

// Holds configuration specific to machine resources.
//
// Machine resources define bootable VM images for the Crucible ecosystem.
// They package a base VM image with additional customisation (installing
// packages, copying configuration files, enabling services, etc.) to produce
// a node image that providers can deploy. The build pipeline is the same
// [Recipe] used by runtimes and services; the difference is that the output
// artifact targets VM image formats (QCOW2, AMI, etc.) rather than OCI
// container images.
type Machine struct {
	Recipe `yaml:",inline"`
}

// Validates the machine configuration.
func (m *Machine) Validate() error {
	return m.Recipe.Validate()
}
