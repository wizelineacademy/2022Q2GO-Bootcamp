package main

import (
	"fmt"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/controller"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/router"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
)

var (
	r                 router.Router                = router.NewRouter()
	csvFileController controller.CSVFileController = controller.NewCsvController(service.NewCsvService())
)

func main() {
	const port string = ":8080"

	r.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	r.POST("/sendCSVFile", csvFileController.PostCSVFile)
	r.SERVER(port)
}
