package main

import (
	"fmt"
	"project/HINTTASK5/database"
	"project/HINTTASK5/router"
)

func main() {
	// Connect to the database
	if err := database.ConnectDB(); err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	fmt.Println("Connected to the database successfully")

	r := router.SetupRouter()

	r.Run(":8080")
}
