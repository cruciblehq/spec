package compute

// Sends the client's protocol version and revocation list (handshake).
//
// crux sends this to [EndpointHandshake] before any other request. cruxd
// compares Version against its own and rejects the connection when the version
// does not match. cruxd stores Revocations if the root signed it and its CRL
// number is higher than the list it has.
type HandshakeRequest struct {
	Version     string `codec:"version"`     // Supported protocol version.
	Revocations []byte `codec:"revocations"` // DER encoded CRL signed by the root.
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
