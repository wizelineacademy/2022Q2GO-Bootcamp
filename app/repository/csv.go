package repository

import (
	"fmt"
	"os"
)

type Reader interface {
	ReadCsvFile() (*os.File, error)
}

type Writer interface {
	ReadCsvFile() (*os.File, error)
}

type Csv interface {
	Reader
}

//Chain struct to separate logic between the next layer
type csvr struct {
}

func NewCsvRepository() Csv {
	return &csvr{}
}

func (c *csvr) ReadCsvFile() (*os.File, error) {
	csvFile, err := os.Open("files/characters.csv")
	if err != nil {
		fmt.Println(err)
	}

	return csvFile, err
}
