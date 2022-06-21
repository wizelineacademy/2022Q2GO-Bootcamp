package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	router "main.com/router"
)

func main() {
	r := mux.NewRouter()
	router.RegisterProductRoutes(r)
	fmt.Println("The server is running")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8090",
		//timeouts for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
