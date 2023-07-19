package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMasterData struct {
	Name string `json:"name" gorm:"not null;size:128"`
	Uuid string `json:"uuid" gorm:"unique;uniqueIndex;not null;size:64"`
}

type UserMaster struct {
	gorm.Model

	UserMasterData
}

func NewUserMaster(name string) *UserMaster {
	return &UserMaster{
		UserMasterData: UserMasterData{
			Name: name,
			Uuid: uuid.NewString(),
		},
	}
}
