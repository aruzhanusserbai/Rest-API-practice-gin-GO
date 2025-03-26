package handlers

import (
	"database/sql"
	"ginExample/config"
	"ginExample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	// Parse page and limit query parameters with defaults
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Convert page and limit to integers
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	// Calculate offset for pagination
	offset := (pageNum - 1) * limitNum

	// Fetch books with LIMIT and OFFSET
	rows, err := config.DB.Query("SELECT id, title, author_id, category_id, price FROM books LIMIT $1 OFFSET $2", limitNum, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.CategoryID, &book.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book"})
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating books"})
		return
	}

	// Return paginated books
	c.JSON(http.StatusOK, gin.H{
		"page":  pageNum,
		"limit": limitNum,
		"total": len(books),
		"books": books,
	})
}

// CreateBook adds a new book to the database.
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Input validation
	if book.Title == "" || book.AuthorID <= 0 || book.CategoryID <= 0 || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	query := `INSERT INTO books (title, author_id, category_id, price) 
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := config.DB.QueryRow(query, book.Title, book.AuthorID, book.CategoryID, book.Price).Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBookByID retrieves a single book by its ID.
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	err := config.DB.QueryRow(
		"SELECT id, title, author_id, category_id, price FROM books WHERE id = $1", id,
	).Scan(&book.ID, &book.Title, &book.AuthorID, &book.CategoryID, &book.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book by its ID.
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Input validation
	if book.Title == "" || book.AuthorID <= 0 || book.CategoryID <= 0 || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	query := `UPDATE books 
              SET title = $1, author_id = $2, category_id = $3, price = $4 
              WHERE id = $5`
	result, err := config.DB.Exec(query, book.Title, book.AuthorID, book.CategoryID, book.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook deletes a book by its ID.
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
