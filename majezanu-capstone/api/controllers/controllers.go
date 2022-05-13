package controllers

import (
	"encoding/json"
	"log"
	"majezanu/capstone/api/services"
	"net/http"
	"strings"
)

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - GetPokemons - ", r.URL.Path)
	quantity := strings.TrimPrefix(r.URL.Path, "/quantity/")
	data, err := services.RetrievePokemon(quantity)
	if err.Code != 0 {
		log.Println("ERROR - GetPokemons - ", err.Description)
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - GetPokemonById - ", r.URL.Path)
	id := strings.TrimPrefix(r.URL.Path, "/id/")
	data, err := services.RetrievePokemonById(id)
	if err.Code != 0 {
		log.Println("ERROR - GetPokemonById - ", err.Description)
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("INFO - GetPokemonByName - ", r.URL.Path)
	name := strings.TrimPrefix(r.URL.Path, "/name/")
	data, err := services.RetrievePokemonByName(name)
	if err.Code != 0 {
		log.Println("ERROR - GetPokemonById - ", err.Description)
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
