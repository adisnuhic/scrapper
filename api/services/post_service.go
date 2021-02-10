package services

import (
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
)

// IPostService -
type IPostService interface {
	GetAll() (*models.Posts, *apperror.AppError)
}

type postService struct{}

// NewPostService -
func NewPostService() IPostService {
	return &postService{}
}

func (postService) GetAll() (*models.Posts, *apperror.AppError) {
	return nil, nil
}
