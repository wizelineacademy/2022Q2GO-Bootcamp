package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luischitala/2022Q2GO-Bootcamp/handlers"
	"github.com/luischitala/2022Q2GO-Bootcamp/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")

	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

	//To bypass the middleware

	//Create a new hub
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/characters", handlers.ListCharacterHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/readCsv", handlers.ReadCsvHandler(s)).Methods(http.MethodGet)
}
