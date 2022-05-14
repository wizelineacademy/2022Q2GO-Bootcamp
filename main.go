package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pokemons = map[int64]string{}

func UnmarshalData(rows [][]string) {
	for _, row := range rows {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		pokemons[id] = row[1]

	}

}

func ReadCsv(filename string) ([][]string, error) {

	//Open CSV
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Cannot open CSV file:", err)
		return [][]string{}, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()

	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	return rows, nil
}

func main() {

	res, _ := ReadCsv("data.csv")
	fmt.Println(res)
	UnmarshalData(res)
	fmt.Println(pokemons)

}
