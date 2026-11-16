package compute

import (
	"bufio"
	"errors"
	"io"

	"github.com/cruciblehq/utils-go/codec"
	"github.com/cruciblehq/utils-go/crex"
)

// Largest encoded exec message in bytes, newline included.
const MaxExecMessageSize = 16 << 20

// Largest output chunk one [ExecOutput] message carries.
//
// Output is base64 inside the message, which grows it by a third, so chunks of
// this size stay well under [MaxExecMessageSize].
const MaxExecChunkSize = 8 << 20

// Writes the streamed response to an [ExecRequest].
//
// Each [ExecResponse] is one JSON object followed by a newline. The stream is
// always JSON, whatever [BodyFormat] is, because the framing needs every
// message on a single line.
type ExecWriter struct {
	w io.Writer // Destination stream.
}

// Returns a writer of exec messages to w.
func NewExecWriter(w io.Writer) *ExecWriter {
	return &ExecWriter{w: w}
}

// Writes msg as one line, then flushes w when it can be flushed.
//
// Returns [ErrExecMessageTooLarge] when the encoded message exceeds
// [MaxExecMessageSize]; nothing is written then.
func (e *ExecWriter) Send(msg ExecResponse) error {
	data, err := codec.Encode(msg, codec.JSON)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	if len(data) > MaxExecMessageSize {
		return crex.Newf(ErrExecMessageTooLarge, "exec message of %d bytes exceeds %d", len(data), MaxExecMessageSize)
	}
	if _, err := e.w.Write(data); err != nil {
		return err
	}
	if f, ok := e.w.(interface{ Flush() }); ok {
		f.Flush()
	}
	return nil
}

// Sends p as [ExecOutput] messages on channel ch.
//
// p is split into chunks of at most [MaxExecChunkSize] bytes, one message each.
func (e *ExecWriter) Output(ch Channel, p []byte) error {
	for len(p) > 0 {
		n := min(len(p), MaxExecChunkSize)
		if err := e.Send(ExecResponse{Type: ExecOutput, Channel: ch, Data: p[:n]}); err != nil {
			return err
		}
		p = p[n:]
	}
	return nil
}

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
