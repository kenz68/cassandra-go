package utils

// Configuration hold the all configurations in a single struct
type Configuration struct {
	Database DBSettings
	Host     HostSettings
}

// DBSettings hold the database settings
type DBSettings struct {
	Url      string
	Keyspace string
}

// HostSettings hold the host settings
type HostSettings struct {
	Port string
}
