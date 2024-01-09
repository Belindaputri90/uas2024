package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"uas2024/database"
)

const secretKey = "your-secret-key" // Replace with your actual secret key

// LoginRequest struct to capture login data from the request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse struct to respond to login results
type LoginResponse struct {
	Token string `json:"token"`
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
	token, err := generateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Successful login, send token as the response
	response := LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

// LogoutHandler for logout
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// Generate JWT token
func generateJWTToken(userID int) (string, error) {
	// Convert user.ID to uint
	uintUserID := uint(userID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": uintUserID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	})

	// Use the constant for the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// IsValidToken validates the JWT token
func IsValidToken(tokenString string) bool {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Check for parsing errors
	if err != nil || !token.Valid {
		return false
	}

	return true
}
