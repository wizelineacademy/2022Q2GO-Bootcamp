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
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		return nil, errors.New("Port is required")
	}
	c.JWTSecret = os.Getenv("JWT_SECRET")
	if c.JWTSecret == "" {
		return nil, errors.New("Secret is required")
	}

	c.DatabaseUrl = os.Getenv("DATABASE_URL")

	if c.DatabaseUrl == "" {
		return nil, errors.New("Database is required")
	}
	c.Port = os.Getenv("PORT")
	c.JWTSecret = os.Getenv("JWT_SECRET")
	c.DatabaseUrl = os.Getenv("DATABASE_URL")

	return c, err
}
