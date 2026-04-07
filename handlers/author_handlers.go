package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author

	if err := config.DB.Find(&authors).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't fetch data"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

func AddAuthor(c *gin.Context) {
	var newAuthor models.Author

	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&newAuthor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
}
