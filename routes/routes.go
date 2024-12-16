package routes

import (
	"project/config"
	"project/internal/handlers"
	"project/internal/repositories"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengatur semua rute dalam aplikasi
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Setup repository untuk User dan Product
	userRepo := repositories.NewUserRepository(config.DB)
	productRepo := repositories.NewProductRepository(config.DB)

	// Setup service untuk User dan Product
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// Setup handler untuk User dan Product
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	// Rute untuk User
	r.GET("/", userHandler.ListUsers)
	r.GET("/users/create", userHandler.CreateUserPage)
	r.POST("/users", userHandler.CreateUser)

	// Rute untuk Product
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/create", productHandler.CreateProductPage)
	r.POST("/products", productHandler.CreateProduct)

	return r
}
