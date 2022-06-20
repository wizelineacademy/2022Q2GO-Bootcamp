package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/errorhandler"
	"github.com/labstack/echo"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers1(e *echo.Group, service service.PokemonService) {
	res := resource{service}

	e.GET("/pokemon/:id", res.GetPokemonById)
	e.GET("/pokemon", res.GetPokemon)
}

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

	response, err := r.service.FindPokemonById(c.Param("id"))
	if err != nil {
		return errorhandler.ErrFailedDependency
	}

	return c.JSON(http.StatusOK, response)
}

func (r resource) GetPokemon(c echo.Context) error {
	log.Println("Get infos about pokemons")

	_, err := r.service.FindPokemon()
	if err != nil {
		return errorhandler.ErrFailedDependency
	}

	return c.JSON(204, nil)
}
