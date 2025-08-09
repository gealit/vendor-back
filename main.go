package main

import (
	"io"
	"log"
	"os"

	"main/middlewares"
	"main/models"
	"main/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file...")
	}
}

func main() {
	f, _ := os.Create("vendor.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Apply middleware
	router.Use(middlewares.CORSMiddleware())
	// router.Use(middlewares.AuthMiddleware())

	// Load the routes
	routes.AuthRoutes(router)

	// Run the server
	router.Run(":8080")
}
