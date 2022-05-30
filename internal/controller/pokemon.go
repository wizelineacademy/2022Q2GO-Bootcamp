package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/services"
	"github.com/labstack/echo"
)

func GetItems(c echo.Context) error {
	log.Println("Get infos about pokemons")
	pokemon := services.RetrievePokemon()
	c.JSON(http.StatusCreated, pokemon)
	return nil
}

func GetPokemonById(c echo.Context) error {
	//first check is id valid or not
	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error")

		//Retrieves the infos about
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}

	pokemon, err1 := services.RetrievePokemonById(c.Param("id"))

	if err1 != nil {
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}
	//send the pokemon
	return c.JSON(http.StatusOK, pokemon)
}
