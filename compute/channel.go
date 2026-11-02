package compute

// A channel identifier.
//
// Protocol output is multiplexed over one or more named channels (not just
// stdout and stderr). Names enclosed in angle brackets (e.g., [ChannelStdout]
// and [ChannelStderr]) are reserved by the protocol. Any other name denotes a
// custom channel.
type Channel string

// Reserved protocol channels.
const (
	ChannelStdout Channel = "<stdout>" // Standard output.
	ChannelStderr Channel = "<stderr>" // Standard error.
)

// Whether the channel name is reserved.
//
// Reserved names are enclosed in angle brackets. The protocol owns that
// namespace, so application-defined channels must use something else.
func (c Channel) Reserved() bool {
	return len(c) >= 2 && c[0] == '<' && c[len(c)-1] == '>'
}
