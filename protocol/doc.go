// Defines the wire protocol for CLI-to-daemon communication.
//
// Messages are newline-delimited JSON envelopes exchanged over a Unix domain
// socket. Each envelope carries a command and an optional typed payload.
package protocol
