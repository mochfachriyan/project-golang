package handlers

import (
	"net/http"
	"project/internal/services"
	"project/models"

	"github.com/gin-gonic/gin"
)

// UserHandler menangani request untuk user
type UserHandler struct {
	service *services.UserService
}

// NewUserHandler membuat handler baru untuk User
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// ListUsers menampilkan daftar user
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"users": users})
}

// CreateUserPage menampilkan form untuk membuat user baru
func (h *UserHandler) CreateUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "users/create.html", nil)
}

// CreateUser menangani pembuatan user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
