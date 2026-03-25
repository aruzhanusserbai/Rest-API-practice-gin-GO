package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Categories = make(map[int]Category)
var NextCategoryID = 1
