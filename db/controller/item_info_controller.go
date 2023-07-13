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

		info := models.NewItemInfo(
			user.ID,
			userUuid,
			item.ID,
			itemUuid,
			1,
		)
		result = tx.Create(info)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (c ItemInfoController) ReadAllUserItems(userUuid string) ([]models.ItemInfo, error) {
	itemInfos := make([]models.ItemInfo, 0)
	result := c.db.
		Select(
			"ID", "CreatedAt", "UpdatedAt", "Uuid", "UserID", "ItemID",
			"Amount", "UserUuid", "ItemUuid",
		).
		Where("user_uuid = ?", userUuid).
		Find(&itemInfos)
	if result.Error != nil {
		return nil, result.Error
	}
	return itemInfos, nil
}

func (c ItemInfoController) UpdateAmount(uuid string, amount uint) (*models.ItemInfoData, error) {
	itemInfo := models.ItemInfo{ItemInfoData: models.ItemInfoData{Amount: amount}}
	// keep itemInfo's pk = 0, otherwise pk will be condition
	result := c.db.
		Model(itemInfo).
		Where("uuid = ?", uuid).
		Updates(itemInfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &itemInfo.ItemInfoData, nil
}
