package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/luischitala/2022Q2GO-Bootcamp/config"
	"github.com/luischitala/2022Q2GO-Bootcamp/controller"
	"github.com/luischitala/2022Q2GO-Bootcamp/database"
	infrastructure "github.com/luischitala/2022Q2GO-Bootcamp/infrastructure/router"
	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/usecase"
)

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
	router.GET("/readCsv", cc.ReadCsv)
	router.SERVE(c.Port)
}
