package business

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/services"
)

// PostBusiness interface
type PostBusiness interface {
	GetByID(id uint64) (*models.Post, *apperror.AppError)
	GetAll() (*models.Posts, *apperror.AppError)
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

// GetByID returns post for provided ID
func (bl *postBusiness) GetByID(id uint64) (*models.Post, *apperror.AppError) {
	return bl.service.GetByID(id)
}

// GetAll reuturns all posts
func (bl *postBusiness) GetAll() (*models.Posts, *apperror.AppError) {
	return bl.service.GetAll()
}
