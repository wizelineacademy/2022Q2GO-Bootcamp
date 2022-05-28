package main

import (
	"fmt"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/controller"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/errorhandler"
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

	app.GET("/pokemon", controller.GetItems)
	app.GET("/pokemon/:id", controller.GetPokemonById)
	app.GET("/exchange", controller.CreateCsv)

	fmt.Printf("API Management Listen to %s port in\n", port)

	app.Logger.Fatal(app.Start(":8000"))
}
