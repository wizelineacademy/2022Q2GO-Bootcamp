package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luischitala/2022Q2GO-Bootcamp/controller"
	infrastructure "github.com/luischitala/2022Q2GO-Bootcamp/infrastructure/router"
	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/usecase"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	const port string = ":5050"

	//Repositories
	rcsv := repository.NewCsvRepository()
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
	router.SERVE(port)
}
