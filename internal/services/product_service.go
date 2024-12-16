package services

import (
	"project/internal/repositories"
	"project/models"
)

// ProductService mengelola logika bisnis terkait produk
type ProductService struct {
	repo repositories.ProductRepository
}

// NewProductService membuat instance baru dari ProductService
func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetProducts mengambil semua produk dari database
func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

// AddProduct menambahkan produk baru
func (s *ProductService) AddProduct(product *models.Product) error {
	return s.repo.Create(product)
}
