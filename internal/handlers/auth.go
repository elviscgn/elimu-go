package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
	sessions    = make(map[string]*User)
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	GoogleID string `json:"google_id"`
}

// loads oauth config from environment
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
// @Description  Redirects to Google OAuth authentication page
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      307
// @Router       /login [get]
func GoogleLogin(c *gin.Context) {
	state := "state_123"
	authURL := oauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// GoogleCallback godoc
// @Summary      OAuth callback
// @Description  Handles Google OAuth callback and creates user session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        code   query     string  true  "Authorization code from Google"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /callback [get]
func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code"})
		return
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed"})
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
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
	sessions[sessionID] = user
	c.SetCookie("session_id", sessionID, 3600, "/", "localhost", false, true)

	fmt.Print("Logged in")
	log.Printf("User: %s", user.Name)
	log.Printf("Email: %s", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"user":    user,
	})
}

// GetCurrentUser godoc
// @Summary      Get current user
// @Description  Returns information about currently logged in user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200    {object}  map[string]interface{}
// @Failure      401    {object}  map[string]interface{}
// @Router       /me [get]
func GetCurrentUser(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	user, exists := sessions[sessionID]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Logout godoc
// @Summary      Logout user
// @Description  Clears user session and logs them out
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /logout [get]
func Logout(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err == nil {
		delete(sessions, sessionID)
	}
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
