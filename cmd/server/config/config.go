package config

// Configuration holds all the configuration data
type Configuration struct {
	MQTT     MQTTConfiguration     `mapstructure:"mqtt"`
	Tenants  []TenantConfiguration `mapstructure:"tenants"`
	Database DatabaseConfiguration `mapstructure:"database"`
	Server   ServerConfiguration   `mapstructure:"server"`
}

// MQTTConfiguration holds MQTT related configuration
type MQTTConfiguration struct {
	Broker   string `mapstructure:"broker"`
	ClientID string `mapstructure:"client_id"`
}

// TenantConfiguration holds tenant related configuration
type TenantConfiguration struct {
	Name   string   `mapstructure:"name"`
	Topics []string `mapstructure:"topics"`
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
