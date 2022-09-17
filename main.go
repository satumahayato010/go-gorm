package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DBConn *gorm.DB

func initDB() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "person.db")
	if err != nil {
		panic("DB connection fail!!")
	}
	log.Println("DB connection successfully!!")
}

func main() {
	initDB()
	defer DBConn.Close()
}
