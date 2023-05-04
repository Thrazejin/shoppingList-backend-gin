package model

import (
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model
	Name     string `gorm:"size:100;not null" json:"name"`
	Username string `gorm:"size:20;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Entries  []Entry
}

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *AppUser) Save() error {
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *AppUser) BeforeSave(*gorm.DB) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *AppUser) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (input *UserInput) FindIfDoesntExist() (AppUser, error) {
	var user AppUser

	if err := DB.Where("username = ?", input.Username).First(&user).Error; err == nil {
		return user, errors.New("Unit 'name' already exists")
	}

	return AppUser{}, nil
}

func FindUserByUsername(username string) (AppUser, error) {
	var user AppUser
	err := DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return AppUser{}, err
	}
	return user, nil
}

func FindUserById(id uint) (AppUser, error) {
	var user AppUser
	err := DB.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return AppUser{}, err
	}
	return user, nil
}
