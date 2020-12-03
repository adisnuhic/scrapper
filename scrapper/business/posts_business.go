package business

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/services"
)

// PostBusiness interface
type PostBusiness interface {
	GetByID(id uint64) (*models.Post, *apperror.AppError)
}

type postBusiness struct {
	service services.PostService
}

// NewPostBusiness new
func NewPostBusiness(svc services.PostService) PostBusiness {
	return &postBusiness{
		service: svc,
	}
}

func (bl *postBusiness) GetByID(id uint64) (*models.Post, *apperror.AppError) {
	return bl.service.GetByID(id)
}
