package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"uas2024/database"
	"uas2024/models"
)

func GetUsers(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var users []models.User
	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	database.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&userInput)
	c.JSON(http.StatusOK, gin.H{"data": userInput})
}

func UpdateUser(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	database.DB.First(&user, id)
	user.Username = userInput.Username
	user.Password = userInput.Password
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.User{}, id)

	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}