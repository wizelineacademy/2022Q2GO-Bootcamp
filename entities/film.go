package entities

import (
	"errors"
	"strconv"
)

type Film struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Director string `json:"director"`
	ReleaseDate string `json:"release_date"`
}

func ToFilm(record map[string]string) (Film, error) {
	id, err := strconv.Atoi(record["id"])
	if err != nil {
		return Film{}, errors.New("Non integer id")
	}

	return Film{
		ID: id,
		Title: record["title"],
		Director: record["director"],
		ReleaseDate: record["release_date"],
	}, nil;
}
