package globals

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	INIT_FILE_NAME                            = "users_data.json"
	PG_DB, PG_USER, PG_PASSWORD, SERVICE_PORT string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("fail to load .env file")
	}
	PG_DB = os.Getenv("PG_DB")
	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	SERVICE_PORT = os.Getenv("SERVICE_PORT")

	if PG_DB == "" || PG_USER == "" || PG_PASSWORD == "" || SERVICE_PORT == "" {
		log.Fatal("please set all env varibles in .env file located in root directory")
	}
}
