package viewmodels

import "github.com/adisnuhic/scrapper_api/models"

// Auth -
type Auth struct {
	User         *models.User `json:"user"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
}
