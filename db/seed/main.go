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
		log.Fatalln(err)
	}

	userBytes, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalln(err)
	}
	itemBytes, err := os.ReadFile("items.json")
	if err != nil {
		log.Fatalln(err)
	}

	users := make([]models.UserMasterData, 0)
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		log.Fatalln(err)
	}
	items := make([]models.ItemMasterData, 0)
	err = json.Unmarshal(itemBytes, &items)
	if err != nil {
		log.Fatalln(err)
	}

	orm, err := db.NewOrm()
	if err != nil {
		log.Fatalln(err)
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
			log.Fatalln(err)
		}
		userUuids[i] = uData.Uuid
	}
	for i := 0; i < itemLen; i++ {
		iData, err := itemCtrl.Create(items[i].Name)
		if err != nil {
			log.Fatalln(err)
		}
		itemUuids[i] = iData.Uuid
	}

	itemStateCtrl := orm.ItemState()
	for i := 0; i < userLen; i++ {
		for j := 0; j < itemLen; j++ {
			err := itemStateCtrl.Create(userUuids[i], itemUuids[j], 1)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	numTX := 0

	for i := 0; i < userLen; i++ {
		states := []*[]models.ItemState{}
		err := itemStateCtrl.ReadAllUserItems(userUuids[i], &states)
		if err != nil {
			log.Fatalln(err)
		}

		ptr := []uint{}
		for j := 0; j < itemLen; j++ {
			for k := 0; k < 100; k++ {
				err := itemStateCtrl.UpdateAmount((*states[j])[0].Uuid, 1, &ptr)
				if err != nil {
					log.Fatalln(err)
				}

				numTX++
				(*states[j])[0].Amount = ptr[0]
				log.Printf("%#v\n", states[j])
			}
		}
	}

	log.Printf("seed done, %d txs", numTX)
}
