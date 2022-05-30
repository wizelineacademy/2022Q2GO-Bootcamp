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

	client := services.NewClient("", "", "mx")
	t, err := client.GetRequest(c.Param("id"))
	if err != nil {
		log.Println("Error")
		return errorhandler.ErrFailedDependency
	}

	err = services.CreateCSV(*t)
	if err != nil {
		log.Println("Error")
		return errorhandler.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, t)
}
