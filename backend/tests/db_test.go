package tests

import (
	"fmt"

	"github.com/user123/URL-shortener-Golang/backend/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	if err = db.AutoMigrate(&handlers.URL{}); err != nil {
		panic(fmt.Sprintf("Failed to migrate schema: %v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	return db, func() {
		sqlDB.Close()
		fmt.Println("Connection closed successfully")
	}
}
