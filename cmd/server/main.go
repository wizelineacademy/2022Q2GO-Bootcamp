package main

import (
	"fmt"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/controller"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/service"
	"github.com/krmirandas/2022Q2GO-Bootcamp/pkg/errorhandler"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Server is lift")
	fmt.Println("Port: 8000")

	app := echo.New()

	//set custom binder to validate payloads
	bi := hook.NewCustomBinderWithValidation()

	app.Binder = bi

	//set custom error handler
	app.HTTPErrorHandler = errorhandler.NewErrorHandler

	//set the port listener
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	dataCSV := "/home/krmirandas/Documentos/Proyectos/Wizeline/Proyecto/2022Q2GO-Bootcamp/data/pokemon.csv"
	writeCSV := "/home/krmirandas/Documentos/Proyectos/Wizeline/Proyecto/2022Q2GO-Bootcamp/data/pokemonAPI.csv"

	group := app.Group("/v2")
	controller.RegisterHandlers1(group,
		service.NewPokemonService(repository.NewPokemonRepo(dataCSV)),
	)
	controller.RegisterHandlers2(group,
		service.NewPokemonInfoService(repository.NewPokemonInfoRepo(writeCSV)),
	)

	fmt.Printf("API Management Listen to %s port in\n", port)

	app.Logger.Fatal(app.Start(":8000"))
}
