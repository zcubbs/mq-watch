package config

import (
	"github.com/spf13/viper"
	"github.com/zcubbs/x/pretty"
)

func LoadConfiguration(configFile string) (Configuration, error) {
	var configuration Configuration

	v := viper.New()
	v.SetConfigFile(configFile)

	viper.SetDefault("database.datasource", "mq-watch.db")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	if err := v.Unmarshal(&configuration); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

func PrintConfiguration(config Configuration) {
	// Print out the configuration
	pretty.PrintJson(config)
}
