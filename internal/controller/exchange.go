package controller

import (
	"net/http"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/services"
	"github.com/labstack/echo"
)

func CreateCsv(c echo.Context) error {
	client := services.NewBasicAuthClient("", "")
	t, _ := client.GetTodo()
	services.CreateCSV(t)

	c.JSON(http.StatusOK, t)
	return nil
}
