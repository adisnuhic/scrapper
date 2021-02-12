package services

import (
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
	"github.com/adisnuhic/scrapper/repositories"
)

// PostService asd
type PostService interface {
	GetByID(id uint64) (*models.Post, *apperror.AppError)
	CreateMany(posts *models.Posts) (*models.Posts, *apperror.AppError)
	GetAll() (*models.Posts, *apperror.AppError)
}

type postService struct {
	repo repositories.PostRepository
}

// NewPostService -
func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{
		repo: repo,
	}
}

// CreateMany creates many posts
func (svc *postService) CreateMany(posts *models.Posts) (*models.Posts, *apperror.AppError) {
	return svc.repo.CreateMany(posts)
}

// GetByID return post for provided ID
func (svc *postService) GetByID(id uint64) (*models.Post, *apperror.AppError) {
	return svc.repo.GetByID(id)
}

// GetAll returns all posts
func (svc *postService) GetAll() (*models.Posts, *apperror.AppError) {
	return svc.repo.GetAll()
}
