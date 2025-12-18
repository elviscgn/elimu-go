package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @Summary      Welcome message
// @Description  Returns a welcome message with API info
// @Tags         general
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Elimu Portal API",
		"version": "1.0.0",
	})
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if the API is running
// @Tags         general
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "Healthy",
		"timestamp": time.Now().Unix(),
	})
}

// RandomEndpoint godoc
// @Summary      Random fact
// @Description  Returns a random student-related fact
// @Tags         general
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /random [get]
func RandomEndpoint(c *gin.Context) {
	facts := []string{
		"Students who study with music retain information better.",
		"The world's oldest university is in Morocco, founded in 859 AD.",
		"Finland doesn't give students standardized tests until college.",
	}
	rand.Seed(time.Now().UnixNano())
	randomFact := facts[rand.Intn(len(facts))]
	c.JSON(http.StatusOK, gin.H{"fact": randomFact})
}

func DebugInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"ip":         c.ClientIP(),
		"timestamp":  time.Now().Format(time.RFC3339),
	})
}
