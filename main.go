package main

import (
	"log"
	"student-management/config"
	"student-management/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := routes.SetupRouter()
	router.Run()
}
