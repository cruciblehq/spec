// Package compute defines the protocol shared by crux and cruxd.
//
// crux is the client and cruxd is the daemon that drives containerd. This
// package is the source of their shared vocabulary.
//
// [Container] and [Runtime] are interfaces that describe the runtime. Two
// implementations satisfy them: one implemented by cruxd internally to manage
// containerd, and an HTTP-client implemented by crux.
//
// The protocol runs over HTTP/2 and JSON on a mTLS-over-TCP stack. HTTP/2 is
// negotiated automatically via TLS ALPN, giving multiplexed, flow-controlled
// streams over one connection.
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
// Both ends are authenticated via mTLS (see package trust). A new machine
// serves only [EndpointEnroll] and [EndpointCertify] until crux installs
// its certificate.
//
// A unary operation that cannot complete returns an [Error] body of type
// [MediaTypeError] whose [ErrorCode] classifies the failure, with the status
// given by [ErrorCode.HTTPStatus]; [Error] satisfies the error interface. A
// streaming operation reports a failure that occurs before its response begins
// the same way, but a failure mid-stream arrives as a terminal [ExecError]
// message, because the status is already committed once streaming has started.
//
// Message bodies are encoded in [BodyFormat] with the project codec using each
// type's codec struct tags. Every body carries a vendor media type following
// the application/vnd.crucible.compute.{name}.v0 pattern. Encoded bodies add
// the suffix of [BodyFormat] (see [mediatype.MediaType.With]); binary and
// streaming bodies use the type as is.
package compute
