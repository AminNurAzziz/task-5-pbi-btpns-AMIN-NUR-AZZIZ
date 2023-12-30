package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var jwtSecret string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	jwtSecret = os.Getenv("JWT_SECRET")
}

func GetJWTSecret() string {
	return jwtSecret
}
