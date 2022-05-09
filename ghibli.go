package main

import (
	"wizeline/ghibli/adapters/csv"
)

func main() {
	csv.ReadCSVFile("data/ghibli.csv")
}
