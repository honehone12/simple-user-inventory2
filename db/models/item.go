package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemData struct {
	Name string `json:"name" gorm:"unique;not null;size:128"`
	Uuid string `json:"uuid" gorm:"unique;uniqueIndex;not null;size:64"`
}

type Item struct {
	gorm.Model

	ItemData

	ItemInfos []*ItemInfo
}

func NewItem(name string) *Item {
	return &Item{
		ItemData: ItemData{
			Name: name,
			Uuid: uuid.NewString(),
		},
	}
}
