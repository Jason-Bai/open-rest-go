package usecase

import (
	"github.com/Jason-Bai/open-rest-go/apis/user"
	"github.com/Jason-Bai/open-rest-go/models"
)

// UserUseCase a service of user
type UserUseCase struct {
	repository user.Repository
}

// NewUseCase create a user service
func NewUseCase(repository user.Repository) *UserUseCase {
	return &UserUseCase{
		repository,
	}
}

// Create create a user
func (us *UserUseCase) Create() (*models.User, error) {
	return nil, nil
}
