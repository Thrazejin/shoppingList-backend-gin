package model

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
	"shoppingList-backend-gin.com/m/config"

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
	if err := database.AutoMigrate(&AppUser{}); err != nil {
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
	if config.GetConfiguration().Server.Environment == "QA" {
		return FormatDSN(
			config.GetConfiguration().QADatabase.User,
			config.GetConfiguration().QADatabase.Password,
			config.GetConfiguration().QADatabase.DBname,
			config.GetConfiguration().QADatabase.Port), nil
	} else if config.GetConfiguration().Server.Environment == "DEV" {
		return FormatDSN(
			config.GetConfiguration().DEVDatabase.User,
			config.GetConfiguration().DEVDatabase.Password,
			config.GetConfiguration().DEVDatabase.DBname,
			config.GetConfiguration().DEVDatabase.Port), nil
	} else if config.GetConfiguration().Server.Environment == "PROD" {
		return FormatDSN(
			config.GetConfiguration().PRODdatabase.User,
			config.GetConfiguration().PRODdatabase.Password,
			config.GetConfiguration().PRODdatabase.DBname,
			config.GetConfiguration().PRODdatabase.Port), nil
	}

	return "", errors.New("Server config dont have a valid environment")
}

func FormatDSN(user string, password string, dbname string, port int) string {
	return "user=" + user + " password=" + password + " dbname=" + dbname + " port=" + strconv.Itoa(port) + " sslmode=disable TimeZone=America/Sao_Paulo"
}
