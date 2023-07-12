package models

import "gorm.io/gorm"

type ItemInfoData struct {
	Amount uint `json:"amount" gorm:"not null"`
}

type ItemInfo struct {
	gorm.Model

	UserID uint
	ItemID uint

	ItemInfoData
}
