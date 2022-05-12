package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const DB string = "./data.csv"

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

var Pokemons []Pokemon

func readDB(){
	log.Println("INFO - Reading DB")
	// TODO: read from csv file
	Pokemons = []Pokemon{
		{Id: "1", Name: "bulbasaur"},
		{Id: "2", Name: "ivasaur"},
	}
}

func index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - Home Page")
	readDB()
	json.NewEncoder(w).Encode(Pokemons)
}

func startServer(){
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main(){
	startServer()
}