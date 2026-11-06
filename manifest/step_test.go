package manifest

import (
	"errors"
	"testing"
)

func TestStepValidate_Unpack(t *testing.T) {
	tests := []struct {
		name string
		step Step
		want error
	}{
		{
			name: "valid unpack",
			step: Step{Unpack: "rootfs.tar.gz /"},
			want: nil,
		},
		{
			name: "unpack with workdir and user",
			step: Step{Unpack: "rootfs.tar.gz /", Workdir: "/app", User: "0:0"},
			want: nil,
		},
		{
			name: "unpack and run are mutually exclusive",
			step: Step{Unpack: "rootfs.tar.gz /", Run: "echo hi"},
			want: ErrMutuallyExclusiveOps,
		},
		{
			name: "unpack and copy are mutually exclusive",
			step: Step{Unpack: "rootfs.tar.gz /", Copy: "a b"},
			want: ErrMutuallyExclusiveOps,
		},
		{
			name: "unpack rejects shell",
			step: Step{Unpack: "rootfs.tar.gz /", Shell: "/bin/bash"},
			want: ErrShellWithCopy,
		},
		{
			name: "unpack rejects env",
			step: Step{Unpack: "rootfs.tar.gz /", Env: map[string]string{"A": "1"}},
			want: ErrEnvWithCopy,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.step.Validate()
			if tt.want == nil {
				if err != nil {
					t.Fatalf("Validate() = %v, want nil", err)
				}
				return
			}
			if !errors.Is(err, tt.want) {
				t.Fatalf("Validate() = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestStepValidate(t *testing.T) {
	cases := []struct {
		name string
		step Step
		want error
	}{
		{"valid run", Step{Run: "echo hi"}, nil},
		{"empty", Step{}, ErrEmptyStep},
		{"mutually exclusive ops", Step{Run: "x", Copy: "y"}, ErrMutuallyExclusiveOps},
		{"steps without platform", Step{Steps: []Step{{Run: "x"}}}, ErrStepsWithoutPlatform},
		{"platform with operation", Step{Run: "x", Platform: "linux/arm64", Steps: []Step{{Run: "y"}}}, ErrPlatformWithOperation},
		{"shell with copy", Step{Copy: "a b", Shell: "sh"}, ErrShellWithCopy},
		{"env with copy", Step{Copy: "a b", Env: map[string]string{"K": "V"}}, ErrEnvWithCopy},
		{
			"nested platform group",
			Step{Platform: "linux/arm64", Steps: []Step{{Platform: "linux/amd64", Steps: []Step{{Run: "x"}}}}},
			ErrNestedPlatformGroup,
		},
		{"valid platform group", Step{Platform: "linux/arm64", Steps: []Step{{Run: "x"}}}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.step.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Step.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
