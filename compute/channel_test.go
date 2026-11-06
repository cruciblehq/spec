package compute

import "testing"

func TestChannelReserved(t *testing.T) {
	cases := []struct {
		channel Channel
		want    bool
	}{
		{ChannelStdout, true},
		{ChannelStderr, true},
		{"<custom>", true},
		{"<>", false},
		{"stdout", false},
		{"", false},
		{"<", false},
		{"<a", false},
		{"a>", false},
	}
	for _, c := range cases {
		if got := c.channel.Reserved(); got != c.want {
			t.Errorf("Channel(%q).Reserved() = %v, want %v", c.channel, got, c.want)
		}
	}
}
