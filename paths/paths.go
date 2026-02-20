package paths

import "os"

const (

	// Default permission mode used when creating directories.
	DefaultDirMode os.FileMode = 0755

	// Default permission mode used when creating files.
	DefaultFileMode os.FileMode = 0644
)

// Path to the daemon's runtime directory.
//
// Contains the Unix socket and PID file while the daemon is running.
func Runtime() string {
	return "/run/cruxd"
}

// Path to the daemon's Unix domain socket.
//
// cruxd listens here for commands from the crux CLI. In development
// (macOS + Lima), this guest socket is forwarded to the host.
func Socket() string {
	return "/run/cruxd/cruxd.sock"
}

// Path to the daemon's PID file.
func PIDFile() string {
	return "/run/cruxd/cruxd.pid"
}
