package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("img-upload.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
