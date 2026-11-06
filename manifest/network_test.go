package manifest

import (
	"errors"
	"testing"
)

func TestValidateProtocol(t *testing.T) {
	cases := []struct {
		name  string
		proto string
		want  error
	}{
		{"empty", "", nil},
		{"ip wildcard", "ip", nil},
		{"tcp", "tcp", nil},
		{"udp uppercase", "UDP", nil},
		{"icmp padded", "  icmp  ", nil},
		{"icmpv6", "icmpv6", nil},
		{"sctp", "sctp", nil},
		{"raw", "raw", nil},
		{"number low", "0", nil},
		{"number high", "255", nil},
		{"number mid", "47", nil},
		{"number too high", "256", ErrInvalidProtocol},
		{"negative", "-1", ErrInvalidProtocol},
		{"unknown word", "gopher", ErrInvalidProtocol},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := validateProtocol(c.proto); !errors.Is(got, c.want) {
				t.Fatalf("validateProtocol(%q) = %v, want %v", c.proto, got, c.want)
			}
		})
	}
}

func TestIngressRuleValidate(t *testing.T) {
	if err := (&IngressRule{Protocol: "tcp"}).Validate(); err != nil {
		t.Fatalf("valid ingress rule: %v", err)
	}
	if err := (&IngressRule{Protocol: "bogus"}).Validate(); !errors.Is(err, ErrInvalidProtocol) {
		t.Fatalf("invalid ingress rule = %v, want ErrInvalidProtocol", err)
	}
}

func TestEgressRuleValidate(t *testing.T) {
	if err := (&EgressRule{Protocol: "udp"}).Validate(); err != nil {
		t.Fatalf("valid egress rule: %v", err)
	}
	if err := (&EgressRule{Protocol: "bogus"}).Validate(); !errors.Is(err, ErrInvalidProtocol) {
		t.Fatalf("invalid egress rule = %v, want ErrInvalidProtocol", err)
	}
}

func TestNetworkValidate(t *testing.T) {
	cases := []struct {
		name string
		net  Network
		want error
	}{
		{"empty", Network{}, nil},
		{
			"valid rules",
			Network{
				Ingress: []IngressRule{{Protocol: "tcp"}},
				Egress:  []EgressRule{{Protocol: "udp"}},
			},
			nil,
		},
		{
			"bad ingress",
			Network{Ingress: []IngressRule{{Protocol: "nope"}}},
			ErrInvalidNetwork,
		},
		{
			"bad egress",
			Network{Egress: []EgressRule{{Protocol: "nope"}}},
			ErrInvalidNetwork,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.net.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Network.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
