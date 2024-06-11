package user

import (
	"log"

	"example.com/larkiee/interview/Globals"
	"example.com/larkiee/interview/db"
	"example.com/larkiee/interview/models"
)

func init() {
	var count int64
	db.DB.Model(models.User{}).Count(&count)
	if count == 10000 {
		log.Println("already initialized")
		return
	}
	log.Println("initializing database ...")
	err := createFromFile(globals.INIT_FILE_NAME)
	if err != nil {
		log.Fatal("failed init database")
	}
	log.Println("finish initialization")
}
