package mac

import (
	"github.com/cruciblehq/spec/affordance/agl"
	"github.com/cruciblehq/utils-go/crex"
)

// Normalised expression node in the MAC subsystem model.
//
// Type selects the expression variant and determines which fields are populated.
// For binary nodes ("and", "or") Left and Right are set; for unary ("not") only
// Operand; for comparisons ("cmp") Op, LHS, and RHS; and so on for "in", "like",
// "between", and "bittest".
type Expr struct {
	Type    string   `codec:"type"`    // Expression type discriminator.
	Left    *Expr    `codec:"left"`    // Left child for binary expressions.
	Right   *Expr    `codec:"right"`   // Right child for binary expressions.
	Operand *Expr    `codec:"operand"` // Operand for unary expressions.
	Op      string   `codec:"op"`      // Comparison operator.
	LHS     *Value   `codec:"lhs"`     // Left-hand side value for comparison expressions.
	RHS     *Value   `codec:"rhs"`     // Right-hand side value for comparison expressions.
	Field   *Value   `codec:"field"`   // Field reference for field-based expressions.
	Values  []*Value `codec:"values"`  // Values for multi-value expressions such as "in".
	Pattern string   `codec:"pattern"` // Pattern for pattern-matching expressions.
	Low     *Value   `codec:"low"`     // Lower bound for between expressions.
	High    *Value   `codec:"high"`    // Upper bound for between expressions.
	Mask    *Value   `codec:"mask"`    // Mask value for bit-test expressions.
}

// Expression type discriminators for the Type field of an Expr.
//
// The boolean operators reuse the AGL operator constants; the remaining
// discriminators are specific to the normalised MAC model.
const (
	exprAnd     = string(agl.OpAnd) // Logical conjunction of Left and Right.
	exprOr      = string(agl.OpOr)  // Logical disjunction of Left and Right.
	exprNot     = string(agl.OpNot) // Logical negation of Operand.
	exprCmp     = "cmp"             // Comparison of LHS and RHS with Op.
	exprIn      = "in"              // Membership test of Field against Values.
	exprLike    = "like"            // Pattern match of Field against Pattern.
	exprBetween = "between"         // Range test of Field within Low and High.
	exprBitTest = "bittest"         // Bit-mask test of Field against Mask.
)

// Deep-clones an Expr.
func cloneExpr(e *Expr) *Expr {
	if e == nil {
		return nil
	}
	vals := make([]*Value, len(e.Values))
	for i, v := range e.Values {
		vals[i] = cloneValue(v)
	}
	return &Expr{
		Type:    e.Type,
		Left:    cloneExpr(e.Left),
		Right:   cloneExpr(e.Right),
		Operand: cloneExpr(e.Operand),
		Op:      e.Op,
		LHS:     cloneValue(e.LHS),
		RHS:     cloneValue(e.RHS),
		Field:   cloneValue(e.Field),
		Values:  vals,
		Pattern: e.Pattern,
		Low:     cloneValue(e.Low),
		High:    cloneValue(e.High),
		Mask:    cloneValue(e.Mask),
	}
}

// Whether two Expr trees are structurally equal.
func exprEqual(a, b *Expr) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Type != b.Type || a.Op != b.Op || a.Pattern != b.Pattern {
		return false
	}
	if !valueEqual(a.LHS, b.LHS) || !valueEqual(a.RHS, b.RHS) {
		return false
	}
	if !valueEqual(a.Field, b.Field) || !valueEqual(a.Mask, b.Mask) {
		return false
	}
	if !valueEqual(a.Low, b.Low) || !valueEqual(a.High, b.High) {
		return false
	}
	return exprEqual(a.Left, b.Left) &&
		exprEqual(a.Right, b.Right) &&
		exprEqual(a.Operand, b.Operand) &&
		valuesEqual(a.Values, b.Values)
}

// Validates the expression tree node.
//
// Type must be one of the known expression types and the required operands for
// that type must be present. Child nodes are validated recursively.
func (e *Expr) Validate() error {
	switch e.Type {
	case exprAnd, exprOr:
		return e.validateBinary()
	case exprNot:
		return e.validateUnary()
	case exprCmp:
		return e.validateComparison()
	case exprIn:
		return e.validateIn()
	case exprLike:
		return e.validateLike()
	case exprBetween:
		return e.validateBetween()
	case exprBitTest:
		return e.validateBitTest()
	default:
		return crex.Newf(ErrInvalidExpr, "unknown type %q", e.Type)
	}
}

// Validates each provided child expression node, wrapping any failure.
func validateExprs(children ...*Expr) error {
	for _, c := range children {
		if err := c.Validate(); err != nil {
			return crex.Wrap(ErrInvalidExpr, err)
		}
	}
	return nil
}

// Validates each provided value, wrapping any failure.
func validateValues(values ...*Value) error {
	for _, v := range values {
		if err := v.Validate(); err != nil {
			return crex.Wrap(ErrInvalidExpr, err)
		}
	}
	return nil
}

// Validates a binary node, requiring both child operands.
func (e *Expr) validateBinary() error {
	if e.Left == nil || e.Right == nil {
		return crex.Newf(ErrInvalidExpr, "%q requires left and right operands", e.Type)
	}
	return validateExprs(e.Left, e.Right)
}

// Validates a unary node, requiring its single operand.
func (e *Expr) validateUnary() error {
	if e.Operand == nil {
		return crex.Newf(ErrInvalidExpr, "%q requires an operand", e.Type)
	}
	return validateExprs(e.Operand)
}

// Validates a comparison node, requiring an operator and both sides.
func (e *Expr) validateComparison() error {
	if e.Op == "" || e.LHS == nil || e.RHS == nil {
		return crex.Newf(ErrInvalidExpr, "%q requires op, lhs, and rhs", e.Type)
	}
	return validateValues(e.LHS, e.RHS)
}

// Validates a membership node, requiring a field and non-nil values.
func (e *Expr) validateIn() error {
	if e.Field == nil || len(e.Values) == 0 {
		return crex.Newf(ErrInvalidExpr, "%q requires a field and at least one value", e.Type)
	}
	if err := validateValues(e.Field); err != nil {
		return err
	}
	for i, v := range e.Values {
		if v == nil {
			return crex.At(crex.Newf(ErrInvalidExpr, "%q has a nil value", e.Type), "value", i+1)
		}
		if err := validateValues(v); err != nil {
			return err
		}
	}
	return nil
}

// Validates a pattern-match node, requiring a field and pattern.
func (e *Expr) validateLike() error {
	if e.Field == nil || e.Pattern == "" {
		return crex.Newf(ErrInvalidExpr, "%q requires a field and pattern", e.Type)
	}
	return validateValues(e.Field)
}

// Validates a range node, requiring a field and both bounds.
func (e *Expr) validateBetween() error {
	if e.Field == nil || e.Low == nil || e.High == nil {
		return crex.Newf(ErrInvalidExpr, "%q requires field, low, and high", e.Type)
	}
	return validateValues(e.Field, e.Low, e.High)
}

// Validates a bit-test node, requiring a field and mask.
func (e *Expr) validateBitTest() error {
	if e.Field == nil || e.Mask == nil {
		return crex.Newf(ErrInvalidExpr, "%q requires field and mask", e.Type)
	}
	return validateValues(e.Field, e.Mask)
}
