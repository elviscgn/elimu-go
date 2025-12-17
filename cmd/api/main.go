package main

import (
	"elimu-go/internal/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}
	log.Println("ENV loaded")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", handlers.Welcome)
		api.GET("/health", handlers.HealthCheck)
		api.GET("/random", handlers.RandomEndpoint)
		api.GET("/debug", handlers.DebugInfo)
		api.GET("/login", handlers.GoogleLogin)
		api.GET("/callback", handlers.GoogleCallback)
		api.GET("/me", handlers.GetCurrentUser)
		api.GET("/logout", handlers.Logout)
	}

	log.Printf("Starting :%s", port)
	r.Run(":" + port)
}
