package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/client"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/slice"
	"github.com/labstack/echo"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers2(e *echo.Group, service service.PokemonInfoService) {
	res := resource2{service}

	e.POST("/pokemon/:id", res.CreatePokemon)
	e.GET("/pokemon/concurrent", res.CreatePokemonConcu)
}

type resource2 struct {
	service service.PokemonInfoService
}

func (r resource2) CreatePokemon(c echo.Context) error {
	log.Println("Create a new pokemons in a CSV file")
	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error")
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}

	url := "https://pokeapi.co/api/v2/pokemon/"

	cliente := client.NewClient(url)
	bytes, err := cliente.GetRequest(c.Param("id"))
	if err != nil {
		log.Println("Error: ErrNotValidItemId")
		return errorhandler.ErrNotValidItemId
	}

	var data entity.PokemonInfo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println(err)
		return errorhandler.ErrFailedDependency
	}

	err1 := r.service.CreatePokemon(data)
	if err1 != nil {
		return errorhandler.ErrFailedDependency
	}

	return c.NoContent(http.StatusNoContent)
}

func (r resource2) CreatePokemonConcu(c echo.Context) error {
	log.Println("Read pokemons of CSV file concurrently")
	typeString := c.QueryParam("type")
	if typeString != "" && !slice.StringInSlice(typeString, []string{"even", "odd"}) {
		log.Println("No type parameter sent")
		return errorhandler.ErrSomeFieldAreNotValid
	}

	noItemsString := c.QueryParam("items")
	noItems, err := strconv.Atoi(noItemsString)
	if err != nil {
		return errorhandler.ErrSomeFieldAreNotValid
	}

	noItemsPerWorkersString := c.QueryParam("items_per_workers")
	noItemsPerWorkers, err := strconv.Atoi(noItemsPerWorkersString)
	if err != nil {
		return errorhandler.ErrSomeFieldAreNotValid
	}

	algo := r.service.CreatePokemonConcu(noItemsPerWorkers)

	if len(algo) < noItems {
		return errorhandler.ErrSomeFieldAreNotValid
	}

	cer := algo[:noItems]

	return c.JSON(http.StatusOK, cer)
}
