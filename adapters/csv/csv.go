package csv

import (
	"encoding/csv"
	"io"
)

func newMap[K comparable, V any](keys []K, vals []V ) (record map[K]V) {
	record = make(map[K]V, len(keys))
	for i := 0; i < len(keys); i++ {
		record[keys[i]] = vals[i]
	}
	return
}

type Structurer[T any] interface {
	ToStruct(map[string]string) (T, error)
}

type CsvRepository[T any] struct {
	reader *csv.Reader
	structurer Structurer[T]
}

func NewCsvRepository[T any](
	reader io.Reader,
	structurer Structurer[T]) CsvRepository[T] {

	return CsvRepository[T]{
		reader: csv.NewReader(reader),
		structurer: structurer,
	}
}

func (repo CsvRepository[T]) ListAll() ([]T, error) {
	elems := make([]T, 0)

	keys, err := repo.reader.Read()
	if err != nil {
		return elems, err
	}

	for {
		line, err := repo.reader.Read()
		if err == io.EOF {
			break
		}
		if err !=  nil {
			return elems, err
		}

		elem, err := repo.structurer.ToStruct(newMap(keys, line))
		if err != nil {
			return elems, err
		}

		elems = append(elems, elem)
	}
	return elems, nil
}

