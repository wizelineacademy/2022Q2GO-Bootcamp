package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krmirandas/2022Q2GO-Bootcamp/api/controller"
)

func main() {
	fmt.Println("Server is lift")
	fmt.Println("Port: 8000")
	router := mux.NewRouter()

	buildCSVRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func buildCSVRoutes(router *mux.Router) {
	prefix := "/pokemon"
	router.HandleFunc(prefix+"/{id}", controller.GetItems).Methods("GET")
}
