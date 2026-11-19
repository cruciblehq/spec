package agent

// Requests a handshake with the client's protocol version and revocation list.
//
// The client sends this to [EndpointHandshake] before any other request. The
// agent stores Revocations if the root signed the list and its CRL number is
// higher than the list it has. The agent does not check Version; it answers
// with its own, and the client decides.
type HandshakeRequest struct {
	Version     string `codec:"version"`     // Supported protocol version.
	Revocations []byte `codec:"revocations"` // DER encoded CRL signed by the root.
}

// Carries the protocol version and build identity of the agent.
//
// Returned from [EndpointHandshake]. The client compares Version against its
// own and refuses incompatible versions.
type HandshakeResponse struct {
	Version string `codec:"version"`          // Supported protocol version.
	Build   string `codec:"build,omitempty"`  // Agent build version.
	Commit  string `codec:"commit,omitempty"` // Agent git commit.
}
