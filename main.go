package main

import (
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
	server.Run(os.Getenv("PORT"))
}
