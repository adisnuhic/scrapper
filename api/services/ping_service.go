package services

import "github.com/adisnuhic/scrapper/repositories"

// IPingService interface
type IPingService interface {
	Ping() string
}

type pingService struct {
	Repository repositories.IPingRepository
}

// NewPingService -
func NewPingService(repo repositories.IPingRepository) IPingService {
	return &pingService{
		Repository: repo,
	}
}

// Ping returns string "pong"
func (svc pingService) Ping() string {
	return svc.Repository.Ping()
}
