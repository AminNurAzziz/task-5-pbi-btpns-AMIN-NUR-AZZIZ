package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"project/HINTTASK5/models"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// Load environment variables from .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func ConnectDB() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	// Establish database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	DB = db

	// Run database migration if tables do not exist
	if err := MigrateDB(); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("Error getting SQL DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		fmt.Println("Error closing database connection:", err)
	}
}

func MigrateDB() error {
	if err := DB.AutoMigrate(&models.User{}, &models.Photo{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}

func TestConnection() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error getting SQL DB: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	return nil
}
