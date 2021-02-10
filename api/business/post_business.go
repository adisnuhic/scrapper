package business

import (
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
	"github.com/adisnuhic/scrapper_api/services"
)

// IPostBusiness intefrace
type IPostBusiness interface {
	GetAll() (*models.Posts, *apperror.AppError)
}

type postBusiness struct {
	Service services.IPostService
}

// NewPostBusiness -
func NewPostBusiness(svc services.IPostService) IPostBusiness {
	return &postBusiness{
		Service: svc,
	}
}

func (bl postBusiness) GetAll() (*models.Posts, *apperror.AppError) {
	return bl.Service.GetAll()
}
