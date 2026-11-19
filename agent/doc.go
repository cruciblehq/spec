// Package agent defines the protocol for talking to an agent.
//
// An agent is the program that runs on each machine and acts for a client: it
// runs and manages the deployments the client sends it, and keeps them as
// intended. The agent serves this protocol once it is enrolled. [Agent] and
// [Container] describe what an agent offers. A server implements them directly,
// and a client implements them by sending each call over this protocol.
//
// The protocol runs HTTP/2 over mutual TLS on TCP. HTTP/2 is negotiated through
// TLS ALPN, which gives many streams with flow control over one connection.
//
// Every operation is one HTTP/2 stream with one request and one response, and
// has a request message and a response message (some of them empty). Binary
// data travels in a multipart/related body. Its first part is the operation's
// message and its second part is the data. Import and copy requests, and
// export and extract responses, are sent this way.
//
// The [ExecRequest] response is a stream of messages. Each [ExecResponse]
// carries a chunk of process output on a named [Channel], or a terminal result
// ([ExecExit] with an exit code, [ExecError] with a failure). Output is not
// fixed to stdout and stderr but shares the stream across channels. Names
// enclosed in angle brackets, such as [ChannelStdout], are reserved by the
// protocol, and any other name denotes a channel the application defines.
//
// [Version] is the version of the wire contract. The client calls
// [EndpointHandshake], reads the [HandshakeResponse], and refuses an
// incompatible agent before it sends any other request.
//
// Both ends are authenticated with mutual TLS (see package trust). The agent
// on a new machine serves only [EndpointEnroll] and [EndpointCertify] until the
// client installs the agent certificate.
//
// A unary operation that cannot complete returns an [Error] body of type
// [MediaTypeError] whose [ErrorCode] classifies the failure, with the status
// given by [ErrorCode.HTTPStatus]; [Error] satisfies the error interface. A
// streaming operation reports a failure that occurs before its response begins
// the same way, but a failure during the stream arrives as a terminal
// [ExecError] message, because the status is already sent once streaming has
// started.
//
// Messages are encoded in [BodyFormat] with package codec, using each type's
// codec struct tags. Every message and data part carries a vendor media type
// following the application/vnd.crucible.agent.{name}.v0 pattern. Messages
// add the suffix of [BodyFormat] (see [mediatype.MediaType.With]). Data parts
// and the exec stream use the type as is.
package agent
