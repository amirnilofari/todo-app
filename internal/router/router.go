package router

import (
	"github.com/amirnilofari/todo-app/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/todos", handlers.GetTodos)

	return r
}
