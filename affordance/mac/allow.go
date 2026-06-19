package mac

import "github.com/cruciblehq/utils-go/crex"

// Single LSM hook allow rule produced by a MAC grant.
//
// Hook names the kernel LSM callback point at which the rule is evaluated.
// Where constrains the allow to invocations matching the predicate tree; a
// nil Where makes the rule unconditional.
type Allow struct {
	Hook  string `codec:"hook"`  // Kernel LSM hook name.
	Where *Expr  `codec:"where"` // Where-clause expression tree.
}

// Deep-clones an Allow rule.
func cloneAllow(a *Allow) *Allow {
	return &Allow{
		Hook:  a.Hook,
		Where: cloneExpr(a.Where),
	}
}

// Validates the allow rule.
func (a *Allow) Validate() error {
	if a.Hook == "" {
		return crex.Newf(ErrInvalidAllow, "hook is empty")
	}
	if a.Where != nil {
		if err := a.Where.Validate(); err != nil {
			return crex.Wrap(ErrInvalidAllow, err)
		}
	}
	return nil
}
