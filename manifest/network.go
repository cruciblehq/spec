package manifest

import (
	"strconv"
	"strings"

	"github.com/cruciblehq/utils-go/crex"
)

// Protocol keywords accepted in network rules.
const (
	ProtocolAny    = "ip"     // Any protocol.
	ProtocolTCP    = "tcp"    // Transmission Control Protocol.
	ProtocolUDP    = "udp"    // User Datagram Protocol.
	ProtocolICMP   = "icmp"   // Internet Control Message Protocol.
	ProtocolICMPv6 = "icmpv6" // ICMP for IPv6.
	ProtocolSCTP   = "sctp"   // Stream Control Transmission Protocol.
	ProtocolRaw    = "raw"    // Raw IP packets.
)

// Largest IP protocol number.
const maxProtocolNumber = 255

// Cloud network perimeter configuration.
//
// Expresses what flows are permitted at the cloud boundary. No flows are
// implied; every allowed path must be declared explicitly.
type Network struct {

	// Inbound flows allowed at the cloud perimeter.
	Ingress []IngressRule `codec:"ingress,omitempty"`

	// Outbound flows allowed at the cloud perimeter.
	Egress []EgressRule `codec:"egress,omitempty"`
}

// Validates all network rules.
func (n *Network) Validate() error {
	for i := range n.Ingress {
		if err := n.Ingress[i].Validate(); err != nil {
			return crex.At(crex.Wrap(ErrInvalidNetwork, err), "ingress rule", i+1)
		}
	}
	for i := range n.Egress {
		if err := n.Egress[i].Validate(); err != nil {
			return crex.At(crex.Wrap(ErrInvalidNetwork, err), "egress rule", i+1)
		}
	}
	return nil
}

// Inbound flow allowed at the cloud perimeter.
//
// Permits traffic from Source on Protocol/Port to reach the service.
type IngressRule struct {

	// Network protocol.
	//
	// A protocol keyword, the "ip" wildcard for any protocol, or a decimal IP
	// protocol number, following the IANA Protocol Numbers registry.
	Protocol string `codec:"protocol"`

	// Destination port on the service.
	//
	// Zero means any port.
	Port uint16 `codec:"port,omitempty"`

	// Traffic source.
	//
	// A CIDR (e.g. "0.0.0.0/0"), a Crucible service name, or "*" for any source.
	Source string `codec:"source,omitempty"`
}

// Validates that the rule specifies a recognised protocol.
func (r *IngressRule) Validate() error {
	return validateProtocol(r.Protocol)
}

// Outbound flow allowed at the cloud perimeter.
//
// Permits traffic from the service to Destination on Protocol/Port.
type EgressRule struct {

	// Network protocol.
	//
	// A protocol keyword, the "ip" wildcard for any protocol, or a decimal IP
	// protocol number, following the IANA Protocol Numbers registry.
	Protocol string `codec:"protocol"`

	// Destination port.
	//
	// Zero means any port.
	Port uint16 `codec:"port,omitempty"`

	// Traffic destination.
	//
	// A CIDR (e.g. "0.0.0.0/0"), a Crucible service name, or "*" for any destination.
	Destination string `codec:"destination,omitempty"`
}

// Validates that the rule specifies a recognised protocol.
func (r *EgressRule) Validate() error {
	return validateProtocol(r.Protocol)
}

// Checks that p is a recognised network protocol name.
//
// Accepts the protocol keywords, the "ip" wildcard for any protocol, and a
// decimal IP protocol number in the range 0–255.
func validateProtocol(p string) error {
	switch strings.ToLower(strings.TrimSpace(p)) {
	case "", ProtocolAny, ProtocolTCP, ProtocolUDP, ProtocolICMP, ProtocolICMPv6, ProtocolSCTP, ProtocolRaw:
		return nil
	}
	if n, err := strconv.Atoi(p); err == nil && n >= 0 && n <= maxProtocolNumber {
		return nil
	}
	return crex.Newf(ErrInvalidProtocol, "unknown protocol %q", p)
}
