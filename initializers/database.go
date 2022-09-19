package initializers

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
