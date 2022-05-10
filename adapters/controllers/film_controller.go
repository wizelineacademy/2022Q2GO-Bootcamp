package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wizeline/ghibli/adapters/repository"
)

func sendEncoded(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getEncodedFilms() ([]byte, error) {
	films, err := repository.ListAllFilms()
	if err != nil {
		return []byte{}, err
	}

	return json.MarshalIndent(films, "", "\t")
}

func getEncodedFilm(id string) ([]byte, error) {
	idn, err := strconv.Atoi(id)
	if err != nil {
		return []byte{}, err
	}

	film, err := repository.GetFilmById(idn)
	if err != nil {
		return []byte{}, err
	}

	return json.MarshalIndent(film, "", "\t")
}

func tryGetEncoded(r *http.Request) ([]byte, error) {
	id := r.URL.Query().Get("id")

	switch {
	case id != "":
		return getEncodedFilm(id)
	default:
		return getEncodedFilms()
	}
}

func ServeFilms(w http.ResponseWriter, r *http.Request) {
	data, err := tryGetEncoded(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		sendEncoded(w, data)
	}
}
