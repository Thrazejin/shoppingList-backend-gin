package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server      ServerConfigurations
	QADatabase  QADatabaseConfigurations
	DEVDatabase DEVDatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port        int
	environment string
	host        string
	port        int
}

// DatabaseConfigurations exported
type QADatabaseConfigurations struct {
	user     string
	password string
	dbname   string
	port     int
}

// DatabaseConfigurations exported
type DEVDatabaseConfigurations struct {
	user     string
	password string
	dbname   string
	port     int
}

var configuration Configurations

func ReadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}

func GetConfiguration() Configurations {
	return configuration
}

func (serverConfig ServerConfigurations) GetEnvironment() string {
	return serverConfig.environment
}
