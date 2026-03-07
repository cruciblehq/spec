package paths

import (
	"os"
	"path/filepath"
)

const (

	// Default permission mode used when creating directories.
	DefaultDirMode os.FileMode = 0755

	// Default permission mode used when creating files.
	DefaultFileMode os.FileMode = 0644

	// Base directory for cruxd runtime state under /run.
	baseDir = "/run/cruxd"
)

// Path to the runtime directory for a cruxd instance.
//
// Contains the Unix socket and PID file while the daemon is running.
//
//	/run/cruxd/instances/<name>
func InstanceDir(name string) string {
	return filepath.Join(baseDir, "instances", name)
}

// Path to the daemon's Unix domain socket for an instance.
//
// cruxd listens here for commands from the crux CLI. In development (macOS +
// Lima), this guest socket is forwarded to the host.
//
//	/run/cruxd/instances/<name>/cruxd.sock
func Socket(name string) string {
	return filepath.Join(InstanceDir(name), "cruxd.sock")
}

// Path to the daemon's PID file for an instance.
//
//	/run/cruxd/instances/<name>/cruxd.pid
func PIDFile(name string) string {
	return filepath.Join(InstanceDir(name), "cruxd.pid")
}
