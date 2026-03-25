package models

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Authors = make(map[int]Author)
var NextAuthorID = 1
