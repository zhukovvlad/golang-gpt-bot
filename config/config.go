// Package config provides utility functions to manage environment variables.
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Attempt to load the .env file only when in development mode.
	// In production or other environments, environment variables should be set externally.
	if os.Getenv("APP_ENV") == "development" {
		if err := godotenv.Load(); err != nil {
			// Log the absence of a .env file. It's fine to not have one in non-development modes.
			log.Println("No .env file found, but that's okay in non-development mode")
		}
	}
}

// GetEnv retrieves the value of the environment variable named by the key.
// If the key does not exist, it returns the fallback value.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
