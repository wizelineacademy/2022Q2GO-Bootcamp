package service

import (
	"testing"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/assert"
)

type mockRepository1 struct {
	repo PokemonInfoRepo
}

func Test_service_Info(t *testing.T) {
	s := NewPokemonInfoService(mockRepository1{})

	err := s.CreatePokemon(entity.PokemonInfo{})
	assert.NotNil(t, err)

	algo := s.CreatePokemonConcu(1)
	assert.NotNil(t, algo)
}

func (m mockRepository1) WritePokemon(pokemon entity.PokemonInfo) error {
	return errCRUD
}

func (m mockRepository1) ConcuRSwWP(itemsPerWorker int) []entity.PokemonInfo {
	var pokemons []entity.PokemonInfo
	pokemonMock := entity.PokemonInfo{
		ID:                     1,
		Name:                   "arbok",
		BaseExperience:         157,
		Height:                 35,
		IsDefault:              true,
		Order:                  33,
		Weight:                 650,
		LocationAreaEncounters: "https://pokeapi.co/api/v2/pokemon/24/encounters",
	}
	pokemons = append(pokemons, pokemonMock)
	return pokemons
}
