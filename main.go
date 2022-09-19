package main

import (
	"go-gorm/controllers"
	"go-gorm/initializers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
