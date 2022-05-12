package main

import (
	"fmt"
)

func readFile(fileName string) []string {
	lines := make([]string, 3)
	fmt.Println("TODO Open and read file, then return slice of lines")
	return lines
}

func main() {

	// Call read csv file
	readLines := readFile("csvfile.csv")
	fmt.Println(readLines)
}
