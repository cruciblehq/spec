package manifest

import "testing"

func TestIsValidName(t *testing.T) {
	valid := []string{"foo", "bar_baz", "x1", "a", "a_1_b"}
	invalid := []string{"", "Foo", "1foo", "_foo", "foo-bar", "foo bar", "foo.bar"}
	for _, s := range valid {
		if !isValidName(s) {
			t.Errorf("isValidName(%q) = false, want true", s)
		}
	}
	for _, s := range invalid {
		if isValidName(s) {
			t.Errorf("isValidName(%q) = true, want false", s)
		}
	}
}

func TestIsValidPlatform(t *testing.T) {
	valid := []string{"linux/amd64", "linux/arm64", "linux/arm64/v8"}
	invalid := []string{"", "linux", "/amd64", "linux/", "Linux/amd64", "linux//amd64"}
	for _, s := range valid {
		if !isValidPlatform(s) {
			t.Errorf("isValidPlatform(%q) = false, want true", s)
		}
	}
	for _, s := range invalid {
		if isValidPlatform(s) {
			t.Errorf("isValidPlatform(%q) = true, want false", s)
		}
	}
}

func TestIsValidRoutePattern(t *testing.T) {
	valid := []string{"/", "/api", "/auth/login", "/a-b_c.d~e", "/foo/bar"}
	invalid := []string{"", "api", "/foo//bar", "//", "/foo?x=1", "/foo#f", "foo/bar"}
	for _, s := range valid {
		if !isValidRoutePattern(s) {
			t.Errorf("isValidRoutePattern(%q) = false, want true", s)
		}
	}
	for _, s := range invalid {
		if isValidRoutePattern(s) {
			t.Errorf("isValidRoutePattern(%q) = true, want false", s)
		}
	}
}
