// login_handlers.go
package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"uas2024/database"
)

// Variabel global untuk menyimpan secret key
var globalSecretKey string

// LoginRequest struct to capture login data from the request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse struct to respond to login results
type LoginResponse struct {
	Token string `json:"token"`
}

// LogoutHandler for logout
func LogoutHandler(c *gin.Context) {
	// Clear the session_token cookie
	c.SetCookie("session_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// LoginHandler handles login requests
func LoginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := database.GetUserByUsername(loginRequest.Username)
	if err != nil || user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, secretKey, err := generateJWTToken(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Successful login, send token as the response
	response := LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{"data": response})

	// Store the dynamic secret key globally
	globalSecretKey = secretKey
}

// Generate JWT token
func generateJWTToken(c *gin.Context, userID int) (string, string, error) {
	// Convert user.ID to uint
	uintUserID := uint(userID)

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": uintUserID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	})

	// Generate dynamic secret key
	secretKey, err := generateSecretKey()
	if err != nil {
		return "", "", err
	}

	// Use the dynamic secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	// Set the session_token cookie
	c.SetCookie("session_token", tokenString, 60*60*24, "/", "localhost", false, true)

	return tokenString, secretKey, nil
}

// Generate secret key dynamically
func generateSecretKey() (string, error) {
	// Menentukan panjang secret key yang diinginkan
	keyLength := 32 // Panjang dalam byte

	// Menggunakan crypto/rand untuk menghasilkan random bytes
	randomBytes := make([]byte, keyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Mengonversi random bytes ke dalam format string yang aman
	secretKey := base64.StdEncoding.EncodeToString(randomBytes)

	return secretKey, nil
}
