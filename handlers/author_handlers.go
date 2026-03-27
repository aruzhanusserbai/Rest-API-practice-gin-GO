package handlers

import (
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authorList []models.Author
	for _, author := range models.Authors {
		authorList = append(authorList, author)
	}
	c.JSON(200, authorList)
}
func AddAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindBodyWithJSON(&author); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if author.Name == "" {
		c.JSON(400, gin.H{"error": "Author's name is required"})
		return
	}

	author.ID = models.NextAuthorID
	models.NextAuthorID++
	models.Authors[author.ID] = author

	c.JSON(201, author)
}
