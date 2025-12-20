package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminOverview godoc
// @Summary      Admin overview
// @Description  Lists active sessions and registered users
// @Tags         General
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /admin/overview [get]
func AdminOverview(c *gin.Context) {
	students, err := getAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to load students",
		})
		return
	}

	staff, err := getAllStaff()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to load staff",
		})
		return
	}

	active := getActiveSessions()

	c.JSON(http.StatusOK, gin.H{
		"active_sessions": active,
		"students":        students,
		"staff":           staff,
	})
}
