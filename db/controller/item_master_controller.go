package controller

import (
	"simple-user-inventory2/db/models"

	"gorm.io/gorm"
)

type ItemMasterController struct {
	db *gorm.DB
}

func NewItemMasterController(db *gorm.DB) ItemMasterController {
	return ItemMasterController{db: db}
}

func (c ItemMasterController) Create(name string) (*models.ItemMasterData, error) {
	item := models.NewItemMaster(name)
	result := c.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item.ItemMasterData, nil
}

func (c ItemMasterController) Read(uuid string) (*models.ItemMaster, error) {
	item := &models.ItemMaster{}
	result := c.db.
		Select(
			"ID", "CreatedAt", "UpdatedAt", "Name", "Uuid",
		).
		Where("uuid = ?", uuid).
		Take(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}
