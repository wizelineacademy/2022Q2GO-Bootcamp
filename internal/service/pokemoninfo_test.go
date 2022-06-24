package service

import (
	"testing"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestService_CreatePokemon(t *testing.T) {
	var testCases = []struct {
		name    string
		pokemon entity.PokemonInfo
		err     error
		// Repository
		repoErr error
	}{
		{
			"Should save a pokemon",
			entity.PokemonInfo{
				ID:                     1,
				Name:                   "arbok",
				BaseExperience:         157,
				Height:                 35,
				IsDefault:              true,
				Order:                  33,
				Weight:                 650,
				LocationAreaEncounters: "https://pokeapi.co/api/v2/pokemon/24/encounters",
			},
			nil,
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// testObj := new(.mockPokemonRepo)
			testObj := mocks.NewPokemonInfoRepository(t)

			// setup expectations
			testObj.On("WritePokemon", tc.pokemon).Return(tc.repoErr)

			// call the code we are testing
			s := NewPokemonInfoService(testObj)

			err := s.CreatePokemon(tc.pokemon)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestService_ConcuRSwWP(t *testing.T) {
	var testCases = []struct {
		name           string
		itemsPerWorker int
		response       []entity.PokemonInfo
		err            error
		// Repository
		repoRes []entity.PokemonInfo
		repoErr error
	}{
		{
			"Should save a pokemon",
			1,
			[]entity.PokemonInfo{
				{
					ID:                     1,
					Name:                   "arbok",
					BaseExperience:         157,
					Height:                 35,
					IsDefault:              true,
					Order:                  33,
					Weight:                 650,
					LocationAreaEncounters: "https://pokeapi.co/api/v2/pokemon/24/encounters",
				},
			},
			nil,
			[]entity.PokemonInfo{
				{
					ID:                     1,
					Name:                   "arbok",
					BaseExperience:         157,
					Height:                 35,
					IsDefault:              true,
					Order:                  33,
					Weight:                 650,
					LocationAreaEncounters: "https://pokeapi.co/api/v2/pokemon/24/encounters",
				},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// testObj := new(.mockPokemonRepo)
			testObj := mocks.NewPokemonInfoRepository(t)

			// setup expectations
			testObj.On("ConcuRSwWP", tc.itemsPerWorker).Return(tc.repoRes)

			// call the code we are testing
			s := NewPokemonInfoService(testObj)

			res := s.CreatePokemonConcu(tc.itemsPerWorker)

			assert.Equal(t, tc.response, res)
		})
	}
}
