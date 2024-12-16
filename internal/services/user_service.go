package services

import (
	"project/internal/repositories"
	"project/models"
)

// UserService mengelola logika bisnis terkait user
type UserService struct {
	repo repositories.UserRepository
}

// NewUserService membuat instance baru dari UserService
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUsers mengambil semua user dari database
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// AddUser menambahkan user baru
func (s *UserService) AddUser(user *models.User) error {
	return s.repo.Create(user)
}
