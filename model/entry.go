package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
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

func (user *User) GetEntryList() []Entry {
	return user.Entries
	/*return EntryList{entries: user.Entries}*/
}
