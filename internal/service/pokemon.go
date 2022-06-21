package service

import (
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

// PokemonService the contract of the pokemon service
type PokemonService interface {
	// FindPokemon gets filtered specific Pokemon
	FindPokemon() ([]entity.Pokemon, error)

	FindPokemonById(filter string) (entity.Pokemon, error)

	Count() (int, error)
}

type PokemonRepo interface {
	ReadPokemon() ([]entity.Pokemon, error)
	ReadOnePokemon(id string) (entity.Pokemon, error)
	Count() (int, error)
}

type pokemonService struct {
	repo PokemonRepo
}

func NewPokemonService(repo PokemonRepo) PokemonService {
	return &pokemonService{repo: repo}
}

func (ps *pokemonService) FindPokemon() ([]entity.Pokemon, error) {
	pokemons, err := ps.repo.ReadPokemon()
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (ps *pokemonService) FindPokemonById(filter string) (entity.Pokemon, error) {
	pokemon, err := ps.repo.ReadOnePokemon(filter)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func (ps *pokemonService) Count() (int, error) {
	return ps.repo.Count()
}
