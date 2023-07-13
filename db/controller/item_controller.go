package controller

import (
	"simple-user-inventory2/db/models"

	"gorm.io/gorm"
)

type ItemController struct {
	db *gorm.DB
}

func NewItemController(db *gorm.DB) ItemController {
	return ItemController{db: db}
}

func (c ItemController) Create(name string) (*models.ItemData, error) {
	item := models.NewItem(name)
	result := c.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item.ItemData, nil
}

func (c ItemController) Read(uuid string) (*models.Item, error) {
	item := &models.Item{}
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
