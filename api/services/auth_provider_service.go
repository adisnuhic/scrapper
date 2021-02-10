package services

import (
	"github.com/adisnuhic/scrapper_api/models"
	apperror "github.com/adisnuhic/scrapper_api/pkg"
	"github.com/adisnuhic/scrapper_api/repositories"
)

// IAuthProviderService interface
type IAuthProviderService interface {
	GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError)
	Delete(auth *models.AuthProvider) *apperror.AppError
	Save(auth *models.AuthProvider) *apperror.AppError
}

type authProviderService struct {
	Repository repositories.IAuthProviderRepository
}

// NewAuthProviderService -
func NewAuthProviderService(repo repositories.IAuthProviderRepository) IAuthProviderService {
	return &authProviderService{
		Repository: repo,
	}
}

// GetByUserID return auth provider for provided ID
func (svc authProviderService) GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError) {
	return svc.Repository.GetByUserID(id)
}

// Save auth data
func (svc authProviderService) Save(auth *models.AuthProvider) *apperror.AppError {
	return svc.Repository.Save(auth)
}

// Deletes auth data
func (svc authProviderService) Delete(auth *models.AuthProvider) *apperror.AppError {
	return svc.Repository.Delete(auth)
}
