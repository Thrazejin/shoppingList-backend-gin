package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string `gorm:"size:100;not null;unique"`
	ShortName string `gorm:"size:10;not null;unique"`
	UserID    uint
	User      AppUser `gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
