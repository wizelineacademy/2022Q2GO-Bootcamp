package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/luischitala/2022Q2GO-Bootcamp/config"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/controller"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/infrastructure/database"
	infrastructure "github.com/luischitala/2022Q2GO-Bootcamp/internal/infrastructure/router"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/usecase"

	_ "github.com/luischitala/2022Q2GO-Bootcamp/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title 2022 Wizeline's Go Bootcamp API
// @version 1.0
// @description This is a sample API for the final deliverable
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email lroman@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5050
// @BasePath /
func main() {
	// Configuration
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	//Repositories
	rcsv := repository.NewCsvRepository()
	// Database
	repo, err := database.NewPostgresRepository(c.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

	//Router
	r := mux.NewRouter()
	router := infrastructure.NewMuxRouter(r)

	//Use cases
	cu := usecase.NewCharacterUseCase(rcsv)

	//Controller
	hc := controller.NewHomeController()
	cc := controller.NewCharacterController(cu)

	//Routes
	router.GET("/", hc.Home)
	router.GET("/characters", cc.ListCharacter)
	router.GET("/charactersApi", cc.ListCharacterApi)
	router.GET("/readCsv", cc.ReadCsv)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.SERVE(c.Port)
}
