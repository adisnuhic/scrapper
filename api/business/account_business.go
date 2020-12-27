package business

import (
	"github.com/adisnuhic/scrapper/models"
	apperror "github.com/adisnuhic/scrapper/pkg"
	"github.com/adisnuhic/scrapper/services"
)

// IAccountBusiness interface
type IAccountBusiness interface {
	Ping() string
	Register(firstName string, lastName string, email string, password string) *apperror.AppError
}

type accountBusiness struct {
	Service services.IAccountService
}

// NewAccountBusiness -
func NewAccountBusiness(svc services.IAccountService) IAccountBusiness {
	return &accountBusiness{
		Service: svc,
	}
}

// Ping returns string "pong"
func (bl accountBusiness) Ping() string {
	return bl.Service.Ping()
}

// Register user
func (bl accountBusiness) Register(firstName string, lastName string, email string, password string) *apperror.AppError {
	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	return bl.Service.Register(user)
}
