package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uas2024/models"
)

var DB *gorm.DB

// InitDB initializes the database connection and auto-migrates the tables
func InitDB() {
	var err error
	dataSourceName := "root:@tcp(127.0.0.1:3306)/uas2024?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the tables
	err = DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		panic("failed to auto migrate tables")
	}

	fmt.Println("Connected to database")
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	return user, result.Error
}
