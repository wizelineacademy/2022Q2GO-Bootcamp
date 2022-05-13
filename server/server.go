package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/TanZng/toh-api/database"
	"github.com/TanZng/toh-api/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

// Server
// Needs to implement Config method that returns a Config struct
type Server interface {
	Config() *Config
}

// Broker
// It's in charge of managing the Servers
type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseURL is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("🚀 Starting server on port", b.Config().Port)

	repo, err := database.NewCSVRepository("./database/toh.csv")
	if err != nil {
		log.Fatal("Error NewCSVRepository:", err)
	}

	repository.SetRepository(repo)

	if err = http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("Error ListenAndServe:", err)
	}
}
