package repository

import (
	"wizeline/ghibli/adapters/csv"
	"wizeline/ghibli/entities"
	"errors"
)

func ListAllFilms() ([]entities.Film, error) {
	const csvFile = "data/ghibli.csv"
	return csv.ReadFilms(csvFile)
}

func GetFilmById(id int) (entities.Film, error) {
	films, err := ListAllFilms()
	if err != nil {
		return entities.Film{}, err
	}

	for _, f := range films {
		if f.ID == id {
			return f, nil
		}
	}
	return entities.Film{}, errors.New("Film by ID not found")
}
