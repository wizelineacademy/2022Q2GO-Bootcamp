package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Add the database path
const DB string = "./data.csv"

// Create pokemon struct
type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// Global array for pokemons
var Pokemons []Pokemon

// readDB reads a CSV file and return the matrix of entries
func readDB() [][]string{
	log.Println("INFO - Reading DB")

	var file, err = os.OpenFile(DB, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, _ := reader.ReadAll();
	return data
}

// mapDBtoStructArr takes a matrix of strings (the result from DB)
// and map it to the Struct Array
func mapDBtoStructArr(data [][]string){
	for _, d := range data {
		Pokemons = append(Pokemons, Pokemon{
			Id: d[0],
			Name: d[1],
		})
	}
}

// index return the index page content
func index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - Home Page")
	data := readDB()
	mapDBtoStructArr(data)
	json.NewEncoder(w).Encode(Pokemons)
}

// startServer starts the server
func startServer(){
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main(){
	startServer()
}