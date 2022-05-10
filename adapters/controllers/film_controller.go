package controllers

import (
	"encoding/json"
	"net/http"
	"wizeline/ghibli/adapters/repository"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	films, _ := repository.ListAllFilms()
	encFilms, _ := json.MarshalIndent(films, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encFilms)
}
