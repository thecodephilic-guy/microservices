package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thecodephilic-guy/user-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //creating and exporting the instance of the database connection

func ConnectDB() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Successfully Connected To The Database!")
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{}) // Creates users table if not exists
}
