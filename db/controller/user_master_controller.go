package controller

import (
	"simple-user-inventory2/db/models"

	"gorm.io/gorm"
)

type UserMasterController struct {
	db *gorm.DB
}

func NewUserMasterController(db *gorm.DB) UserMasterController {
	return UserMasterController{db: db}
}

func (c UserMasterController) Create(name string) (*models.UserMasterData, error) {
	user := models.NewUserMaster(name)
	result := c.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.UserMasterData, nil
}

func (c UserMasterController) Read(uuid string) (*models.UserMaster, error) {
	user := &models.UserMaster{}
	result := c.db.
		Select(
			"ID", "CreatedAt", "UpdatedAt", "Name", "Uuid",
		).
		Where("uuid = ?", uuid).
		Take(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
