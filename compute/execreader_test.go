package compute

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestExecReaderSkipsEmptyLines(t *testing.T) {
	r := NewExecReader(strings.NewReader("\n\n{\"type\":\"exit\",\"code\":1}\n\n"))

	got, err := r.Next()
	if err != nil {
		t.Fatalf("Next() = %v, want nil", err)
	}
	if got.Type != ExecExit || got.Code != 1 {
		t.Errorf("Next() = %+v, want exit with code 1", got)
	}
	if _, err := r.Next(); !errors.Is(err, io.EOF) {
		t.Errorf("Next() at end = %v, want %v", err, io.EOF)
	}
}

func TestExecReaderInvalidMessage(t *testing.T) {
	r := NewExecReader(strings.NewReader("{bad\n"))

	if _, err := r.Next(); err == nil {
		t.Fatal("Next() = nil, want an error")
	}
}

func TestExecReaderLineTooLong(t *testing.T) {
	line := bytes.Repeat([]byte("a"), MaxExecMessageSize+1)
	r := NewExecReader(bytes.NewReader(line))

	if _, err := r.Next(); !errors.Is(err, ErrExecMessageTooLarge) {
		t.Fatalf("Next() = %v, want %v", err, ErrExecMessageTooLarge)
	}
}
