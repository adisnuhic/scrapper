package repositories

import (
	"github.com/adisnuhic/scrapper/db"
)

// IPingRepository interface
type IPingRepository interface {
	Ping() string
}

type pingRepository struct {
	Store db.Store
}

// NewPingRepository -
func NewPingRepository() IPingRepository {
	return &pingRepository{}
}

// Ping returns string "pong"
func (repo pingRepository) Ping() string {
	return "pong..."
}
