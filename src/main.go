package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

type Person struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDatabase() {
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
		panic("DB Connection fail!!")
	}
	defer db.Close()
	log.Println("DB Connection Successfully!!")

	db.AutoMigrate(&Person{})
	log.Println("DB Migration Complete!!")
}

func main() {
	r := gin.Default()
	initDatabase()

	r.Run()

}
