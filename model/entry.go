package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Content   string `gorm:"type:text" json:"content"`
	AppUserID uint
	User      AppUser `gorm:"not null;foreignKey:AppUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EntryList struct {
	entries []Entry
}

func (entry *Entry) Save() (*Entry, error) {

	err := DB.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func (user *AppUser) GetEntryList() []Entry {
	return user.Entries
	/*return EntryList{entries: user.Entries}*/
}
