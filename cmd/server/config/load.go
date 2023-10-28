package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfiguration(configFile string) (Configuration, error) {
	v := viper.New()
	v.SetConfigFile(configFile)

	viper.SetDefault("database.datasource", "mq-watch.db")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	var configuration Configuration
	if err := v.Unmarshal(&configuration); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

func PrintConfiguration(config Configuration) {
	// Print out the configuration
	fmt.Printf("MQTT Broker: %s\n", config.MQTT.Broker)
	fmt.Printf("MQTT Topic: %s\n", config.MQTT.Topic)
	fmt.Printf("Database Dialect: %s\n", config.Database.Dialect)
	fmt.Printf("Database Datasource: %s\n", config.Database.Datasource)
	fmt.Printf("Database AutoMigrate: %t\n", config.Database.AutoMigrate)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
}
