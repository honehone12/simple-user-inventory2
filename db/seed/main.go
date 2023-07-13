package main

import (
	"encoding/json"
	"log"
	"os"
	"simple-user-inventory2/db"
	"simple-user-inventory2/db/models"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}

	userBytes, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	itemBytes, err := os.ReadFile("items.json")
	if err != nil {
		log.Fatal(err)
	}

	users := make([]models.UserData, 0)
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		log.Fatal(err)
	}
	items := make([]models.ItemData, 0)
	err = json.Unmarshal(itemBytes, &items)
	if err != nil {
		log.Fatal(err)
	}

	orm, err := db.NewOrm()
	if err != nil {
		log.Fatal(err)
	}
	userCtrl := orm.User()
	userLen := len(users)
	userUuids := make([]string, userLen)
	itemCtrl := orm.Item()
	itemLen := len(items)
	itemUuids := make([]string, itemLen)

	for i := 0; i < userLen; i++ {
		uData, err := userCtrl.Create(users[i].Name)
		if err != nil {
			log.Fatal(err)
		}
		userUuids[i] = uData.Uuid
	}
	for i := 0; i < itemLen; i++ {
		iData, err := itemCtrl.Create(items[i].Name)
		if err != nil {
			log.Fatal(err)
		}
		itemUuids[i] = iData.Uuid
	}

	itemInfoCtrl := orm.ItemInfo()
	for i := 0; i < userLen; i++ {
		for j := 0; j < itemLen; j++ {
			err := itemInfoCtrl.Create(userUuids[i], itemUuids[j], 1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("seed done")
}
