package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(100);unique" json:"code" validate:"required"`
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price" validate:"required"`
	Stock     int       `json:"stock" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
