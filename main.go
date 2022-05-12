package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type data struct {
	Index string
	Item  string
}

func readFile(fileName string) [][]string {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println()
	}

	return csvLines
}

func main() {

	// Call read csv file
	readLines := readFile("csvfile.csv")
	fmt.Println(readLines)
}
