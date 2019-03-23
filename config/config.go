package config

// Config .
type Config struct {
	Service  map[string]string
	Database map[string]string
	Etc      map[string]string
}

var configFile = "service.conf"
