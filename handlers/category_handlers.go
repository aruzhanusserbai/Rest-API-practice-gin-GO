package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't fetch data"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var newCategory models.Category

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&newCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
