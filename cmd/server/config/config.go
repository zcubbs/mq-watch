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
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// TenantConfiguration holds tenant related configuration
type TenantConfiguration struct {
	Name   string   `mapstructure:"name"`
	Topics []string `mapstructure:"topics"`
}

// DatabaseConfiguration holds database related configuration
type DatabaseConfiguration struct {
	Sqlite   SQLiteConfiguration   `mapstructure:"sqlite"`
	Postgres PostgresConfiguration `mapstructure:"postgres"`
}

// SQLiteConfiguration holds SQLite related configuration
type SQLiteConfiguration struct {
	Enabled     bool   `mapstructure:"enabled"`
	Datasource  string `mapstructure:"datasource"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

// PostgresConfiguration holds PostgreSQL related configuration
type PostgresConfiguration struct {
	Enabled     bool   `mapstructure:"enabled"`
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	SSLMode     bool   `mapstructure:"sslmode"`
	AutoMigrate bool   `mapstructure:"auto_migrate"`
}

// ServerConfiguration holds HTTP server related configuration
type ServerConfiguration struct {
	Port        int    `mapstructure:"port"`
	TlsEnabled  bool   `mapstructure:"tls_enabled"`
	TlsCertFile string `mapstructure:"tls_cert_file"`
	TlsKeyFile  string `mapstructure:"tls_key_file"`
}
