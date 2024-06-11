package user

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	"example.com/larkiee/interview/db"
	"example.com/larkiee/interview/models"
)

type user struct {}

type addressJson struct {
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type userJson struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Addresses []addressJson `json:"addresses"`
}

func createFromFile(filename string) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var users []userJson
	err = json.Unmarshal(buf, &users)
	if err != nil {
		return err
	}
	userStream := make(chan userJson)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ju := range userStream {
			processUser(ju)	
		}
	}()
	for ind, jsonU := range users {
		userStream <- jsonU
		if ind == len(users) - 1 {
			close(userStream)
		}
	}
	wg.Wait()
	return nil
}

func processUser(ju userJson) {
	splitedName := strings.Split(ju.Name, " ")
			dbUser := models.User{
				UserID: ju.ID,
				FirstName: splitedName[0],
				LastName: splitedName[1],
				Email: ju.Email,
				PhoneNumber: ju.PhoneNumber,
			}
			var addrs []models.Address 
			for _, addr := range ju.Addresses {
				addrs = append(addrs, models.Address{
					Street: addr.Street,
					City: addr.City,
					State: addr.State,
					ZipCode: addr.ZipCode,
					Country: addr.Country,
				})
			}
			dbUser.Addresses = addrs
			tx := db.DB.Create(&dbUser)
			if tx.Error != nil {
				log.Fatal(tx.Error)
			}
}

func (u *user) GetUser(id string) (*models.User, error) {
	var res models.User
	tx := db.DB.Where("user_id = ?", id).Preload("Addresses").Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if res.UserID == "" {
		return nil, errors.New("not found")
	}
	return &res, nil
}


func New() *user {
	return &user{}
}