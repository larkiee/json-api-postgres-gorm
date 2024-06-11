package db

import (
	"fmt"
	"log"

	globals "example.com/larkiee/interview/Globals"
	"example.com/larkiee/interview/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init(){
	dsn := fmt.Sprintf(
	"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
	globals.PG_USER,
	globals.PG_PASSWORD,
	globals.PG_DB,

)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed postgres connection")
	}
	err = db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}