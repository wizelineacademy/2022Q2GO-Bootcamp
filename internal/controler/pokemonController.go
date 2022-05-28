package controler

import (
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"
	"github.com/labstack/echo/v4"
)

type PokemonController interface {
	HandleGetAllPokemons(ctx echo.Context) error
	HandleGetPokemonById(ctx echo.Context) error
	HandleAddNewPokemons(ctx echo.Context) error
}

type PokemonService interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	AddNewPokemons([]entity.Pokemon) error
}

type pokemonController struct {
	pokemonService PokemonService
}

func NewPokemonController(pokemonService PokemonService) PokemonController {
	return &pokemonController{
		pokemonService: pokemonService,
	}
}

func (pKC *pokemonController) HandleGetAllPokemons(ctx echo.Context) error {
	return nil
}

func (pKC *pokemonController) HandleGetPokemonById(ctx echo.Context) error {
	return nil
}

func (pKC *pokemonController) HandleAddNewPokemons(ctx echo.Context) error {
	return nil
}
