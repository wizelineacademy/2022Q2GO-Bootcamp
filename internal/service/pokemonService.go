package service

import (
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"
)

type pokemonRepository interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	GetPokemonItemsFromCSV(typeReading string, items, itemsPerWorkers int) ([]entity.Pokemon, error)
	AddNewPokemons(newPokemons []string) ([]entity.Pokemon, error)
}

type pokemonService struct {
	pokemonRepository pokemonRepository
}

func NewPokemonService(pokemonRepository pokemonRepository) *pokemonService {
	return &pokemonService{
		pokemonRepository: pokemonRepository,
	}
}

func (pKS *pokemonService) GetAllPokemons() ([]entity.Pokemon, error) {
	return pKS.pokemonRepository.GetAllPokemons()
}

func (pKS *pokemonService) GetPokemonById(id int) (*entity.Pokemon, error) {
	return pKS.pokemonRepository.GetPokemonById(id)
}

func (pKS *pokemonService) GetPokemonItemsFromCSV(typeReading string, items, itemsPerWorkers int) ([]entity.Pokemon, error) {
	return pKS.pokemonRepository.GetPokemonItemsFromCSV(typeReading, items, itemsPerWorkers)
}

func (pKS *pokemonService) AddNewPokemons(newPokemons []string) ([]entity.Pokemon, error) {
	return pKS.pokemonRepository.AddNewPokemons(newPokemons)
}
