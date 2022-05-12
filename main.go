package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type data struct {
	Index int64
	Item  string
}

func readFile(fileName string) []data {
	items := make([]data, 0)

	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println()
	}

	for _, line := range csvLines {
		index, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		item := data{
			Index: index,
			Item:  line[1],
		}
		items = append(items, item)
	}

	return items
}

func main() {

	// Call read csv file
	readLines := readFile("csvfile.csv")
	fmt.Println(readLines)
}
