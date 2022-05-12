package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Struct to read all environment variables
type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

// Init function for the constructor struct
func NewConfig() (*Config, error) {
	c := &Config{}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	if os.Getenv("PORT") == "" {
		return nil, errors.New("Port is required")
	}
	if os.Getenv("JWT_SECRET") == "" {
		return nil, errors.New("Secret is required")
	}
	if os.Getenv("DATABASE_URL") == "" {
		return nil, errors.New("Database is required")
	}
	c.Port = os.Getenv("PORT")
	c.JWTSecret = os.Getenv("JWT_SECRET")
	c.DatabaseUrl = os.Getenv("DATABASE_URL")

	return c, err
}
