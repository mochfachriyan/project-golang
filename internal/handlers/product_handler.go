package handlers

import (
	"net/http"
	"project/internal/services"
	"project/models"

	"github.com/gin-gonic/gin"
)

// ProductHandler menangani request untuk produk
type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler membuat handler baru untuk Product
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// ListProducts menampilkan daftar produk
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "products/list.html", gin.H{"products": products})
}

// CreateProductPage menampilkan form untuk membuat produk baru
func (h *ProductHandler) CreateProductPage(c *gin.Context) {
	c.HTML(http.StatusOK, "products/create.html", nil)
}

// CreateProduct menangani pembuatan produk baru
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.AddProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/products")
}
