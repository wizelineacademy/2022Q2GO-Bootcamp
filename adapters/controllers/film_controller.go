package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wizeline/ghibli/entities"
	"wizeline/ghibli/services/repository"
)

type FilmController struct {
	repo repository.Repository[entities.Film]
}

func NewFilmController(
	repo repository.Repository[entities.Film]) FilmController {
	return FilmController{repo: repo}
}

func sendEncoded(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (f FilmController) getEncodedFilms() ([]byte, error) {
	films, err := f.repo.ListAll()
	if err != nil {
		return []byte{}, err
	}

	return json.MarshalIndent(films, "", "\t")
}

func (f FilmController) getEncodedFilm(id string) ([]byte, error) {
	idn, err := strconv.Atoi(id)
	if err != nil {
		return []byte{}, err
	}

	film, err := repository.GetById(f.repo, idn)
	if err != nil {
		return []byte{}, err
	}

	return json.MarshalIndent(film, "", "\t")
}

func (f FilmController) tryGetEncoded(r *http.Request) ([]byte, error) {
	id := r.URL.Query().Get("id")

	switch {
	case id != "":
		return f.getEncodedFilm(id)
	default:
		return f.getEncodedFilms()
	}
}

func (f FilmController) ServeFilms(w http.ResponseWriter, r *http.Request) {
	data, err := f.tryGetEncoded(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		sendEncoded(w, data)
	}
}
