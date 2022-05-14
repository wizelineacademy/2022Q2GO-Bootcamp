package controller

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"../util" 
	"../model" 
)

const mugQuantity = 50

func GetItems(writer http.ResponseWriter, request *http.Request) {
	log.Println("Accessing the endpoint for reading csv")
	initHeaders(writer)
	
	pokemons := make(map[string]model.Pokemon)

	url := "https://gist.githubusercontent.com/armgilles/194bcff35001e7eb53a2a8b441e8b2c6/raw/92200bc0a673d5ce2110aaad4544ed6c4010f687/pokemon.csv"
	data, err := util.ReadCSVFromUrl(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	id := mux.Vars(request)["id"]

	for idx, row := range data {
		if idx == 0 {
			continue
		}

		newpokemon := model.Pokemon{
			ID: row[0],
			Name: row[1],
			Type1: row[2],
			Type2: row[3],
			Total: row[4],
			HP: row[5],
			Attack: row[6],
			Defense: row[7],
			SpAtk: row[8],
			SpDef: row[9],
			Speed: row[10],
			Generation: row[11],
			Legendary: row[12],
		}

		pokemons[row[0]] = newpokemon
	}

	fmt.Println(pokemons[id])
	json.NewEncoder(writer).Encode(pokemons[id])
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}