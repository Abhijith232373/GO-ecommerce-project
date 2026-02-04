package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")),&gorm.Config{})
	if err!=nil{
		log.Fatal("Database connection failed !")
	}
	DB=db
	log.Println("db connected !")
}