package agent

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type flushRecorder struct {
	bytes.Buffer
	flushes int
}

func (f *flushRecorder) Flush() {
	f.flushes++
}

func TestExecWriterSend(t *testing.T) {
	var buf flushRecorder
	w := NewExecWriter(&buf)
	sent := []ExecResponse{
		{Type: ExecOutput, Channel: ChannelStdout, Data: []byte("hello\n")},
		{Type: ExecExit, Code: 3},
	}
	for _, msg := range sent {
		if err := w.Send(msg); err != nil {
			t.Fatalf("Send() = %v, want nil", err)
		}
	}
	if buf.flushes != len(sent) {
		t.Errorf("flushes = %d, want %d", buf.flushes, len(sent))
	}

	r := NewExecReader(&buf.Buffer)
	for _, want := range sent {
		got, err := r.Next()
		if err != nil {
			t.Fatalf("Next() = %v, want nil", err)
		}
		if got.Type != want.Type || got.Code != want.Code || !bytes.Equal(got.Data, want.Data) {
			t.Errorf("Next() = %+v, want %+v", got, want)
		}
	}
	if _, err := r.Next(); !errors.Is(err, io.EOF) {
		t.Errorf("Next() at end = %v, want %v", err, io.EOF)
	}
}

func TestExecWriterOutputChunks(t *testing.T) {
	var buf bytes.Buffer
	data := bytes.Repeat([]byte("x"), MaxExecChunkSize+10)

	if err := NewExecWriter(&buf).Output(ChannelStderr, data); err != nil {
		t.Fatalf("Output() = %v, want nil", err)
	}

	r := NewExecReader(&buf)
	var got []byte
	for {
		msg, err := r.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("Next() = %v, want nil", err)
		}
		if msg.Type != ExecOutput || msg.Channel != ChannelStderr {
			t.Errorf("message = %+v, want stderr output", msg)
		}
		if len(msg.Data) > MaxExecChunkSize {
			t.Errorf("len(Data) = %d, want at most %d", len(msg.Data), MaxExecChunkSize)
		}
		got = append(got, msg.Data...)
	}
	if !bytes.Equal(got, data) {
		t.Errorf("reassembled %d bytes, want %d", len(got), len(data))
	}
}

func TestExecWriterOutputEmpty(t *testing.T) {
	var buf bytes.Buffer
	if err := NewExecWriter(&buf).Output(ChannelStdout, nil); err != nil {
		t.Fatalf("Output() = %v, want nil", err)
	}
	if buf.Len() != 0 {
		t.Errorf("wrote %d bytes, want 0", buf.Len())
	}
}

func TestExecWriterMessageTooLarge(t *testing.T) {
	var buf bytes.Buffer
	msg := ExecResponse{Type: ExecOutput, Channel: ChannelStdout, Data: make([]byte, MaxExecMessageSize)}

	if err := NewExecWriter(&buf).Send(msg); !errors.Is(err, ErrExecMessageTooLarge) {
		t.Fatalf("Send() = %v, want %v", err, ErrExecMessageTooLarge)
	}
	if buf.Len() != 0 {
		t.Errorf("wrote %d bytes, want 0", buf.Len())
	}
}
