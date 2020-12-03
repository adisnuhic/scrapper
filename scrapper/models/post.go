package models

import "time"

// Post type
type Post struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at" sql:"-"`
}

// Posts collection
type Posts []Post
