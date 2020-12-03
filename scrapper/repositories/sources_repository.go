package repositories

import (
	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/ecode"
	"github.com/adisnuhic/scrapper/models"
	"github.com/adisnuhic/scrapper/pkg/apperror"
)

// SourceRepository interface
type SourceRepository interface {
	GetAll() (*models.Sources, *apperror.AppError)
}

type sourceRepository struct {
	Store db.Store
}

// NewSourceRepository create new post repo
func NewSourceRepository(s db.Store) SourceRepository {
	return &sourceRepository{
		Store: s,
	}
}

// GetAll returns all sources
func (repo *sourceRepository) GetAll() (*models.Sources, *apperror.AppError) {
	sources := new(models.Sources)

	tx := repo.Store.Find(sources)

	if tx.Error != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchSourcesCode, tx.Error, ecode.ErrUnableToFetchSourcesMsg)
	}

	return sources, nil
}
