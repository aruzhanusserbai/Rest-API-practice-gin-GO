package models

type Book struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	AuthorID   int     `json:"author_id"`
	CategoryID int     `json:"category_id"`
	Price      float64 `json:"price"`
}

var Books = make(map[int]Book)
var NextBookID = 1
