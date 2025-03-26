package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCategories retrieves all categories from the database.
func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning category"})
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, categories)
}

// AddCategory adds a new category to the database.
func AddCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newCategory.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	err := config.DB.QueryRow(query, newCategory.Name).Scan(&newCategory.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}
