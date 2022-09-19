package main

import (
	"go-gorm/initializers"
	"go-gorm/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
