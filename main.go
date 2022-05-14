package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	controller "./controller" 
)


func main(){
	fmt.Println("Server is lift")
	fmt.Println("Port: 8000")
	router := mux.NewRouter()

	buildCSVRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func buildCSVRoutes(router *mux.Router) {
	prefix := "/pokemon"
	router.HandleFunc(prefix + "/{id}", controller.GetItems).Methods("GET")
}