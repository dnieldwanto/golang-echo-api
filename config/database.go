package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading file configuration: ", err.Error())
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable Timezone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SCHEMA"),
		os.Getenv("DB_TZ"),
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("DB CONNECTED!!")
	return db
}
