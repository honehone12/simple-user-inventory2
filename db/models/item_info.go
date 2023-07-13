package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemInfoData struct {
	Amount   uint   `json:"amount" gorm:"not null"`
	Uuid     string `json:"uuid" gorm:"not null;unique;uniqueIndex;size:64"`
	UserUuid string `json:"user_uuid" gorm:"not null;index;size:64"`
	ItemUuid string `json:"item_uuid" gorm:"not null;size:64"`
}

type ItemInfo struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	ItemID uint `gorm:"not null"`

	ItemInfoData
}

func NewItemInfo(
	userId uint,
	userUuid string,
	itemId uint,
	itemUuid string,
	amount uint,
) *ItemInfo {
	return &ItemInfo{
		UserID: userId,
		ItemID: itemId,
		ItemInfoData: ItemInfoData{
			Amount:   amount,
			Uuid:     uuid.NewString(),
			UserUuid: userUuid,
			ItemUuid: itemUuid,
		},
	}
}
