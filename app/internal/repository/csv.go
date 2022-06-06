package repository

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Reader interface {
	ReadCsvFile() (*os.File, error)
}

type ReaderConcurrently interface {
	ReadCsvFileConcurrently() ([]string, error)
}

type Writer interface {
	WriteCsvFile() (*csv.Writer, *os.File, error)
}

type Csv interface {
	Reader
	Writer
	ReaderConcurrently
}

//Chain struct to separate logic between the next layer
type csvr struct {
}

func NewCsvRepository() Csv {
	return &csvr{}
}

func (c *csvr) ReadCsvFile() (*os.File, error) {
	// Open the csv File
	csvFile, err := os.Open("files/characters.csv")
	if err != nil {
		fmt.Println(err)
	}
	// Return the file ready to iterate the rows
	return csvFile, err
}
func (c *csvr) ReadCsvFileConcurrently() ([]string, error) {
	// Open the csv File
	csvFile, err := ioutil.ReadFile("files/characters.csv")
	// Error handling
	if err != nil {
		fmt.Println(err)
	}
	// Split the content in strings
	content := strings.Split(string(csvFile), "\n")
	// Return the content of the csv file
	return content[:len(content)-1], nil
}

func (c *csvr) WriteCsvFile() (*csv.Writer, *os.File, error) {
	//Create Csv
	csvFile, err := os.Create("files/characterResult.csv")

	// Error handling
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Write into the csv
	csvwriter := csv.NewWriter(csvFile)

	return csvwriter, csvFile, err
}
