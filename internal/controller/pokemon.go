package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/errorhandler"
	"github.com/labstack/echo"
)

type resource struct {
	service service.PokemonService
}

func (r resource) GetPokemonById(c echo.Context) error {
	log.Println("Get infos about pokemons")

	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error")
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}

	client, err := r.service.FindPokemon(c, c.Param("id"))
	if err != nil {
		log.Println("Error")
		return errorhandler.ErrFailedDependency
	}

	return c.JSON(http.StatusOK, client)
}

// func GetPokemonConcurrent(c echo.Context) error {
// 	// noItemsString := c.QueryParam("items")
// 	// noItems, err := strconv.Atoi(noItemsString)
// 	// if err != nil {
// 	// 	return errorhandler.ErrSomeFieldAreNotValid
// 	// }

// 	// noItemsPerWorkersString := c.QueryParam("items_per_workers")
// 	// noItemsPerWorkers, err := strconv.Atoi(noItemsPerWorkersString)
// 	// if err != nil {
// 	// 	return errorhandler.ErrSomeFieldAreNotValid
// 	// }

// 	// f1, _ := os.Open(hook.Getcwd())
// 	// defer f1.Close()

// 	// algo := repository.ConcuRSwWP(f1, noItemsPerWorkers)
// 	// cer := algo[:noItems]

// 	// return c.JSON(http.StatusOK, cer)
// }
