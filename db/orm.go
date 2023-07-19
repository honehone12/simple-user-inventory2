package db

import (
	"errors"
	"log"
	"os"
	"simple-user-inventory2/db/controller"
	"simple-user-inventory2/db/models"

	"github.com/redis/rueidis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orm struct {
	db *gorm.DB
	kv rueidis.Client
}

func NewOrm() (Orm, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		return Orm{nil, nil}, errors.New("env param POSTGRES_DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Orm{nil, nil}, err
	}

	if err = db.AutoMigrate(
		&models.UserMaster{}, &models.ItemMaster{}, &models.ItemState{},
	); err != nil {
		return Orm{nil, nil}, err
	}

	log.Println("established new db connection")

	addr := os.Getenv("REDIS_ADDR")
	if len(addr) == 0 {
		return Orm{nil, nil}, errors.New("env param REDIS_ADDR is empty")
	}

	kv, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{addr},
	})
	if err != nil {
		return Orm{nil, nil}, err
	}

	log.Println("established new kv connection")

	return Orm{
		db: db,
		kv: kv,
	}, nil
}

func (orm Orm) User() controller.UserMasterController {
	return controller.NewUserMasterController(orm.db)
}

func (orm Orm) Item() controller.ItemMasterController {
	return controller.NewItemMasterController(orm.db)
}

func (orm Orm) ItemState() controller.ItemStateController {
	return controller.NewItemStateController(orm.kv)
}
