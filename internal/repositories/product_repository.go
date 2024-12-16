package repositories

import (
	"project/models"

	"gorm.io/gorm"
)

// ProductRepository mengelola interaksi dengan tabel produk
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository membuat instance baru dari ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll mengambil semua produk dari database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Create menyimpan produk baru ke dalam database
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}
