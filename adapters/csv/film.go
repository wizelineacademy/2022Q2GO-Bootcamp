package csv

import (
	"encoding/csv"
	"io"
	"os"
	e "wizeline/ghibli/entities"
)

func newMap[K comparable, V any](keys []K, vals []V ) (record map[K]V) {
	record = make(map[K]V, len(keys))
	for i := 0; i < len(keys); i++ {
		record[keys[i]] = vals[i]
	}
	return
}

func ReadFilms(path string) ([]e.Film, error) {
	films := make([]e.Film, 0)
	f, err := os.Open(path)
	if err != nil {
		return films, err
	}

	c := csv.NewReader(f)
	keys, err := c.Read()
	if err != nil {
		return films, err
	}

	for {
		r, err := c.Read()
		if err == io.EOF {
			break
		}
		if err !=  nil {
			return films, err
		}

		film, err := e.ToFilm(newMap(keys, r))
		if err != nil {
			return films, err
		}

		films = append(films, film)
	}
	return films, nil
}
