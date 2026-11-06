package manifest

import (
	"errors"
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func TestStageValidate(t *testing.T) {
	cases := []struct {
		name  string
		stage Stage
		want  error
	}{
		{"empty scratch stage", Stage{}, nil},
		{"valid named stage", Stage{Name: "builder", From: "crucible/runtime 0.1.0"}, nil},
		{"invalid name", Stage{Name: "bad name"}, ErrInvalidStageName},
		{"invalid platform", Stage{Platform: "not-a-platform"}, ErrInvalidPlatform},
		{"valid platform", Stage{Platform: "linux/arm64"}, nil},
		{"args without from", Stage{Args: Args{"k": "v"}}, ErrMissingFrom},
		{"args with from", Stage{From: "img", Args: Args{"key": "v"}}, nil},
		{
			"step uses platform in platform stage",
			Stage{Platform: "linux/arm64", Steps: []Step{{Platform: "linux/amd64"}}},
			ErrPlatformInPlatformStage,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.stage.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Stage.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestStageEncode(t *testing.T) {
	s := &Stage{
		Name:     "builder",
		Platform: "linux/arm64",
		From:     "img",
		Args:     Args{"key": "val"},
		Steps:    []Step{{Run: "echo hi"}},
	}
	m, err := s.Encode(&codec.Codec{})
	if err != nil {
		t.Fatalf("Stage.Encode: %v", err)
	}
	out := m.(map[string]any)
	for _, k := range []string{"name", "platform", "from", "args", "steps"} {
		if _, ok := out[k]; !ok {
			t.Fatalf("stage encode missing %q: %v", k, out)
		}
	}
}

func TestStageDecodeRejectsNonMap(t *testing.T) {
	if err := (&Stage{}).Decode(&codec.Codec{}, "not a map"); !errors.Is(err, ErrInvalidStage) {
		t.Fatalf("Stage.Decode(non-map) = %v, want ErrInvalidStage", err)
	}
}
