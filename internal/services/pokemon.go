package services

import (
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
)

func RetrievePokemon() (output []model.Pokemon) {
	return hook.Pokemonss
}

func RetrievePokemonById(idInput string) (output model.Pokemon, errorOutput *errorhandler.Error) {
	_, err := strconv.Atoi(idInput)

	if err != nil {
		errorOutput = errorhandler.ErrNotValidItemId
	}

	for _, pokemon := range hook.Pokemonss {
		if pokemon.ID == idInput {
			output = pokemon
			break
		}
	}

	return
}
