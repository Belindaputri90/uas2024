package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"uas2024/database"
	"uas2024/models"
)

func GetProducts(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var products []models.Product
	database.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	database.DB.First(&product, id)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&productInput)
	c.JSON(http.StatusOK, gin.H{"data": productInput})
}

func UpdateProduct(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var productInput models.Product
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	database.DB.First(&product, id)
	product.Name = productInput.Name
	product.Price = productInput.Price
	database.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
	// Pastikan user telah login
	if !IsUserLoggedIn(c.Request) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, gin.H{"data": "Product deleted"})
}
