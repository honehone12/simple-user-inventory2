package db

import (
	"errors"
	"log"
	"os"
	"simple-user-inventory2/db/controller"
	"simple-user-inventory2/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orm struct {
	db *gorm.DB
}

func NewOrm() (Orm, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		return Orm{nil}, errors.New("env param POSTGRES_DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Orm{nil}, err
	}

	if err = db.AutoMigrate(
		&models.User{}, &models.Item{}, &models.ItemInfo{},
	); err != nil {
		return Orm{nil}, err
	}

	log.Println("established new database connection")
	return Orm{db: db}, nil
}

func (orm Orm) User() controller.UserController {
	return controller.NewUserController(orm.db)
}

func (orm Orm) Item() controller.ItemController {
	return controller.NewItemController(orm.db)
}

func (orm Orm) ItemInfo() controller.ItemInfoController {
	return controller.NewItemInfoController(orm.db)
}
