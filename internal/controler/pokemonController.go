package controler

import (
	"net/http"
	"strconv"

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
	AddNewPokemons([]entity.Pokemon) ([]entity.Pokemon, error)
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
	pokemons, err := pKC.pokemonService.GetAllPokemons()
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, pokemons)
}

func (pKC *pokemonController) HandleGetPokemonById(ctx echo.Context) error {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	pokemon, err := pKC.pokemonService.GetPokemonById(intId)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, pokemon)
}

func (pKC *pokemonController) HandleAddNewPokemons(ctx echo.Context) error {
	pokemons, err := pKC.pokemonService.AddNewPokemons([]entity.Pokemon{})
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, pokemons)
}
