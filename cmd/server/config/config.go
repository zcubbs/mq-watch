package config

// Configuration holds all the configuration data
type Configuration struct {
	MQTT     MQTTConfiguration     `mapstructure:"mqtt"`
	Database DatabaseConfiguration `mapstructure:"database"`
	Server   ServerConfiguration   `mapstructure:"server"`
}

// MQTTConfiguration holds MQTT related configuration
type MQTTConfiguration struct {
	Broker   string `mapstructure:"broker"`
	Topic    string `mapstructure:"topic"`
	ClientID string `mapstructure:"client_id"`
}

// DatabaseConfiguration holds database related configuration
type DatabaseConfiguration struct {
	Dialect     string `mapstructure:"dialect"`
	Datasource  string `mapstructure:"datasource"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

// ServerConfiguration holds HTTP server related configuration
type ServerConfiguration struct {
	Port int `mapstructure:"port"`
}
