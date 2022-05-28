package repository

import "github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"

type PokemonRepository interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	AddNewPokemons([]entity.Pokemon) error
}

type pokemonRepository struct {
	data []entity.Pokemon
}

func NewPokemonRepository(fileName string) PokemonRepository {
	return &pokemonRepository{}
}

func (pR *pokemonRepository) GetAllPokemons() ([]entity.Pokemon, error) {
	return nil, nil
}

func (pR *pokemonRepository) GetPokemonById(id int) (*entity.Pokemon, error) {
	return nil, nil
}

func (pR *pokemonRepository) AddNewPokemons([]entity.Pokemon) error {
	return nil
}
