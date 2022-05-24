package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krmirandas/2022Q2GO-Bootcamp/api/services"
)

func GetItems(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get infos about beers")
	initHeaders(writer)
	json.NewEncoder(writer).Encode(services.RetrievePokemon())
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
