package main

import (
	"os"
	"project/config"
	"project/routes"
)

func main() {
	// Memuat konfigurasi dari file .env
	config.LoadEnv()

	// Menghubungkan ke database dan membuat database jika belum ada
	config.ConnectDatabase()

	// Melakukan migrasi database secara otomatis
	config.AutoMigrate()

	// Menyiapkan router untuk aplikasi
	r := routes.SetupRouter()

	// Menjalankan aplikasi di port yang telah ditentukan
	r.Run(":" + os.Getenv("APP_PORT"))
}
