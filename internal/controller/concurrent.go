package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
	"github.com/labstack/echo"
)

func GetPokemonConcurrent(c echo.Context) error {
	// typeParam := c.QueryParam("type")

	// if !policies.IsValidParams(typeParam) {
	// 	return errorhandler.ErrSomeFieldAreNotValid
	// }

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
	// if err != nil {
	// 	log.Println("Error")

	// 	return errorhandler.ErrNotFoundAnyItemWithThisId
	// }

	// pokemon, err1 := services.RetrievePokemonById(c.Param("id"))

	// if err1 != nil {
	// 	return errorhandler.ErrNotFoundAnyItemWithThisId
	// }
	f1, _ := os.Open(hook.Getcwd())
	defer f1.Close()

	algo := repository.ConcuRSwWP(f1, noItemsPerWorkers)
	cer := algo[:noItems]
	//send the pokemon
	return c.JSON(http.StatusOK, cer)
}
