package models

import "time"

// Source type
type Source struct {
	ID        uint64    `json:"id"`
	Source    string    `json:"source"`
	SourceURL string    `json:"source_url"`
	CreatedAt time.Time `json:"created_at"`
}

// Sources collection
type Sources []Source
