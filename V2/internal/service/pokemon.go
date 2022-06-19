package service

import (
	"context"

	"github.com/krmirandas/2022Q2GO-Bootcamp/V2/internal/entity"
)

// PokemonService the contract of the pokemon service
type PokemonService interface {
	// CreatePokemon create new record Pokemon
	CreatePokemon(ctx context.Context, Pokemon *entity.Pokemon) error

	// FindPokemon gets filtered specific Pokemon
	FindPokemon(ctx context.Context, filter *entity.Pokemon) ([]entity.Pokemon, error)

	Count(ctx context.Context) (int, error)
}

type PokemonRepo interface {
	ReadPokemon() ([]entity.Pokemon, error)
	WritePokemon(pokemon *entity.Pokemon) error
	Count(ctx context.Context) (int, error)
}

type pokemonService struct {
	repo PokemonRepo
}

func NewPokemonService(repo PokemonRepo) PokemonService {
	return &pokemonService{repo: repo}
}

func (ps *pokemonService) CreatePokemon(ctx context.Context, Pokemon *entity.Pokemon) error {
	return nil
}

func (ps *pokemonService) FindPokemon(ctx context.Context, filter *entity.Pokemon) ([]entity.Pokemon, error) {
	pokemons, err := ps.repo.ReadPokemon(ctx)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (ps *pokemonService) Count(ctx context.Context) (int, error) {
	return ps.repo.Count(ctx)
}
