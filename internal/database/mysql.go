package database

import (
	"gorm.io/gorm"
	"os"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
