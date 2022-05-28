package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Reader interface {
	ReadCsvFile() (*os.File, error)
}

type Writer interface {
	WriteCsvFile() (*csv.Writer, *os.File, error)
}

type Csv interface {
	Reader
	Writer
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

func (c *csvr) WriteCsvFile() (*csv.Writer, *os.File, error) {
	//Create Csv
	csvFile, err := os.Create("files/characterResult.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Write into the csv
	csvwriter := csv.NewWriter(csvFile)

	return csvwriter, csvFile, err
}
