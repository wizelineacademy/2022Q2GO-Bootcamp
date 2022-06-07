package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/utils"
	"github.com/labstack/echo/v4"
)

type pokemonControllerI interface {
	HandleGetAllPokemons(ctx echo.Context) error
	HandleGetPokemonById(ctx echo.Context) error
	HandleGetPokemonItemsFromCSV(ctx echo.Context) error
	HandleAddNewPokemons(ctx echo.Context) error
}

type pokemonService interface {
	GetAllPokemons() ([]entity.Pokemon, error)
	GetPokemonById(id int) (*entity.Pokemon, error)
	GetPokemonItemsFromCSV(typeReading string, items, itemsPerWorkers int) ([]entity.Pokemon, error)
	AddNewPokemons(newPokemons []string) ([]entity.Pokemon, error)
}

type pokemonController struct {
	pokemonService pokemonService
}

func NewPokemonController(pokemonService pokemonService) pokemonControllerI {
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
		return ctx.JSON(http.StatusNotFound, utils.ErrorJsonStructResponse("pokemon not found"))
	}

	return ctx.JSON(http.StatusOK, pokemon)
}

func (pKC *pokemonController) HandleGetPokemonItemsFromCSV(ctx echo.Context) error {
	typeReading := ctx.QueryParam("type")
	if !utils.IsValidReadingType(typeReading) {
		return ctx.JSON(http.StatusBadRequest, utils.ErrorJsonStructResponse("invalid type"))
	}

	noItemsString := ctx.QueryParam("items")
	noItems, err := strconv.Atoi(noItemsString)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ErrorJsonStructResponse("invalid items"))
	}

	noItemsPerWorkersString := ctx.QueryParam("items_per_workers")
	noItemsPerWorkers, err := strconv.Atoi(noItemsPerWorkersString)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ErrorJsonStructResponse("invalid items_per_workers"))
	}

	pokemons, err := pKC.pokemonService.GetPokemonItemsFromCSV(typeReading, noItems, noItemsPerWorkers)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusNotFound, utils.ErrorJsonStructResponse("pokemons not found"))
	}

	return ctx.JSON(http.StatusOK, pokemons)
}

func (pKC *pokemonController) HandleAddNewPokemons(ctx echo.Context) error {
	generationID := ctx.Param("id")
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	resp, err := http.Get("https://pokeapi.co/api/v2/generation/" + generationID)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return err
	}

	newPokemons := jsonBody["pokemon_species"].([]interface{})
	pokemonsNames := make([]string, len(newPokemons))

	for i, pokemon := range newPokemons {
		pokemonParsed := pokemon.(map[string]interface{})
		pokemonName := pokemonParsed["name"].(string)
		pokemonsNames[i] = strings.ToLower(pokemonName)
	}

	pokemons, err := pKC.pokemonService.AddNewPokemons(pokemonsNames)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, pokemons)
}
