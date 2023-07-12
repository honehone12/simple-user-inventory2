package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"simple-user-inventory2/db"
	"simple-user-inventory2/db/models"
	"sync"

	"github.com/google/uuid"
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

	wg := sync.WaitGroup{}

	for j := 0; j < 1; j++ {
		orm, err := db.NewOrm()
		if err != nil {
			log.Fatal(err)
		}
		userCtrl := orm.User()
		userLen := len(users)
		itemCtrl := orm.Item()
		itemLen := len(items)

		wg.Add(1)

		go func() {
			for i := 0; i < userLen; i++ {
				userCtrl.Create(fmt.Sprintf("%s%s", users[i].Name, uuid.NewString()))
			}
			for i := 0; i < itemLen; i++ {
				itemCtrl.Create(fmt.Sprintf("%s%s", items[i].Name, uuid.NewString()))
			}
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("seed done")
}
