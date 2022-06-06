package utils

import "github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"

func CleanPokemonsResponse(pokemons []entity.Pokemon) []entity.Pokemon {
	newPokemons := []entity.Pokemon{}

	for _, pokemon := range pokemons {
		if pokemon.Id == 0 {
			break
		}

		newPokemons = append(newPokemons, pokemon)
	}

	return newPokemons
}
