package service

import (
	"sync"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

var mu sync.Mutex

// PokemonService the contract of the pokemon service
type PokemonInfoService interface {
	// CreatePokemon create new record Pokemon
	CreatePokemon(pokemonInfo entity.PokemonInfo) error
	CreatePokemonConcu(itemsPerWorker int) []entity.PokemonInfo
}

type PokemonInfoRepo interface {
	WritePokemon(pokemonInfo entity.PokemonInfo) error
	ConcuRSwWP(itemsPerWorker int) []entity.PokemonInfo
}

type pokemonInfoService struct {
	repo PokemonInfoRepo
}

func NewPokemonInfoService(repo PokemonInfoRepo) PokemonInfoService {
	return &pokemonInfoService{repo: repo}
}

func (ps *pokemonInfoService) CreatePokemon(pokemonInfo entity.PokemonInfo) error {
	err := ps.repo.WritePokemon(pokemonInfo)

	return err
}

func (ps *pokemonInfoService) CreatePokemonConcu(itemsPerWorker int) []entity.PokemonInfo {
	pokemons := ps.repo.ConcuRSwWP(itemsPerWorker)

	return pokemons
}
