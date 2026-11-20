package agent

// Writer that sends each write as output on one channel of an exec stream.
type channelWriter struct {
	stream  *ExecWriter // Stream the output messages are written to.
	channel Channel     // Channel the output messages are tagged with.
}

// Sends p as output on the writer's channel.
func (c *channelWriter) Write(p []byte) (int, error) {
	if err := c.stream.Output(c.channel, p); err != nil {
		return 0, err
	}
	return len(p), nil
}
