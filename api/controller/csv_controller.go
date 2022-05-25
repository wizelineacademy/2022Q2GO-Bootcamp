package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krmirandas/2022Q2GO-Bootcamp/api/services"
	"github.com/labstack/echo"
)

func CreateArticle(c echo.Context) error {
	//get article from request
	article := services.RetrievePokemon()
	//return the new article
	c.JSON(http.StatusCreated, article)
	return nil
}

func GetItems(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get infos about pokemons")
	initHeaders(writer)
	json.NewEncoder(writer).Encode(services.RetrievePokemon())
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
