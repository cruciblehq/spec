package manifest

import (
	"errors"
	"testing"

	"github.com/cruciblehq/utils-go/codec"
)

func TestMergeMap(t *testing.T) {
	dst := map[string]any{"a": 1}
	src := map[string]any{"b": 2}
	got, err := mergeMap(dst, src)
	if err != nil {
		t.Fatalf("disjoint merge: %v", err)
	}
	if got["a"] != 1 || got["b"] != 2 {
		t.Fatalf("merged map = %v, want a=1 b=2", got)
	}

	if _, err := mergeMap(map[string]any{"x": 1}, map[string]any{"x": 2}); !errors.Is(err, ErrEncodeFailed) {
		t.Fatalf("conflicting merge = %v, want ErrEncodeFailed", err)
	}
}

func TestEncodeToMap(t *testing.T) {
	svc := &Service{
		Recipe:     Recipe{Stages: []Stage{{From: "img"}}},
		Entrypoint: []string{"/bin/app"},
	}
	m, err := encodeToMap(&codec.Codec{}, svc)
	if err != nil {
		t.Fatalf("encodeToMap(service): %v", err)
	}
	if _, ok := m["entrypoint"]; !ok {
		t.Fatalf("encoded service missing entrypoint: %v", m)
	}

	m, err = encodeToMap(&codec.Codec{}, &ComputeLocal{Host: "h"})
	if err != nil {
		t.Fatalf("encodeToMap(local): %v", err)
	}
	if m["Host"] != "h" {
		t.Fatalf("encoded local = %v, want Host=h", m)
	}
}
