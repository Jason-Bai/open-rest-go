package user

import "github.com/Jason-Bai/open-rest-go/models"

type Repository interface {
	Create() (*models.User, error)
}
