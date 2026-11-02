// Package compute defines the protocol shared by crux and cruxd.
//
// crux is the client and cruxd is the daemon that drives containerd. This
// package is the single source of their shared vocabulary.
//
// [Container] and [Runtime] describe the runtime, amnd two implementations
// satisfy them (interfaces): one that cruxd runs internally, and the HTTP
// client in crux. cruxd also runs the server, but the server does not satisfy
// these interfaces (it decodes each request and calls cruxd's internal
// implementation).
//
// The protocol runs over HTTP/2 and JSON over mTLS-over-TCP. HTTP/2 is
// negotiated automatically via TLS ALPN.
//
// [ExecRequest] is an HTTP/2 stream: each [ExecResponse] includes a chunk of
// process output on a named [Channel], or a terminal result ([ExecExit] with
// an exit code, [ExecError] with a failure).
//
// [Version] is the protocol version. crux calls [EndpointHandshake], reads the
// [HandshakeResponse], and refuses an incompatible cruxd.
//
// crux presents a bearer token (in [HeaderAuthorization]) with each request.
// The token is verified by cruxd using an [Authenticator].
//
// The HTTP status reflects only the transport (whether the request was
// received, authenticated, routed, and answered). A malformed, unauthenticated,
// or unroutable request gets a non-2xx status and the operation never runs; a
// server fault that prevents any response is a 5xx.
//
// An operation that runs (e.g., the request was well-formed and authenticated)
// returns the result in the response payload (2xx status). In case of failure,
// the response contains an [Error] body ([MediaTypeError]) with an [ErrorCode].
// A streaming operation carries the same [Error] in its terminal [ExecError]
// message. HTTP status is never used to signal an operation failure.
package compute
