package compute

// Sends the client's protocol version (handshake).
//
// crux sends this to [EndpointHandshake] before any other request. cruxd
// compares Version against its own and rejects incompatible versions.
type HandshakeRequest struct {
	Version string `codec:"version"` // Supported protocol version.
}

// Reports the protocol version and build identity of a cruxd instance.
//
// Returned from [EndpointHandshake]. crux compares Version against its own
// and refuses incompatible versions.
type HandshakeResponse struct {
	Version string `codec:"version"`          // Supported protocol version.
	Build   string `codec:"build,omitempty"`  // cruxd build version.
	Commit  string `codec:"commit,omitempty"` // cruxd git commit.
}
