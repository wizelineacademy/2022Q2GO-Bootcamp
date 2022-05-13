package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Fish struct {
	Id   int64
	Name string
}

type MyFish struct {
	Items []Fish
}

func csvReader(res http.ResponseWriter, req *http.Request) {
	// 1. Open the file
	recordFile, err := os.Open("./fish.csv")
	if err != nil {
		fmt.Println("CSV not valid: ", err)
		return
	}
	// 2. Initialize the reader
	reader := csv.NewReader(recordFile)
	// 3. Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error ocurred while reading the file: ", err)
		return
	}
	// 4. Show records in the response
	data := MyFish{}
	for _, row := range records {
		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		item := Fish{Id: id, Name: row[1]}
		data.Items = append(data.Items, item)
	}

	fmt.Println(data.Items)
	fmt.Fprintln(res, data.Items)

	//5. Close csv file
	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered closing the csv file ", err)
		return
	}
}

func setupRoutes() {
	http.HandleFunc("/readcsv", csvReader)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func main() {
	setupRoutes()
}
