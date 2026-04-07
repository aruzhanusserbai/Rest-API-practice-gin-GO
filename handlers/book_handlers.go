package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var filteredBooks []models.Book

	category := c.Query("category_id")
	page, _ := strconv.Atoi(c.DefaultQuery("p", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("l", "5"))
	offset := (page - 1) * limit

	query := config.DB.Limit(limit).Offset(offset)

	if category != "" {
		query = query.Where("category_id=?", category)
	}

	query.Find(&filteredBooks)
	c.JSON(http.StatusOK, filteredBooks)
}

func AddBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if newBook.Title == "" {
		c.JSON(http.StatusBadRequest, "Title can not be empty!")
		return
	}
	if newBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater that 0"})
		return
	}

	if err := config.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("id")
	var book models.Book

	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	book.Title = updatedBook.Title
	book.Price = updatedBook.Price
	book.CategoryID = updatedBook.CategoryID
	book.AuthorID = updatedBook.AuthorID

	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.First(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
