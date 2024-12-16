package config

import (
	"fmt"
	"log"
	"os"
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah referensi global untuk koneksi database
var DB *gorm.DB

// ConnectDatabase menghubungkan ke database MySQL dan membuat database jika belum ada
func ConnectDatabase() {
	// Format DSN tanpa nama database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Koneksi ke MySQL tanpa menyebutkan database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Membuat database jika belum ada
	dbName := os.Getenv("DB_NAME")
	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	err = database.Exec(createDBQuery).Error
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	// Koneksi kembali dengan database yang sudah dibuat
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
	)

	database, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to reconnect to the database:", err)
	}

	DB = database
	log.Println("Database connected successfully!")
}

// AutoMigrate melakukan migrasi otomatis untuk model yang ada
func AutoMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Product{}) // Menambahkan model lainnya jika perlu
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}
	log.Println("Database migrated successfully!")
}
