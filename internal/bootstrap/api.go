package bootstrap

import (
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/controller"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/repository"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/service"
	"github.com/labstack/echo/v4"
)

type app struct {
	server *echo.Echo
}

func ServeAPI() {
	a := new(app)
	a.setupServer()
	a.start()
}

func (a *app) start() {
	a.server.Start(":5000")
}

func (a *app) setupServer() {
	a.server = echo.New()
	baseGroup := a.server.Group("/api/v1")
	setAPIRoute(baseGroup)
}

func setAPIRoute(g *echo.Group) {
	c := controller.NewPokemonController(
		service.NewPokemonService(
			repository.NewPokemonRepository("pokemons.csv"),
		),
	)

	g.GET("/all", c.HandleGetAllPokemons)
	g.GET("/:id", c.HandleGetPokemonById)
	g.GET("/items", c.HandleGetPokemonItemsFromCSV)
	g.PUT("/add/generation/:id", c.HandleAddNewPokemons)
}
