package user

import "github.com/Jason-Bai/open-rest-go/models"

type UseCase interface {
	Create() (*models.User, error)
}
