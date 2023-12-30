package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var jwtSecret string
var appPort string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	jwtSecret = os.Getenv("JWT_SECRET")
	appPort = os.Getenv("APP_PORT")
}

func GetJWTSecret() string {
	return jwtSecret
}

func GetAppPort() string {
	return appPort
}
