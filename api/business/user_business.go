package business

import (
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
	"github.com/adisnuhic/scrapper_api/services"
)

// IUserBusiness interface
type IUserBusiness interface {
	GetByID(id uint64) (*models.User, *apperror.AppError)
	GetByEmail(email string) (*models.User, *apperror.AppError)
}

type userBusiness struct {
	Service services.IUserService
}

// NewUserBusiness -
func NewUserBusiness(svc services.IUserService) IUserBusiness {
	return &userBusiness{
		Service: svc,
	}
}

// GetByID returns user for provided ID
func (bl userBusiness) GetByID(id uint64) (*models.User, *apperror.AppError) {
	return bl.Service.GetByID(id)
}

// GetByEmail returns user for provided email
func (bl userBusiness) GetByEmail(email string) (*models.User, *apperror.AppError) {
	return bl.Service.GetByEmail(email)
}
