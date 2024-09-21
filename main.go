package main

import (
	"log"
	"pr/models"
	"pr/routes"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	LoadEnv()
	models.ConnectDatabase()
	routes.SetupRoutes()
}
