package middleware

import (
	"elimu-go/internal/handlers"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func RequireLogin(sessions *sync.Map) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
			c.Abort()
			return
		}

		user, ok := sessions.Load(sessionID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}

		c.Set("current_user", user)
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get("current_user")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No user in context"})
			c.Abort()
			return
		}

		user := u.(*handlers.User)
		if user.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
