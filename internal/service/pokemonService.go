package service

import "github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"

type PokemonService interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	AddNewPokemons([]entity.Pokemon) ([]entity.Pokemon, error)
}

type PokemonRepository interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	AddNewPokemons([]entity.Pokemon) error
}

type pokemonService struct {
	pokemonRepository PokemonRepository
}

func NewPokemonService(pokemonRepository PokemonRepository) PokemonService {
	return &pokemonService{
		pokemonRepository: pokemonRepository,
	}
}

func (pKS *pokemonService) GetAllPokemons() ([]entity.Pokemon, error) {
	return nil, nil
}

func (pKS *pokemonService) GetPokemonById(id int) (*entity.Pokemon, error) {
	return nil, nil
}

func (pKS *pokemonService) AddNewPokemons([]entity.Pokemon) ([]entity.Pokemon, error) {
	return nil, nil
}
