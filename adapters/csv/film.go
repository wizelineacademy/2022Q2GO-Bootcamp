package csv

import (
	"errors"
	"strconv"
	"wizeline/ghibli/entities"
)


type filmStructurer struct {}

func NewFilmStructurer() Structurer[entities.Film]{
	return filmStructurer{}
}

func (f filmStructurer) ToStruct(
	record map[string]string) (entities.Film, error) {
	id, err := strconv.Atoi(record["id"])
	if err != nil {
		return entities.Film{}, errors.New("Non integer id")
	}

	return entities.Film{
		Id: id,
		Title: record["title"],
		Director: record["director"],
		ReleaseDate: record["release_date"],
	}, nil;
}
