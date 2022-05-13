package app

import (
	"fmt"
	"log"
	"majezanu/capstone/api/controllers"
	"net/http"
)

var port = ":1000"

func handleRequests() {
	http.HandleFunc("/quantity/", controllers.GetPokemons)
	http.HandleFunc("/id/", controllers.GetPokemonById)
	http.HandleFunc("/name/", controllers.GetPokemonByName)
	log.Fatal(http.ListenAndServe(port, nil))
}

func Run() {
	fmt.Println("Running server on PORT", port)
	handleRequests()
}
