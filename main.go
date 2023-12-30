package main

import (
	"fmt"
	"project/HINTTASK5/database"
	"project/HINTTASK5/router"
	"project/HINTTASK5/config"
)

func main() {
	// Connect to the database
	if err := database.ConnectDB(); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	fmt.Println("Connected to the database successfully")

	r := router.SetupRouter()
	
	// Get the application port from the config package
	appPort := config.GetAppPort()
	r.Run(":" + appPort)
}
