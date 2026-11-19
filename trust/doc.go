// Package trust defines how a client and the agents on its machines establish
// and verify trust.
//
// The client holds a root key. The root signs issuer certificates that expire
// after a month, and issuers sign the client certificates the client presents
// in mutual TLS. The root also signs each agent's certificate once, at
// enrollment. That certificate names the agent's machine and lasts for the life
// of the machine. The root revokes an issuer or a machine by listing its
// certificate in a CRL (RFC 5280). Agents keep the list with the highest CRL
// number they have seen and refuse any issuer it names, and the client refuses
// any machine it names.
//
// A new machine starts untrusted. The client places an [Enrollment] in the data
// the machine boots from, holding the root certificate, the current revocation
// list, and a token that works once. On first contact the agent proves it holds
// the token with a [Proof] bound to the TLS session through a TLS exporter
// (RFC 5705), the client signs the agent's certificate, and only then does the
// agent serve its protocol.
//
// Every key is Ed25519. Every connection uses TLS 1.3 with the hybrid X25519
// and ML-KEM-768 key exchange ([ServerConfig], [ClientConfig], and
// [EnrollConfig]).
package trust
