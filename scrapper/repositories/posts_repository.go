package repositories

import (
	"errors"

	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/ecode"
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
)

// PostRepository interface
type PostRepository interface {
	GetByID(id uint64) (*models.Post, *apperror.AppError)
	CreateMany(posts *models.Posts) (*models.Posts, *apperror.AppError)
	GetAll() (*models.Posts, *apperror.AppError)
}

type postRepository struct {
	Store db.Store
}

// NewPostRepository create new post repo
func NewPostRepository(s db.Store) PostRepository {
	return &postRepository{
		Store: s,
	}
}

// GetByID get post for provided id
func (repo *postRepository) GetByID(id uint64) (*models.Post, *apperror.AppError) {
	post := new(models.Post)

	tx := repo.Store.Where("id = ?", id).Find(post)

	if tx.Error != nil {
		return nil, apperror.New(ecode.ErrExampleTestCode, tx.Error, ecode.ErrExampleTestMsg)
	}

	return post, nil
}

// CreateMany creates many posts
func (repo *postRepository) CreateMany(posts *models.Posts) (*models.Posts, *apperror.AppError) {
	tx := repo.Store.Begin()
	myPosts := make(models.Posts, 0)

	for _, p := range *posts {
		if err := tx.Create(&p).Error; err != nil {
			// tx.Rollback()
			// return nil, apperror.New(ecode.ErrUnableToCreatePostCode, err, ecode.ErrUnableToCreatePostMsg)
		} else {
			myPosts = append(myPosts, p)
		}

	}

	tx.Commit()

	return &myPosts, nil
}

// GetAll returns all posts
func (repo *postRepository) GetAll() (*models.Posts, *apperror.AppError) {
	model := make(models.Posts, 0)

	tx := repo.Store.Find(&model)

	if tx.Error != nil && !tx.RecordNotFound() {
		return nil, apperror.New(ecode.ErrUnableToFetchPostCode, errors.New(ecode.ErrUnableToFetchPostMsg), ecode.ErrUnableToFetchPostMsg)
	}

	return &model, nil
}
