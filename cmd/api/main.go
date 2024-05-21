package main

import (
	"dating-app/internal/models"
	"dating-app/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDB()

	router := gin.Default()

	api.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}
