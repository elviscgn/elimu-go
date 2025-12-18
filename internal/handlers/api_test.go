package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWelcomeEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Welcome(c)

	if w.Code != http.StatusOK {
		t.Errorf("Welcome should return 200, got %d", w.Code)
	}

	if w.Body.String() == "" {
		t.Error("Welcome returned empty response")
	}
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HealthCheck(c)

	if w.Code != http.StatusOK {
		t.Errorf("HealthCheck should return 200, got %d", w.Code)
	}
}
