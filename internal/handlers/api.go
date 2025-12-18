package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @Summary      Welcome message
// @Description  Returns API welcome message with version info
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
// @Example      Response
//
//	{
//	  "message": "Elimu Portal API",
//	  "version": "1.0.0"
//	}
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Elimu Portal API",
		"version": "1.0.0",
	})
}

// HealthCheck godoc
// @Summary health check
// @Description  Check if API is running properly
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
// @Example      Response
//
//	{
//	  "status": "Healthy",
//	  "timestamp": 1702867200
//	}
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "Healthy",
		"timestamp": time.Now().Unix(),
	})
}

// RandomEndpoint godoc
// @Summary      Random student fact
// @Description  Returns a random educational fact
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /random [get]
// @Example      Response
//
//	{
//	  "fact": "Students who study with music retain information better."
//	}
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

// DebugInfo godoc
// @Summary      Debug information
// @Description  Returns detailed request debug information
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /debug [get]
// @Example      Response
//
//	{
//	  "method": "GET",
//	  "path": "/api/debug",
//	  "user_agent": "Mozilla/5.0...",
//	  "ip": "192.168.1.100",
//	  "timestamp": "2024-12-17T10:30:00Z"
//	}
func DebugInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"ip":         c.ClientIP(),
		"timestamp":  time.Now().Format(time.RFC3339),
	})
}
