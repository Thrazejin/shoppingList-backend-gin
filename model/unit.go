package model

import (
	"errors"

	"gorm.io/gorm"
)

/// --------------
/// Unit functions
/// --------------

func GetUnitById(user User, unit *Unit, id int) error {

	if err := DB.Where("id = ?", id).First(unit).Error; err != nil {
		return err
	}

	if unit.UserID != user.ID {
		return errors.New("This Unit does not belong to you")
	}

	return nil
}

func (unit *Unit) ParseOutput() *UnitOutput {

	return &UnitOutput{ID: unit.ID, Name: unit.Name, ShortName: unit.ShortName}
}

func FindAllUnities(user User) ([]Unit, error) {
	var unities []Unit

	if err := DB.Where("userID = ?", user.ID).Find(&unities).Error; err != nil {
		return make([]Unit, 0), err
	}

	return unities, nil
}

func MapForOutput(array []Unit, f func(Unit) UnitOutput) []UnitOutput {
	vsm := make([]UnitOutput, len(array))
	for i, v := range array {
		vsm[i] = f(v)
	}
	return vsm
}

func ParseArrayOutput(array []Unit) []UnitOutput {
	return MapForOutput(array,
		func(u Unit) UnitOutput {
			return *u.ParseOutput()
		})
}

func (unit *Unit) Delete() error {
	return DB.Delete(unit).Error
}

/// -------------------
/// UnitInput functions
/// -------------------

func (input *UnitInput) FindIfDoesntExist(user User) (Unit, error) {
	var unit Unit

	if result := DB.Where("name = ? and userId = ?", input.Name, user.ID).First(&unit); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return unit, errors.New("Unit 'name' already exists")
	}

	if result := DB.Where("shortName = ? and userId = ?", input.ShortName, user.ID).First(&unit); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return unit, errors.New("Unit 'shortName' already exists")
	}

	return Unit{}, nil
}

func (input *UnitInput) Save(user User) (*Unit, error) {
	unit := input.Parse(user)

	if err := DB.Create(unit).Error; err != nil {
		return &Unit{}, err
	}

	return unit, nil
}

func (input *UnitInput) Parse(user User) *Unit {
	unit := Unit{User: user, UserID: user.ID, Name: input.Name, ShortName: input.ShortName}

	return &unit
}

/// -------------------------
/// UnitUpdateInput functions
/// -------------------------

func (input *UnitUpdateInput) FindIfDoesntExist(user User, unitIn Unit) (Unit, error) {
	var unit Unit

	if result := DB.Where("name = ? and userID = ? and ID <> ?", input.Name, user.ID, unitIn.ID).First(&unit); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return unit, errors.New("Unit 'name' already exists")
	}

	if result := DB.Where("shortName = ? and userID = ? and ID <> ?", input.ShortName, user.ID, unitIn.ID).First(&unit); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return unit, errors.New("Unit 'shortName' already exists")
	}

	return Unit{}, nil
}

func (input *UnitUpdateInput) Update(unit Unit) (*Unit, error) {
	DB.Model(&unit).Updates(input)

	return &unit, nil
}

func (input *UnitUpdateInput) ParseUnit(id int) (*Unit, error) {
	var unit Unit

	if err := DB.Where("id = ?", id).First(&unit).Error; err != nil {
		return &Unit{}, err
	}

	DB.Model(&unit).Updates(input)

	return &unit, nil
}

type Unit struct {
	gorm.Model
	Name      string `gorm:"size:100;not null;unique"`
	ShortName string `gorm:"size:10;not null;unique"`
	UserID    uint
	User      User `gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UnitInput struct {
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"shortName" binding:"required"`
}

type UnitUpdateInput struct {
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type UnitOutput struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}
