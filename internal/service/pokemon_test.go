package service

import (
	"errors"
	"strconv"
	"testing"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")

type mockRepository struct {
	repo PokemonRepo
}

func Test_service(t *testing.T) {
	s := NewPokemonService(mockRepository{})
	pokemonMock := entity.Pokemon{
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
	}
	id := pokemonMock.ID

	// get
	_, err := s.FindPokemonById("none")
	assert.NotNil(t, err)
	pokemon, err := s.FindPokemonById(id)
	assert.Nil(t, err)
	assert.Equal(t, pokemonMock.Name, pokemon.Name)
	assert.Equal(t, id, pokemon.ID)

	// fiind
	_, err = s.FindPokemon()
	assert.NotNil(t, err)

	// count
	count, _ := s.Count()
	assert.Equal(t, 1, count)
}

func (m mockRepository) Count() (int, error) {
	var pokemons []entity.Pokemon
	pokemonMock := entity.Pokemon{
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
	}
	pokemons = append(pokemons, pokemonMock)
	return len(pokemons), nil
}

func (m mockRepository) ReadOnePokemon(id string) (entity.Pokemon, error) {
	var pokemon entity.Pokemon
	_, err := strconv.Atoi(id)
	if err != nil {
		return pokemon, err
	}
	pokemonMock := entity.Pokemon{
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
	}
	return pokemonMock, err
}

func (m mockRepository) ReadPokemon() ([]entity.Pokemon, error) {
	var pokemons []entity.Pokemon
	pokemonMock := entity.Pokemon{
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
	}
	pokemons = append(pokemons, pokemonMock)
	return pokemons, errCRUD
}
