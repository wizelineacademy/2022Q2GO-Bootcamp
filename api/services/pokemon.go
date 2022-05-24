package services

import (
	"fmt"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/api/hooks"
	"github.com/krmirandas/2022Q2GO-Bootcamp/api/model"
)

func RetrievePokemon() (output []model.Pokemon) {
	return hooks.Pokemonss
}

func RetrieveAll(quantityInput string) (output []model.Pokemon) {
	quantity, err := strconv.Atoi(quantityInput)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(quantity)

	output = make([]model.Pokemon, 0)

	return
}
