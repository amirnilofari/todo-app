package database

import (
	"github.com/amirnilofari/todo-app/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
