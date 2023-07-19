package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemMasterData struct {
	Name string `json:"name" gorm:"unique;not null;size:128"`
	Uuid string `json:"uuid" gorm:"unique;uniqueIndex;not null;size:64"`
}

type ItemMaster struct {
	gorm.Model

	ItemMasterData
}

func NewItemMaster(name string) *ItemMaster {
	return &ItemMaster{
		ItemMasterData: ItemMasterData{
			Name: name,
			Uuid: uuid.NewString(),
		},
	}
}
