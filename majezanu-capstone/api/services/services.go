package services

import (
	"fmt"
	"majezanu/capstone/api/dataclass"
	"majezanu/capstone/api/repository"
	"net/http"
	"strconv"
)

func RetrievePokemon(quantityInput string) (output []dataclass.PokemonDataclass, errorOutput dataclass.ErrorDataclass) {
	quantity, err := strconv.Atoi(quantityInput)

	if err != nil {
		quantity = -1
	}

	output = make([]dataclass.PokemonDataclass, 0)
	input, err := repository.FetchPokemonData()

	if err != nil {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusNotFound,
			Description: fmt.Sprint(err),
		}
		return
	}
	inputSize := len(input)
	if inputSize < quantity {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusUnprocessableEntity,
			Description: "Max quantity",
		}
		return
	}

	for index, item := range input {
		data := dataclass.PokemonDataclass{
			Id:   item.Id,
			Name: item.Name,
		}
		if quantity == index {
			break
		}
		output = append(output, data)
	}
	return
}

func RetrievePokemonById(idInput string) (output dataclass.PokemonDataclass, errorOutput dataclass.ErrorDataclass) {
	id, err := strconv.Atoi(idInput)

	if err != nil {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusUnprocessableEntity,
			Description: "Bad format :id should be a number",
		}
	}

	input, err := repository.FetchPokemonDataById(id)

	if err != nil {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusNotFound,
			Description: fmt.Sprint(err),
		}
		return
	}

	output = dataclass.PokemonDataclass{
		Id:   input.Id,
		Name: input.Name,
	}
	return
}

func RetrievePokemonByName(nameInput string) (output dataclass.PokemonDataclass, errorOutput dataclass.ErrorDataclass) {

	if nameInput == "" {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusUnprocessableEntity,
			Description: "Bad format :name should be a string",
		}
	}

	input, err := repository.FetchPokemonDataByName(nameInput)

	if err != nil {
		errorOutput = dataclass.ErrorDataclass{
			Code:        http.StatusNotFound,
			Description: fmt.Sprint(err),
		}
		return
	}

	output = dataclass.PokemonDataclass{
		Id:   input.Id,
		Name: input.Name,
	}
	return
}
