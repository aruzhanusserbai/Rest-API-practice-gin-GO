package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookList []models.Book
	for _, book := range models.Books {
		bookList = append(bookList, book)
	}
	json.NewEncoder(w).Encode(bookList)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	book.ID = models.NextBookID
	models.NextBookID++
	models.Books[book.ID] = book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, exists := models.Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	_, exists := models.Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook.ID = id
	models.Books[id] = updatedBook
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	_, exists := models.Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(models.Books, id)
	w.WriteHeader(http.StatusNoContent)
}
