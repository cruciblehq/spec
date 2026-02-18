// Defines the wire protocol for CLI-to-daemon communication.
//
// Messages are newline-delimited JSON envelopes exchanged over a Unix domain
// socket. Each envelope carries a protocol version, a command string, and an
// optional typed payload. The version field is checked on every decode and
// the message is rejected with [ErrUnsupportedVersion] when it does not match
// the compiled-in [Version] constant, ensuring the CLI and daemon always agree
// on the wire format.
//
// Encoding and decoding are handled in two phases. [Encode] marshals an
// envelope to JSON in a single step, automatically stamping the current
// version. [Decode] unmarshals only the outer envelope and leaves the payload
// as raw JSON so the caller can route on the command before committing to a
// concrete type. [DecodePayload] then parses the payload into a typed struct.
//
// The package defines request commands ([CmdBuild], [CmdStatus], [CmdShutdown])
// and response commands ([CmdOK], [CmdError]). Each request command has a
// corresponding payload type that carries the command's arguments. Payload
// types expose a Validate method that checks for missing or invalid fields.
//
// Encoding a build request:
//
//	data, err := protocol.Encode(protocol.CmdBuild, protocol.BuildRequest{
//	    Recipe:   &recipe,
//	    Resource: "my-service",
//	    Output:   "/tmp/out",
//	    Root:     "/project",
//	})
//
// Decoding a response:
//
//	env, raw, err := protocol.Decode(data)
//	if err != nil {
//	    log.Fatal(err) // version mismatch or malformed message
//	}
//
//	switch env.Command {
//	case protocol.CmdOK:
//	    result, err := protocol.DecodePayload[protocol.BuildResult](raw)
//	case protocol.CmdError:
//	    result, err := protocol.DecodePayload[protocol.ErrorResult](raw)
//	}
package protocol
