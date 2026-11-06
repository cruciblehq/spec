// Package compute defines the protocol shared by crux and cruxd.
//
// crux is the client and cruxd is the daemon that drives containerd. This
// package is the source of their shared vocabulary.
//
// [Container] and [Runtime] are interfaces that describe the runtime.
// Two implementations satisfy them: one implemented by cruxd internally to
// manage containerd, and an HTTP-client implemented by crux.
//
// The protocol runs over HTTP/2 and JSON on a mTLS-over-TCP stack, even on the
// local VM. HTTP/2 is negotiated automatically via TLS ALPN, giving multiplexed,
// flow-controlled streams over one connection.
//
// Every operation is one HTTP/2 stream (one request and one response). A
// response may include one or more messages. The [ExecRequest] response is
// such a stream: each [ExecResponse] carries a chunk of process output on a
// named [Channel], or a terminal result ([ExecExit] with an exit code,
// [ExecError] with a failure). Output is multiplexed over channels rather
// than fixed to stdout and stderr; names enclosed in angle brackets, such
// as [ChannelStdout], are reserved by the protocol, and any other name
// denotes an application-defined channel.
//
// [Protocol] is the wire-contract version. crux calls [EndpointHandshake],
// reads the [HandshakeResponse], and refuses an incompatible cruxd before
// issuing any other request.
//
// mTLS authenticates the transport. On top of it, crux presents a [Credential]
// as a bearer token in the [HeaderAuthorization] header, and cruxd verifies it
// through an [Authenticator]. [AllowAll] is the interim implementation that
// admits every credential, so the auth boundary is wired end-to-end and later
// hardening is an implementation swap, not a contract change.
//
// A unary operation that cannot complete returns a non-2xx status with an
// [Error] body ([MediaTypeError]) whose [ErrorCode] classifies the failure;
// [Error] satisfies the error interface. A streaming operation reports a
// failure that occurs before its response begins the same way, but a failure
// mid-stream arrives as a terminal [ExecError] message, because the status is
// already committed once streaming has started.
//
// JSON bodies are serialized with the project codec using each type's codec
// struct tags. Every body carries a vendor [MediaType] following the
// application/vnd.crucible.compute.{name}.v0 pattern; JSON bodies use the
// [MediaType.JSON] form and binary and streaming bodies use the type as-is.
package compute
