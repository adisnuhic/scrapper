package business

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/services"
)

// SourceBusiness interface
type SourceBusiness interface {
	GetAll() (*models.Sources, *apperror.AppError)
}

type sourceBusiness struct {
	service services.SourceService
}

// NewSourceBusiness new
func NewSourceBusiness(svc services.SourceService) SourceBusiness {
	return &sourceBusiness{
		service: svc,
	}
}

// GetAll returns all sources
func (bl *sourceBusiness) GetAll() (*models.Sources, *apperror.AppError) {
	return bl.service.GetAll()
}
