package main

import (
    "log"
    "net/http"
	"fmt"
	"os" 
	"encoding/csv"
	"encoding/json"
	"strconv"
)
//Type to store records from CSV file
type record struct {
	Key int64
	Value  string
}

var results []record

// GetHandler handles the index route
func GetGreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!, 2022Q2GO-Bootcamp"))
}

// GetCSVHandler handles the CSV route
func GetCSVHandler(w http.ResponseWriter, r *http.Request) {
	extractRecords("data.csv")
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}


// return data from file
func extractRecords(fileName string) {
	// open file
    f, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
	
    // convert records to array of structs
    results = createRecordList(data)
	fmt.Println(results)
  
}

func createRecordList(data [][]string) []record {
    var recordList []record
    for _, line := range data {
		var rec record
		
		for j, field := range line {
			if j == 0 {
				key, _ := strconv.ParseInt(field, 10, 64)
				rec.Key = key
			} else if j == 1 {
				rec.Value = field
			}
		}
		if rec.Key > 0 { // Omit items with invalid key
			recordList = append(recordList, rec)
		}
    }
    return recordList
}


func main() {
	//	extractRecords("data.csv") 
 
    // Create a mux for routing incoming requests
    m := http.NewServeMux()

    //  URLs will be handled by a function
    m.HandleFunc("/", GetGreetingHandler)
	m.HandleFunc("/getdata", GetCSVHandler)

	fmt.Println("Server Listening at port 8000")

    // Create a server listening on port 8000
    s := &http.Server{
        Addr:    ":8000",
        Handler: m,
    }
	
    // Continue to process new requests until an error occurs
    log.Fatal(s.ListenAndServe())
 
}