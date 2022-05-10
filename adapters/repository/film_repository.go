package repository

import (
	"wizeline/ghibli/adapters/csv"
	"wizeline/ghibli/entities"
)

func ListAllFilms() ([]entities.Film, error) {
	const csvFile = "data/ghibli.csv"
	return csv.ReadFilms(csvFile)
}
