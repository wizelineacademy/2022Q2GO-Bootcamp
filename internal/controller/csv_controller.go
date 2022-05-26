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
	//get pokemon from request
	pokemon := services.RetrievePokemon()
	//return the new pokemon
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
	c.JSON(http.StatusOK, pokemon)
	return nil
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
