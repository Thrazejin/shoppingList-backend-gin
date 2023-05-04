package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	QADatabase   QADatabaseConfigurations
	DEVDatabase  DEVDatabaseConfigurations
	PRODdatabase PRODdatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Environment string
	Host        string
	Port        int
}

// DatabaseConfigurations exported
type QADatabaseConfigurations struct {
	User     string
	Password string
	DBname   string
	Port     int
}

// DatabaseConfigurations exported
type DEVDatabaseConfigurations struct {
	User     string
	Password string
	DBname   string
	Port     int
}

// DatabaseConfigurations exported
type PRODdatabaseConfigurations struct {
	User     string
	Password string
	DBname   string
	Port     int
}

var configuration Configurations

func ReadConfig(environment string) error {
	viper.SetConfigName("config")
	if environment == "QA" {
		viper.AddConfigPath("..")
	} else {
		viper.AddConfigPath(".")
	}
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return err
	}

	viper.Set("server.environment", environment)

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
