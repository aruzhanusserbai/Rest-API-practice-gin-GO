package main

import (
	"bookstore/config"
	"bookstore/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	PORT := os.Getenv("PORT")
	config.ConnectDB()

	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.AddBook)
	r.GET("/books/:id", handlers.GetBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.AddAuthor)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)

	log.Println("Server running on port 8080")
	r.Run(fmt.Sprintf(":%v", PORT))
}
