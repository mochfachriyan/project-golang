package repositories

import (
	"project/models"

	"gorm.io/gorm"
)

// UserRepository handles database operations for the User model
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user to the database
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
