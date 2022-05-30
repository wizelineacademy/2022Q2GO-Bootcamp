package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/errorhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/services"
	"github.com/labstack/echo"
)

func CreateCsv(c echo.Context) error {
	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("Error")
		return errorhandler.ErrNotFoundAnyItemWithThisId
	}

	client := services.NewBasicAuthClient("", "")
	t, err := client.GetTodo(c.Param("id"))
	if err != nil {
		log.Println("Error")
		return errorhandler.ErrFailedDependency
	}

	err = services.CreateCSV(*t)
	if err != nil {
		log.Println("Error")
		return errorhandler.ErrInternalServerError
	}

	c.JSON(http.StatusOK, t)
	return nil
}
