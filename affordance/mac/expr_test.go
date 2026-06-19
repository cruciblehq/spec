package mac

import (
	"errors"
	"testing"
)

func TestMACExprValidateAndMissingChildren(t *testing.T) {
	e := &Expr{Type: "and"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateAndInvalidLeft(t *testing.T) {
	e := &Expr{
		Type:  "and",
		Left:  &Expr{Type: "bogus"},
		Right: &Expr{Type: "cmp", Op: "=", LHS: &Value{IsField: true, Field: "task.uid"}, RHS: &Value{}},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateOrInvalidRight(t *testing.T) {
	e := &Expr{
		Type:  "or",
		Left:  &Expr{Type: "cmp", Op: "=", LHS: &Value{IsField: true, Field: "task.uid"}, RHS: &Value{}},
		Right: &Expr{Type: "bogus"},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateNotMissingOperand(t *testing.T) {
	e := &Expr{Type: "not"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateNotInvalidOperand(t *testing.T) {
	e := &Expr{Type: "not", Operand: &Expr{Type: "bogus"}}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateCmpMissing(t *testing.T) {
	e := &Expr{Type: "cmp"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateCmpInvalidLHS(t *testing.T) {
	e := &Expr{
		Type: "cmp",
		Op:   "=",
		LHS:  &Value{IsField: true, Field: ""},
		RHS:  &Value{},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateCmpInvalidRHS(t *testing.T) {
	e := &Expr{
		Type: "cmp",
		Op:   "=",
		LHS:  &Value{IsField: true, Field: "task.uid"},
		RHS:  &Value{IsField: true, Field: ""},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateCmpOK(t *testing.T) {
	e := &Expr{
		Type: "cmp",
		Op:   "=",
		LHS:  &Value{IsField: true, Field: "task.uid"},
		RHS:  &Value{IntVal: 0},
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACExprValidateInMissing(t *testing.T) {
	e := &Expr{Type: "in"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateInInvalidField(t *testing.T) {
	e := &Expr{
		Type:   "in",
		Field:  &Value{IsField: true, Field: ""},
		Values: []*Value{{IntVal: 1}},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateInNilValue(t *testing.T) {
	e := &Expr{
		Type:   "in",
		Field:  &Value{IsField: true, Field: "task.uid"},
		Values: []*Value{nil},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateInInvalidValue(t *testing.T) {
	e := &Expr{
		Type:   "in",
		Field:  &Value{IsField: true, Field: "task.uid"},
		Values: []*Value{{IsField: true, Field: ""}},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateLikeMissing(t *testing.T) {
	e := &Expr{Type: "like"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateLikeInvalidField(t *testing.T) {
	e := &Expr{
		Type:    "like",
		Field:   &Value{IsField: true, Field: ""},
		Pattern: "/tmp/*",
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateLikeOK(t *testing.T) {
	e := &Expr{
		Type:    "like",
		Field:   &Value{IsField: true, Field: "file.path"},
		Pattern: "/tmp/*",
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACExprValidateBetweenMissing(t *testing.T) {
	e := &Expr{Type: "between"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBetweenInvalidField(t *testing.T) {
	e := &Expr{
		Type:  "between",
		Field: &Value{IsField: true, Field: ""},
		Low:   &Value{IntVal: 0},
		High:  &Value{IntVal: 100},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBetweenInvalidLow(t *testing.T) {
	e := &Expr{
		Type:  "between",
		Field: &Value{IsField: true, Field: "task.uid"},
		Low:   &Value{IsField: true, Field: ""},
		High:  &Value{IntVal: 100},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBetweenInvalidHigh(t *testing.T) {
	e := &Expr{
		Type:  "between",
		Field: &Value{IsField: true, Field: "task.uid"},
		Low:   &Value{IntVal: 0},
		High:  &Value{IsField: true, Field: ""},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBetweenOK(t *testing.T) {
	e := &Expr{
		Type:  "between",
		Field: &Value{IsField: true, Field: "task.uid"},
		Low:   &Value{IntVal: 0},
		High:  &Value{IntVal: 100},
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACExprValidateBittestMissing(t *testing.T) {
	e := &Expr{Type: "bittest"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBittestInvalidField(t *testing.T) {
	e := &Expr{
		Type:  "bittest",
		Field: &Value{IsField: true, Field: ""},
		Mask:  &Value{IntVal: 7},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBittestInvalidMask(t *testing.T) {
	e := &Expr{
		Type:  "bittest",
		Field: &Value{IsField: true, Field: "task.uid"},
		Mask:  &Value{IsField: true, Field: ""},
	}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}

func TestMACExprValidateBittestOK(t *testing.T) {
	e := &Expr{
		Type:  "bittest",
		Field: &Value{IsField: true, Field: "task.uid"},
		Mask:  &Value{IntVal: 7},
	}
	if err := e.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMACExprValidateUnknownType(t *testing.T) {
	e := &Expr{Type: "unknown"}
	if err := e.Validate(); !errors.Is(err, ErrInvalidExpr) {
		t.Fatalf("err = %v, want ErrInvalidExpr", err)
	}
}
