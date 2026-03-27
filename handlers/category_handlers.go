package handlers

import (
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categoryList []models.Category

	for _, category := range models.Categories {
		categoryList = append(categoryList, category)
	}

	c.JSON(200, categoryList)
}
func AddCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindBodyWithJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if category.Name == "" {
		c.JSON(400, gin.H{"error": "Category's name is required"})
		return
	}

	category.ID = models.NextCategoryID
	models.NextCategoryID++
	models.Categories[category.ID] = category

	c.JSON(200, category)
}
