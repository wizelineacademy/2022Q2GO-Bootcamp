package server

import (
	"fmt"
	"net/http"
	"os"
	"encoding/csv"
	"encoding/json"
)

func index(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "Method not supported")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	list := datas
	for _, line := range records {
		data := CsvDataLines{
		   Column1: line[0],
		   Column2: line[1],
		}
		list = append(list, data)
	 }
	json.NewEncoder(w).Encode(list)
}
