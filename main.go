package main

import (
	"gamestatsticregistry/mrgrease.com/db"
	"gamestatsticregistry/mrgrease.com/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file in the current directory
	godotenv.Load()
	server := gin.Default()
	routes.RegisterRoutes(server)
	db.InitCurrentDb()

	_, portPresent := os.LookupEnv("PORT")

	if !portPresent {
		panic("Incomplete Configuration Error: No port provided!")
	}

	server.Run(os.Getenv("PORT"))
}
