package config

import (
	"bookstore/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.Author{},
		&models.Category{},
		&models.Book{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	log.Println("Connected successfully")

}
