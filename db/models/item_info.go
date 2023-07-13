package models

import "gorm.io/gorm"

type ItemInfoData struct {
	Amount   uint   `json:"amount" gorm:"not null"`
	UserUuid string `json:"user_uuid" gorm:"not null;index;size:64"`
	ItemUuid string `json:"item_uuid" gorm:"not null;size:64"`
}

type ItemInfo struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	ItemID uint `gorm:"not null"`

	ItemInfoData
}
