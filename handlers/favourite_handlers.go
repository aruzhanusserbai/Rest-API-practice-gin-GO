package handlers

import (
	"ginExample/config"
	"ginExample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFavourites(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userId := c.GetUint("user_id")

	if err := config.DB.First(&models.Book{}, bookId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	favourite := models.Favourite{
		BookID: uint(bookId),
		UserID: userId,
	}
	if err := config.DB.Create(&favourite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book to favourites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added to favourites!"})
}

func GetFavourites(c *gin.Context) {
	userId := c.GetUint("user_id")
	var bookIds []uint

	if err := config.DB.Model(&models.Favourite{}).
		Select("book_id").
		Where("user_id = ?", userId).
		Scan(&bookIds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't find books"})
		return
	}

	var books []models.Book
	if err := config.DB.Model(&models.Book{}).Where("id IN ?", bookIds).Scan(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't find books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func DeleteFavourite(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userId := c.GetUint("user_id")

	result := config.DB.Where("user_id = ? AND book_id = ?", userId, bookId).
		Delete(&models.Favourite{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete favourite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book removed from favourites!"})
}
