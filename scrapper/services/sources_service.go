package services

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/repositories"
)

// SourceService -
type SourceService interface {
	GetAll() (*models.Sources, *apperror.AppError)
}

type sourceService struct {
	repo repositories.SourceRepository
}

// NewSourceService -
func NewSourceService(repo repositories.SourceRepository) SourceService {
	return &sourceService{
		repo: repo,
	}
}

// GetAll returns all sources
func (svc *sourceService) GetAll() (*models.Sources, *apperror.AppError) {
	return svc.repo.GetAll()
}
