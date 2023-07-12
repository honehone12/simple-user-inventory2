package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserData struct {
	Name string `json:"name" gorm:"not null;size:128"`
	Uuid string `json:"uuid" gorm:"unique;uniqueIndex;not null;size:64"`
}

type User struct {
	gorm.Model

	UserData

	ItemInfos []*ItemInfo
}

func NewUser(name string) *User {
	return &User{
		UserData: UserData{
			Name: name,
			Uuid: uuid.NewString(),
		},
	}
}
