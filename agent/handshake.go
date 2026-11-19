package agent

// Sends the client's protocol version and revocation list (handshake).
//
// The client sends this to [EndpointHandshake] before any other request. The agent
// compares Version against its own and rejects the connection when the version
// does not match. The agent stores Revocations if the root signed it and its CRL
// number is higher than the list it has.
type HandshakeRequest struct {
	Version     string `codec:"version"`     // Supported protocol version.
	Revocations []byte `codec:"revocations"` // DER encoded CRL signed by the root.
}

// Reports the protocol version and build identity of an agent.
//
// Returned from [EndpointHandshake]. The client compares Version against its own
// and refuses incompatible versions.
type HandshakeResponse struct {
	Version string `codec:"version"`          // Supported protocol version.
	Build   string `codec:"build,omitempty"`  // Agent build version.
	Commit  string `codec:"commit,omitempty"` // Agent git commit.
}
