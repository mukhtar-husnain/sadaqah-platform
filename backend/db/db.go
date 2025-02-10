package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mukhtar-husnain/sadaqah-platform/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file.")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open DB: ", err)
	}

	// err = DB.Ping()
	// if err != nil {
	// 	log.Fatal("Database connection is not alive:", err)
	// }
	DB.AutoMigrate(&models.User{}, &models.VolunteerRequest{}, &models.Volunteer{}, &models.Feedback{})

	fmt.Println("✅ Connected to PostgreSQL database successfully!")
}
