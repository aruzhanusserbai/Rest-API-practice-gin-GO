package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuthors retrieves all authors from the database.
func GetAuthors(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name FROM authors")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning author"})
			return
		}
		authors = append(authors, author)
	}

	c.JSON(http.StatusOK, authors)
}

// AddAuthor adds a new author to the database.
func AddAuthor(c *gin.Context) {
	var newAuthor models.Author
	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newAuthor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author name is required"})
		return
	}

	query := "INSERT INTO authors (name) VALUES ($1) RETURNING id"
	err := config.DB.QueryRow(query, newAuthor.Name).Scan(&newAuthor.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add author"})
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
}
