package services

import (
	"github.com/adisnuhic/scrapper/models"
	apperror "github.com/adisnuhic/scrapper/pkg"
	"github.com/adisnuhic/scrapper/repositories"
)

// IAccountService interface
type IAccountService interface {
	Ping() string
	Register(user *models.User) *apperror.AppError
}

type accountService struct {
	Repository repositories.IAccountRepository
}

// NewAccountService -
func NewAccountService(repo repositories.IAccountRepository) IAccountService {
	return &accountService{
		Repository: repo,
	}
}

// Ping returns string "pong"
func (svc accountService) Ping() string {
	return svc.Repository.Ping()
}

// Register user
func (svc accountService) Register(user *models.User) *apperror.AppError {
	return svc.Repository.Register(user)
}
