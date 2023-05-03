package model

import (
	"errors"

	"gorm.io/gorm"
	"shoppingList-backend-gin.com/m/config"
	"shoppingList-backend-gin.com/m/utils"

	"gorm.io/driver/postgres"
)

func ConnectDatabase() error {

	DSN, err := GetDSN()

	if err != nil {
		return err
	}

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := AutoMigrateTables(database); err != nil {
		return err
	}

	DB = database
	return nil
}

func AutoMigrateTables(database *gorm.DB) error {
	if err := database.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err := database.AutoMigrate(&Entry{}); err != nil {
		return err
	}
	if err := database.AutoMigrate(&Unit{}); err != nil {
		return err
	}

	return nil
}

func GetDSN() (string, error) {
	if config.GetConfiguration().Server.GetEnvironment() == "QA" {
		return utils.Format("user={.user} password={.password} dbname={.dbname} port={.port} sslmode=disable TimeZone=America/Sao_Paulo", config.GetConfiguration().QADatabase), nil
	} else if config.GetConfiguration().Server.GetEnvironment() == "DEV" {
		return utils.Format("user={.user} password={.password} dbname={.dbname} port={.port} sslmode=disable TimeZone=America/Sao_Paulo", config.GetConfiguration().DEVDatabase), nil
	} else if config.GetConfiguration().Server.GetEnvironment() == "PROD" {
		return utils.Format("user={.user} password={.password} dbname={.dbname} port={.port} sslmode=disable TimeZone=America/Sao_Paulo", config.GetConfiguration().DEVDatabase), nil
	}

	return "", errors.New("Server config dont have a valid environment")
}
