package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

var DBConn *gorm.DB

func initDB() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("DB connection fail!!")
	}
	log.Println("DB connection successfully!!")

	DBConn.AutoMigrate(&Person{})
	log.Println("DB Migrated completed!!")
}

func main() {
	initDB()
	defer DBConn.Close()
}
