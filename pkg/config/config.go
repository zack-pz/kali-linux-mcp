package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Kali_url string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Kali_url: getEnv("KALI_URL", "local://"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
