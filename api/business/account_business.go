package business

import (
	"errors"

	"github.com/adisnuhic/scrapper/ecode"
	"github.com/adisnuhic/scrapper/models"
	apperror "github.com/adisnuhic/scrapper/pkg"
	"github.com/adisnuhic/scrapper/services"
	"golang.org/x/crypto/bcrypt"
)

// IAccountBusiness interface
type IAccountBusiness interface {
	Register(firstName string, lastName string, email string, password string) *apperror.AppError
}

type accountBusiness struct {
	Service     services.IAccountService
	UserService services.IUserService
}

// NewAccountBusiness -
func NewAccountBusiness(svc services.IAccountService, userSvc services.IUserService) IAccountBusiness {
	return &accountBusiness{
		Service:     svc,
		UserService: userSvc,
	}
}

// Register user
func (bl accountBusiness) Register(firstName string, lastName string, email string, password string) *apperror.AppError {
	exists := bl.UserService.Exists(email)
	if exists {
		return apperror.New(ecode.ErrUserExistsCode, errors.New(ecode.ErrUserExistsMsg), ecode.ErrUserExistsMsg)
	}

	hashByte, errHashByte := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHashByte != nil {
		return nil
	}

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashByte),
	}

	return bl.Service.Register(user)
}
