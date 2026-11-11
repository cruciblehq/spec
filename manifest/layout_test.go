package manifest

import (
	"path/filepath"
	"testing"
)

func TestLayout(t *testing.T) {
	dir := filepath.Join("a", "b")
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"Path", Path(dir), filepath.Join(dir, "crucible.yaml")},
		{"PlanPath", PlanPath(dir), filepath.Join(dir, "plan.yaml")},
		{"ImagePath", ImagePath(dir), filepath.Join(dir, "image.tar")},
		{"WidgetMainPath", WidgetMainPath(dir), filepath.Join(dir, "index.js")},
		{"BuildDir", BuildDir(dir), filepath.Join(dir, "build")},
		{"PackagePath", PackagePath(dir), filepath.Join(dir, "dist", "package.tar.zst")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}
