package business

import (
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
)

// IPostBusiness intefrace
type IPostBusiness interface {
	GetAll() (*models.Posts, *apperror.AppError)
}

type postBusiness struct {
}

// NewPostBusiness -
func NewPostBusiness() IPostBusiness {
	return &postBusiness{}
}

func (bl postBusiness) GetAll() (*models.Posts, *apperror.AppError) {
	return nil, nil
}
