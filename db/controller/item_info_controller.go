package controller

import (
	"simple-user-inventory2/db/models"

	"gorm.io/gorm"
)

type ItemInfoController struct {
	db *gorm.DB
}

func NewItemInfoController(db *gorm.DB) ItemInfoController {
	return ItemInfoController{db: db}
}

func (c ItemInfoController) Create(userUuid string, itemUuid string, amount uint) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		user := &models.User{}
		result := tx.Select("ID", "Name").Where("uuid = ?", userUuid).Take(user)
		if result.Error != nil {
			return nil
		}
		item := &models.Item{}
		result = tx.Select("ID", "Name").Where("uuid = ?", itemUuid).Take(item)
		if result.Error != nil {
			return nil
		}

		info := &models.ItemInfo{
			UserID: user.ID,
			ItemID: item.ID,
			ItemInfoData: models.ItemInfoData{
				UserUuid: userUuid,
				ItemUuid: itemUuid,
				Amount:   amount,
			},
		}
		result = tx.Create(info)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
