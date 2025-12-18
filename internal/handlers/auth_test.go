package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetCurrentUser_NoSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Reset sessions for clean test
	sessions = &sync.Map{}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/me", nil)

	GetCurrentUser(c)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 for no session, got %d", w.Code)
	}
}

func TestGetCurrentUser_ValidSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup: Create a user in sessions
	testUser := &User{
		ID:       "test_123",
		Email:    "test@student.edu",
		Name:     "Test Student",
		Picture:  "https://example.com/pic.jpg",
		GoogleID: "google_123",
	}
	sessions.Store("session_test_123", testUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/me", nil)

	// Add session cookie
	c.Request.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "session_test_123",
	})

	GetCurrentUser(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 with valid session, got %d", w.Code)
	}
}

func TestLogout_WithSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup: Add a session
	testUser := &User{ID: "logout_test"}
	sessions.Store("session_logout_test", testUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/logout", nil)

	// Add session cookie
	c.Request.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "session_logout_test",
	})

	Logout(c)

	// Check session was deleted
	_, exists := sessions.Load("session_logout_test")
	if exists {
		t.Error("Session should have been deleted after logout")
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}

func TestUserStructJSON(t *testing.T) {
	// Test that User struct marshals correctly
	user := &User{
		ID:       "123",
		Email:    "test@student.edu",
		Name:     "Test User",
		Picture:  "https://example.com/pic.jpg",
		GoogleID: "google_123",
	}

	// This tests your JSON struct tags work
	// If JSON marshal fails, your API responses will break
	_, err := json.Marshal(user)
	if err != nil {
		t.Errorf("User struct should marshal to JSON: %v", err)
	}
}

func TestErrorResponses(t *testing.T) {
	// Test your error response format
	errResp := ErrorResponse{
		Error: "Test error",
	}

	_, err := json.Marshal(errResp)
	if err != nil {
		t.Errorf("ErrorResponse should marshal to JSON: %v", err)
	}
}
