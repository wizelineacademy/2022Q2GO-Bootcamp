package main

import (
	"log"
	"net/http"
	"wizeline/ghibli/adapters/controllers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/films", controllers.ServeFilms)

	log.Fatal(http.ListenAndServe(":8009", router))
}
