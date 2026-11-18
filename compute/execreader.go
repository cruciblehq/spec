package compute

import (
	"bufio"
	"errors"
	"io"

	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/crex"
)

// Reads the streamed response to an [ExecRequest].
type ExecReader struct {
	s *bufio.Scanner // Line scanner over the source stream.
}

// Returns a reader of exec messages from r.
func NewExecReader(r io.Reader) *ExecReader {
	s := bufio.NewScanner(r)
	s.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), MaxExecMessageSize)
	return &ExecReader{s: s}
}

// Returns the next message.
//
// Empty lines are skipped. Returns [io.EOF] when the stream ends cleanly, and
// [ErrExecMessageTooLarge] when a line exceeds [MaxExecMessageSize].
func (e *ExecReader) Next() (ExecResponse, error) {
	var msg ExecResponse
	for e.s.Scan() {
		if len(e.s.Bytes()) == 0 {
			continue
		}
		err := codec.Decode(e.s.Bytes(), &msg, codec.JSON)
		return msg, err
	}
	err := e.s.Err()
	if errors.Is(err, bufio.ErrTooLong) {
		return msg, crex.Wrapf(ErrExecMessageTooLarge, err, "exec message exceeds %d bytes", MaxExecMessageSize)
	}
	if err != nil {
		return msg, err
	}
	return msg, io.EOF
}
