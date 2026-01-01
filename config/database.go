package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	return db
}
