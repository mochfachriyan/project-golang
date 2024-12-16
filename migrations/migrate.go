package migrations

import (
	"log"
	"project/config"
	"project/models"
)

func RunMigrations() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrated successfully!")
}
