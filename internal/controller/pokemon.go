package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/pagination"
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
	log.Println("Get infos about pokemons by Id")

	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error")
		return errorhandler.ErrNotValidItemId
	}

	response, err := r.service.FindPokemonById(c.Param("id"))
	if err != nil {
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}

	return c.JSON(http.StatusOK, response)
}

func (r resource) GetPokemon(c echo.Context) error {
	log.Println("Get infos about pokemons")
	count, err := r.service.Count()
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request(), count)
	response, err := r.service.FindPokemon(pages.Offset(), pages.Limit())
	if err != nil {
		return errorhandler.ErrFailedDependency
	}

	return c.JSON(http.StatusOK, response)
}
