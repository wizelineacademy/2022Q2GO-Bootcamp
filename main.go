package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/TanZng/toh-api/handlers"
	"github.com/TanZng/toh-api/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("SERVER_PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URL := os.Getenv("DB_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DB_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/character/{id}", handlers.GetCharacterByIdHandler(s)).Methods(http.MethodGet)
}
