package repository

import "github.com/Jason-Bai/open-rest-go/models"

// Repository the repository of user
type Repository struct{}

// NewRepository create a new user repository
func NewRepository() *Repository {
	return &Repository{}
}

// Create insert a user record
func (r *Repository) Create() (*models.User, error) {
	return nil, nil
}
