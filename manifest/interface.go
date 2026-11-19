package manifest

import (
	"github.com/cruciblehq/spec/reference"
	"github.com/cruciblehq/utils-go/crex"
)

// Defines an interface resource.
//
// Interfaces are symbolic resources that describe the expected behavior of an
// API. Services can declare they implement an interface and can also declare a
// dependency on an interface rather than a specific service.
type Interface struct {
}

// Validates the interface configuration.
//
// An interface has no fields, so it is always valid.
func (i *Interface) Validate() error {
	return nil
}

// Whether the given implementation reference satisfies the required interface.
//
// Both must name the same interface. Returns an error when either reference
// is invalid or has no version, or when the implementation version is not
// exact.
func InterfaceReferenceMatches(requirement, implementation string) (bool, error) {
	reqRef, err := reference.Parse(requirement, string(TypeInterface))
	if err != nil {
		return false, crex.Wrap(ErrInvalidRef, err)
	}
	implRef, err := reference.Parse(implementation, string(TypeInterface))
	if err != nil {
		return false, crex.Wrap(ErrInvalidRef, err)
	}

	if reqRef.Identifier.Path() != implRef.Identifier.Path() {
		return false, nil
	}
	if reqRef.Version() == nil || implRef.Version() == nil {
		return false, crex.Wrap(ErrInvalidRef, ErrMissingRefTarget)
	}

	implVersion, ok := implRef.Version().ExactVersion()
	if !ok {
		return false, crex.Wrap(ErrInvalidRef, ErrMissingRefTarget)
	}

	return reqRef.Version().MatchesVersion(implVersion)
}
