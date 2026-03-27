package handlers

import (
	"bookstore/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	category := c.Query("category_id")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}

	var filtered []models.Book

	// filter
	for _, book := range models.Books {
		if category != "" {
			catID, _ := strconv.Atoi(category)
			if book.CategoryID != catID {
				continue
			}
		}
		filtered = append(filtered, book)
	}

	// pagination
	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	c.JSON(200, filtered[start:end])
}

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if book.Title == "" {
		c.JSON(400, gin.H{"error": "Title is required"})
		return
	}
	if book.Price <= 0 {
		c.JSON(400, gin.H{"error": "Price must be greater than 0"})
		return
	}

	book.ID = models.NextBookID
	models.NextBookID++
	models.Books[book.ID] = book

	c.JSON(200, book)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}

	book, exists := models.Books[id]
	if !exists {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}
	_, exists := models.Books[id]
	if !exists {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindBodyWithJSON(&updatedBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updatedBook.Title == "" {
		c.JSON(400, gin.H{"error": "Title is required"})
		return
	}
	if updatedBook.Price <= 0 {
		c.JSON(400, gin.H{"error": "Price must be greater than 0"})
		return
	}

	updatedBook.ID = id
	models.Books[id] = updatedBook
	c.JSON(200, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid book ID"})
		return
	}
	_, exists := models.Books[id]
	if !exists {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	delete(models.Books, id)
	c.Status(204)
}
