package configs

import (
	"os"
)

type Config struct {
	Port string
	Env  string
}

// LoadConfig loads config like (PORT, ENV, etc...) from .env or sets it to default val
func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // default environment
	}

	return &Config{
		Port: port,
		Env:  env,
	}, nil
}