package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	Port   string
	DBPath string
	Env    string
}

// Load loads configuration from environment variables
func Load() *Config {
	// Load .env file if it exists (ignore error if file doesn't exist)
	_ = godotenv.Load()

	config := &Config{
		Port:   getEnv("PORT", "8080"),
		DBPath: getEnv("DB_PATH", "./expenses.db"),
		Env:    getEnv("ENV", "development"),
	}

	return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetConfig returns the application configuration
var GetConfig = func() *Config {
	cfg := Load()
	log.Printf("Configuration loaded: Port=%s, DBPath=%s, Env=%s", cfg.Port, cfg.DBPath, cfg.Env)
	return cfg
}
