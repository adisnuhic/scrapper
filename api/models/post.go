package models

// Post model
type Post struct{
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

// Posts collection
type Posts []Post