package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	e "wizeline/ghibli/entities"
)

func toFilm(record map[string]string) (e.Film, error) {
	id, err := strconv.Atoi(record["id"])
	if err != nil {
		return e.Film{}, errors.New("Non integer id")
	}

	return e.Film{
		ID: id,
		Title: record["title"],
		Director: record["director"],
		ReleaseDate: record["release_date"],
	}, nil;
}

func newMap[K comparable, V any](keys []K, vals []V ) (record map[K]V) {
	record = make(map[K]V, len(keys))
	for i := 0; i < len(keys); i++ {
		record[keys[i]] = vals[i]
	}
	return
}

func ReadCSVFile(path string) {
	f, _ := os.Open(path)
	c := csv.NewReader(f)
	keys, _ := c.Read()

	for {
		r, e := c.Read()
		if e == io.EOF {
			break
		}

		film, _ := toFilm(newMap(keys, r))
		fmt.Println(film)
	}
}
