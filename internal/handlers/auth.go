package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
	sessions    = &sync.Map{}
)

// User represents an authenticated user
// swagger:model User
type User struct {
	// User's unique ID
	// example: 12345
	ID string `json:"id"`

	// User's email address
	// example: elvischege@student.school.edu
	Email string `json:"email"`

	// User's full name
	// example: Elvis Chege
	Name string `json:"name"`

	// URL to user's profile picture
	// example: https://lh3.googleusercontent.com/a/...
	Picture string `json:"picture"`

	// Google's unique identifier
	// example: 12345678901234567890
	GoogleID string `json:"google_id"`
}

// ErrorResponse represents an API error
// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error message
	// example: Invalid authorization code
	Error string `json:"error"`
}

// LoginResponse represents successful login
// swagger:model LoginResponse
type LoginResponse struct {
	// Success message
	// example: Login successful!
	Message string `json:"message"`

	// Authenticated user data
	User *User `json:"user"`
}

func init() {
	godotenv.Load()

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		log.Println("Client ID empty in init")
	}

	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/api/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GoogleLogin godoc
// @Summary      Start Google OAuth login
// @Description  Redirects to Google OAuth with secure state parameter for authentication
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      307  "Redirect to Google"
// @Failure      500  {object}  ErrorResponse  "Server configuration error"
// @Router       /login [get]
// @Example      Request
// GET /api/login
func GoogleLogin(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	c.SetCookie("oauth_state", state, 300, "/", "", false, true)

	authURL := oauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// GoogleCallback godoc
// @Summary      Handle OAuth callback
// @Description  Processes Google OAuth callback, verifies state, exchanges code for token, and creates user session
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        code   query  string  true  "Authorization code from Google"  example("4/0AX4XfWgYw...")
// @Param        state  query  string  true  "State parameter for CSRF protection"  example("abc123xyz")
// @Success      200    {object}  LoginResponse  "Login successful"
// @Failure      400    {object}  ErrorResponse  "Missing or invalid authorization code"
// @Failure      401    {object}  ErrorResponse  "Invalid OAuth state parameter"
// @Failure      500    {object}  ErrorResponse  "Google API error or server error"
// @Router       /callback [get]
// @Example      Response
//
//	{
//	  "message": "Login successful!",
//	  "user": {
//	    "id": "12345",
//	    "email": "elvischege@student.school.edu",
//	    "name": "Elvis Chege",
//	    "picture": "https://lh3.googleusercontent.com/a/...",
//	    "google_id": "12345678901234567890"
//	  }
//	}
func GoogleCallback(c *gin.Context) {
	receivedState := c.Query("state")
	expectedState, err := c.Cookie("oauth_state")

	if err != nil || receivedState != expectedState {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid state parameter"})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "No code provided"})
		return
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Token exchange failed"})
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var userInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	json.Unmarshal(body, &userInfo)

	user := &User{
		ID:       userInfo.ID,
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Picture:  userInfo.Picture,
		GoogleID: userInfo.ID,
	}

	sessionID := "session_" + user.ID
	sessions.Store(sessionID, user)

	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)
	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	fmt.Print("Logged in")
	log.Printf("User: %s", user.Name)
	log.Printf("Email: %s", user.Email)

	c.JSON(http.StatusOK, LoginResponse{
		Message: "Login successful!",
		User:    user,
	})
}

// GetCurrentUser godoc
// @Summary      Get current user
// @Description  Returns information about currently logged in user
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200  {object}  User  "User data"
// @Failure      401  {object}  ErrorResponse  "Not logged in or session expired"
// @Router       /me [get]
// @Example      Response
//
//	{
//	  "id": "12345",
//	  "email": "elvischege@student.school.edu",
//	  "name": "Elvis Chege",
//	  "picture": "https://lh3.googleusercontent.com/a/...",
//	  "google_id": "12345678901234567890"
//	}
func GetCurrentUser(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Not logged in"})
		return
	}

	value, exists := sessions.Load(sessionID)
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Session expired"})
		return
	}

	user, ok := value.(*User)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Invalid session"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Logout godoc
// @Summary      Logout user
// @Description  Clears user session and logs them out
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /logout [get]
// @Example      Response
//
//	{
//	  "message": "Logged out successfully"
//	}
func Logout(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err == nil {
		sessions.Delete(sessionID)
	}
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
