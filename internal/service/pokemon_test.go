package service

import (
	"errors"
	"testing"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")

func TestService_FindPokemonById(t *testing.T) {
	var testCases = []struct {
		name     string
		filter   string
		response entity.Pokemon
		err      error
		// Repository
		repoRes entity.Pokemon
		repoErr error
	}{
		{
			"Should return a pokemon",
			"8",
			entity.Pokemon{
				ID:         "700",
				Name:       "Sylveon",
				Type1:      "Fairy",
				Type2:      "",
				Total:      "525",
				HP:         "95",
				Attack:     "65",
				Defense:    "65",
				SpAtk:      "110",
				SpDef:      "130",
				Speed:      "60",
				Generation: "6",
				Legendary:  "False",
			},
			nil,
			entity.Pokemon{
				ID:         "700",
				Name:       "Sylveon",
				Type1:      "Fairy",
				Type2:      "",
				Total:      "525",
				HP:         "95",
				Attack:     "65",
				Defense:    "65",
				SpAtk:      "110",
				SpDef:      "130",
				Speed:      "60",
				Generation: "6",
				Legendary:  "False",
			},
			nil,
		},
		{
			"Should return a pokemon",
			"8",
			entity.Pokemon{},
			errors.New("error find"),
			entity.Pokemon{},
			errors.New("error find"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// testObj := new(.mockPokemonRepo)
			testObj := mocks.NewPokemonRepository(t)

			// setup expectations
			testObj.On("ReadOnePokemon", tc.filter).Return(tc.repoRes, tc.repoErr)

			// call the code we are testing
			s := NewPokemonService(testObj)

			res, err := s.FindPokemonById(tc.filter)

			assert.Equal(t, tc.response, res)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestService_FindPokemon(t *testing.T) {
	var testCases = []struct {
		name     string
		response []entity.Pokemon
		err      error
		// Repository
		repoRes []entity.Pokemon
		repoErr error
	}{
		{
			"Should return array of pokemons",
			[]entity.Pokemon{
				{
					ID:         "700",
					Name:       "Sylveon",
					Type1:      "Fairy",
					Type2:      "",
					Total:      "525",
					HP:         "95",
					Attack:     "65",
					Defense:    "65",
					SpAtk:      "110",
					SpDef:      "130",
					Speed:      "60",
					Generation: "6",
					Legendary:  "False",
				},
			},
			nil,
			[]entity.Pokemon{
				{
					ID:         "700",
					Name:       "Sylveon",
					Type1:      "Fairy",
					Type2:      "",
					Total:      "525",
					HP:         "95",
					Attack:     "65",
					Defense:    "65",
					SpAtk:      "110",
					SpDef:      "130",
					Speed:      "60",
					Generation: "6",
					Legendary:  "False",
				},
			},
			nil,
		},
		{
			"Should throw errors",
			[]entity.Pokemon(nil),
			errCRUD,
			[]entity.Pokemon(nil),
			errCRUD,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testObj := mocks.NewPokemonRepository(t)

			// setup expectations
			testObj.On("ReadPokemon").Return(tc.repoRes, tc.repoErr)

			// call the code we are testing
			s := NewPokemonService(testObj)

			res, err := s.FindPokemon()

			assert.Equal(t, tc.response, res)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestService_Count(t *testing.T) {
	var testCases = []struct {
		name     string
		response int
		err      error
		// Repository
		repoRes int
		repoErr error
	}{
		{
			"Should return a count",
			1,
			nil,
			1,
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testObj := mocks.NewPokemonRepository(t)

			// setup expectations
			testObj.On("Count").Return(tc.repoRes, tc.repoErr)

			// call the code we are testing
			s := NewPokemonService(testObj)

			res, err := s.Count()

			assert.Equal(t, tc.response, res)
			assert.Equal(t, tc.err, err)
		})
	}
}
