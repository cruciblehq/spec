package compute

import "testing"

func TestIsValidErrorCode(t *testing.T) {
	valid := []ErrorCode{
		ErrorCodeBadRequest,
		ErrorCodeUnauthenticated,
		ErrorCodeNotFound,
		ErrorCodeConflict,
		ErrorCodePreconditionFailed,
		ErrorCodeInternalError,
	}
	for _, code := range valid {
		if !isValidErrorCode(code) {
			t.Errorf("isValidErrorCode(%q) = false, want true", code)
		}
	}

	for _, code := range []ErrorCode{"", "unknown", "not_found "} {
		if isValidErrorCode(code) {
			t.Errorf("isValidErrorCode(%q) = true, want false", code)
		}
	}
}
