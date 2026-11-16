// Package trust defines how crux and cruxd establish and verify trust.
//
// crux holds a root key. The root signs issuer certificates that expire after
// a month, and issuers sign the client certificates crux presents in mutual
// TLS. The root also signs each machine's certificate once, at enrollment,
// for the life of the machine. The root revokes an issuer or a machine by
// listing it in a CRL (RFC 5280). Machines keep the list with the highest CRL
// number they have seen and refuse any issuer it names, and crux refuses any
// machine it names.
//
// A new machine starts untrusted. crux places an [Enrollment] in the data the
// machine boots from, holding the root certificate, the current revocation
// list, and a token that works once. On first contact cruxd proves it holds
// the token with a [Proof] bound to the TLS session through a TLS exporter
// (RFC 5705), crux signs the machine's certificate, and only then does cruxd
// serve its API.
//
// Every key is Ed25519. Every connection uses TLS 1.3 with the hybrid X25519
// and ML-KEM-768 key exchange ([ServerConfig], [ClientConfig], and
// [EnrollConfig]).
package trust
