package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct untuk menampung konfigurasi
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig untuk mengambil konfigurasi dari environment variables
func LoadConfig() *Config {
	// Memuat file .env jika ada
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "mydatabase"),
		ServerPort: getEnv("SERVER_PORT", "5000"),
	}

	return config
}

// Helper function untuk mendapatkan environment variable dengan default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
