package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Config holds the configuration values for the application
type Config struct {
	MongoURI      string
	MongoDB       string
	ServerAddress string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load environment variables from .enc file if it exists
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Retrieve configuration values from environment variables
	config := &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:       getEnv("MONGO_DB", "urlshortener"),
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
	}

	return config, nil
}

// getEnv retrieves the value of the environment variable names by the key.
// If the variable is empty, it returns the defaultValue.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
