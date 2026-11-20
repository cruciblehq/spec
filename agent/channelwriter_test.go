package agent

import (
	"bytes"
	"testing"
)

func TestExecWriterChannel(t *testing.T) {
	var buf bytes.Buffer
	w := NewExecWriter(&buf)
	n, err := w.Channel(ChannelStderr).Write([]byte("oops"))
	if err != nil || n != 4 {
		t.Fatalf("Write() = (%d, %v), want (4, nil)", n, err)
	}
	msg, err := NewExecReader(&buf).Next()
	if err != nil {
		t.Fatalf("Next() = %v, want nil", err)
	}
	if msg.Type != ExecOutput || msg.Channel != ChannelStderr || string(msg.Data) != "oops" {
		t.Errorf("message = %+v, want stderr output \"oops\"", msg)
	}
}
