package repository

import (
	"fmt"
	"os"
)

type Reader interface {
	ReadCsvFile() (*os.File, error)
}

// Csv - interface that concatenates all repository interfaces
type Csv interface {
	Reader
}

func ReadCsvFile() (*os.File, error) {
	csvFile, err := os.Open("files/characters.csv")
	if err != nil {
		fmt.Println(err)
	}

	return csvFile, err
}
