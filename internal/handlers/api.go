package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Student Portal API",
		"version": "1.0.0",
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "Healthy",
		"timestamp": time.Now().Unix(),
	})
}

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
