// Package paths holds the default paths for devludns.
// Values are confirmed in issues #3, #6, #12, and #19.
package paths

// Default paths — may be overridden via CLI flags or config file.
const (
	// DefaultStateDB is the default path to the SQLite state database.
	DefaultStateDB = "/var/lib/dvludns/dvludns.db"

	// DefaultSocketPath is the default path for the control Unix socket.
	DefaultSocketPath = "/run/dvludns.sock"

	// DefaultConfigFile is the default path for the YAML configuration file.
	DefaultConfigFile = "/etc/devlusoft/devludns.yaml"

	// DefaultCertsDir is the default directory for TLS certificates.
	DefaultCertsDir = "/etc/devlusoft/certs"
)
