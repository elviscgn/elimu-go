// @title           Elimu Portal API
// @version         0.1
// @description     Core backend api for elimu student lms
// @contact.name    API Support
// @contact.email   iwill@create.later

// @host            localhost:8080
// @BasePath        /api

// @securityDefinitions.apiKey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Enter your Google OAuth token

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"elimu-go/internal/handlers"
	"log"
	"os"

	_ "elimu-go/docs" // Add this line (docs folder will be auto-generated)
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main godoc
// @Summary      Start the API server
// @Description  Initializes and runs the Student Portal API on port 8080
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
	}

	log.Printf("Starting :%s", port)
	r.Run(":" + port)
}
