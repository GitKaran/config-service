package database

import (
	"github.com/hellofreshdevtests/GitKaran-devops-test/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//DB is connected database object
var DB *gorm.DB

func Setup() {
	db, err := gorm.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	// Creates a table name Config from the struct that we have defined
	db.AutoMigrate([]models.Config{})
	DB = db
}

// To get the DB connection
func GetDB() *gorm.DB {
	return DB
}
