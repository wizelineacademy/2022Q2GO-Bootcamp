package repository

import (
	"encoding/csv"
	"errors"
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

type Util interface {
	VerifyIfFileExists(name string) (bool, error)
}

type Csv interface {
	Reader
	Writer
	ReaderConcurrently
	Util
}

//Chain struct to separate logic between the next layer
type csvr struct {
}

func NewCsvRepository() (Csv, error) {
	return &csvr{}, nil
}

// Functions that verify if a csv file, used before reading and writting
func (c *csvr) VerifyIfFileExists(name string) (bool, error) {
	//Create Csv
	// _, err := os.Stat("files/characterResults.csv")
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	// Handle if the file does not exist
	if errors.Is(err, os.ErrNotExist) {
		log.Println("File does not exist")
		return false, nil
	}
	log.Println("Error while evaluating")
	return false, err
}

func (c *csvr) ReadCsvFile() (*os.File, error) {
	file := "files/characters.csv"
	exists, err := c.VerifyIfFileExists(file)
	// Handle if the file does not exist
	if !exists {
		log.Println(err)
		return nil, err
	}
	// Open the csv File
	csvFile, err := os.Open(file)
	// Handle if can not open the file
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Return the file ready to iterate the rows
	return csvFile, nil
}

// Function that reads a csv file and return its content as a string array
func (c *csvr) ReadCsvFileConcurrently() ([]string, error) {
	file := "files/characters.csv"
	exists, err := c.VerifyIfFileExists(file)
	// Handle if the file does not exist
	if !exists {
		log.Println("File does not exist")
		return nil, err
	}
	// Open the csv File
	csvFile, err := ioutil.ReadFile(file)
	// Handle if can not read the selected file
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Split the content in strings
	content := strings.Split(string(csvFile), "\n")

	// Return the content of the csv file
	return content[:len(content)-1], nil
}

// Function that verify if a csv file, used before reading and writting

func (c *csvr) WriteCsvFile() (*csv.Writer, *os.File, error) {
	//Create Csv
	csvFile, err := os.Create("files/characterResult.csv")

	// Handle if can not create the file
	if err != nil {
		log.Printf("Failed creating file: %s", err)
		return nil, nil, err
	}

	// Write into the csv
	csvwriter := csv.NewWriter(csvFile)

	// Handle if can not create the csv writer
	if err != nil {
		log.Printf("Failed creating writer: %s", err)
		return nil, nil, err
	}

	return csvwriter, csvFile, err
}
