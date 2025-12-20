// @title           Elimu Portal API
// @version         0.1
// @description     Core backend api for elimu student lms
// @contact.name    API Support
// @contact.email   iwill@create.later
// @host            localhost:8080
// @BasePath        /api
// @tag.name        Authentication
// @tag.description User authentication and session management via Google OAuth
// @tag.name        General
// @tag.description Core API endpoints, health checks, and debug utilities
package main

import (
	"log"
	"os"

	_ "elimu-go/docs"
	"elimu-go/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		api.GET("/admin/overview", handlers.AdminOverview)
	}

	log.Printf("Starting :%s", port)
	r.Run(":" + port)
}
