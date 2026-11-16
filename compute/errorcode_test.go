package compute

import (
	"net/http"
	"testing"
)

func TestIsValidErrorCode(t *testing.T) {
	valid := []ErrorCode{
		ErrorCodeBadRequest,
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

func TestErrorCodeHTTPStatus(t *testing.T) {
	cases := []struct {
		code ErrorCode
		want int
	}{
		{ErrorCodeBadRequest, http.StatusBadRequest},
		{ErrorCodeNotFound, http.StatusNotFound},
		{ErrorCodeConflict, http.StatusConflict},
		{ErrorCodePreconditionFailed, http.StatusPreconditionFailed},
		{ErrorCodeInternalError, http.StatusInternalServerError},
		{"unknown", http.StatusInternalServerError},
	}
	for _, c := range cases {
		if got := c.code.HTTPStatus(); got != c.want {
			t.Errorf("ErrorCode(%q).HTTPStatus() = %d, want %d", c.code, got, c.want)
		}
	}
}
