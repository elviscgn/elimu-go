package middleware

import (
	"net/http"
	"sync"

	"elimu-go/internal/models"

	"github.com/gin-gonic/gin"
)

type contextKey string

const CurrentUserKey contextKey = "current_user"

func RequireLogin(sessions *sync.Map) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
			c.Abort()
			return
		}

		user, ok := sessions.Load(sessionID)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			c.Abort()
			return
		}

		u, ok := user.(*models.User)
		println("User", u.ID, u.Role, u.Email, "Wha", ok)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid session user"})
			c.Abort()
			return
		}

		c.Set(string(CurrentUserKey), u)
		c.Next()
	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	roleSet := make(map[string]struct{}, len(allowedRoles))
	for _, r := range allowedRoles {
		roleSet[r] = struct{}{}
	}

	return func(c *gin.Context) {
		u, exists := c.Get(string(CurrentUserKey))
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user missing from context"})
			c.Abort()
			return
		}

		user, ok := u.(*models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user in context"})
			c.Abort()
			return
		}

		if _, allowed := roleSet[user.Role]; !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
