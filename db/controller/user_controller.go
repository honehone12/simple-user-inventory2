package controller

import (
	"simple-user-inventory2/db/models"

	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{db: db}
}

func (c UserController) Create(name string) error {
	user := models.NewUser(name)
	result := c.db.Create(user)
	return result.Error
}

func (c UserController) Read(uuid string) (*models.User, error) {
	user := &models.User{}
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
