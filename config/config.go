package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found, using system environment variables.")
	}

	portStr := getEnv("DB_PORT", "5432")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB port: %w", err)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBSSLMode:  getEnv("DB_SSLMODE", ""),
		ServerPort: getEnv("SERVER_PORT", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
