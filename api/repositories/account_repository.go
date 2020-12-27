package repositories

import (
	"errors"

	"github.com/adisnuhic/scrapper/db"
	"github.com/adisnuhic/scrapper/ecode"
	"github.com/adisnuhic/scrapper/models"
	apperror "github.com/adisnuhic/scrapper/pkg"
)

// IAccountRepository interface
type IAccountRepository interface {
	Ping() string
	Register(user *models.User) *apperror.AppError
}

type accountRepository struct {
	Store db.Store
}

// NewAccountRepository -
func NewAccountRepository() IAccountRepository {
	return &accountRepository{}
}

// Ping returns string "pong"
func (repo accountRepository) Ping() string {
	return "pong..."
}

// Register user
func (repo accountRepository) Register(user *models.User) *apperror.AppError {
	// DO REGISTER
	return apperror.New(ecode.ErrExampleCode, errors.New(ecode.ErrExampleMsg), ecode.ErrExampleMsg)
}
