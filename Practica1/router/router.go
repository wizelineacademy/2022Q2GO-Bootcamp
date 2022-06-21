package router

import (
	"github.com/gorilla/mux"
	controller "main.com/controllers"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/readCSVFile", controller.HandlerReadFile).Methods("GET")
	r.HandleFunc("/searchID", controller.HandlerSearchID).Methods("GET")
}
