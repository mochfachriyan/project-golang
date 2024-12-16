package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv mengatur pemuatan variabel dari file .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
