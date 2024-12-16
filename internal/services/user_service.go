package services

import (
	"project/internal/repositories"
	"project/models"
)

// UserService mengelola logika bisnis terkait user
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUsers retrieves all users via the repository
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// AddUser creates a new user via the repository
func (s *UserService) AddUser(user *models.User) error {
	return s.repo.Create(user)
}
