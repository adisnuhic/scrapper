package repositories

import (
	"github.com/adisnuhic/scrapper_api/db"
	"github.com/adisnuhic/scrapper_api/ecode"
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
)

// ITokenRepository interface
type ITokenRepository interface {
	CreateToken(token *models.Token) (*models.Token, *apperror.AppError)
	GetByToken(token string) (*models.Token, *apperror.AppError)
}

type tokenRepository struct {
	Store db.Store
}

// NewTokenRepository -
func NewTokenRepository(store db.Store) ITokenRepository {
	return &tokenRepository{
		Store: store,
	}
}

// CreateToken creates token
func (repo tokenRepository) CreateToken(token *models.Token) (*models.Token, *apperror.AppError) {
	if err := repo.Store.Create(token).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateTokenCode, err, ecode.ErrUnableToCreateTokenMsg)
	}

	return token, nil
}

// GetByToken returns token for provided token string
func (repo tokenRepository) GetByToken(token string) (*models.Token, *apperror.AppError) {
	model := new(models.Token)

	tx := repo.Store.Where("token = ?", token).Find(&model)

	if tx.Error != nil {
		return nil, apperror.New(ecode.ErrUnableToGetTokenCode, tx.Error, ecode.ErrUnableToGetTokenMsg)
	}

	return model, nil
}
