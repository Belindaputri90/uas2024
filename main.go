package main

import (
	"github.com/gin-gonic/gin"
	"uas2024/database"
	"uas2024/handlers"
)

func main() {
	router := gin.Default()

	// Inisialisasi koneksi ke database
	database.InitDB()

	// Middleware otentikasi
	router.Use(func(c *gin.Context) {
		if true {
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	})

	// Routes for Users
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Routes for Products
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProduct)
	router.POST("/products", handlers.CreateProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	// Route for Login
	router.POST("/login", handlers.LoginHandler)

	// Route for Logout (memerlukan otentikasi)
	router.POST("/logout", handlers.LogoutHandler)

	// Jalankan server
	router.Run(":8080")
}
