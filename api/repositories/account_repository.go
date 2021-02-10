package repositories

import (
	"github.com/adisnuhic/scrapper_api/db"
	"github.com/adisnuhic/scrapper_api/ecode"
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
)

// IAccountRepository interface
type IAccountRepository interface {
	Register(user *models.User) (*models.User, *apperror.AppError)
}

type accountRepository struct {
	Store db.Store
}

// NewAccountRepository -
func NewAccountRepository(store db.Store) IAccountRepository {
	return &accountRepository{
		Store: store,
	}
}

// Register user
func (repo accountRepository) Register(user *models.User) (*models.User, *apperror.AppError) {

	if err := repo.Store.Create(&user).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateUserCode, err, ecode.ErrUnableToCreateUserMsg)
	}

	return user, nil
}
