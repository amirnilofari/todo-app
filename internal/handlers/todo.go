package handlers

import (
	"github.com/amirnilofari/todo-app/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := c.MustGet("db").(*gorm.DB).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
