package main

import (
	"ginExample/config"
	"ginExample/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	config.InitDB()

	r := gin.Default()

	// Book routes
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Author routes
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.AddAuthor)

	// Category routes
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)

	r.Run(":8080")
}
