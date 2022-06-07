package controller

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type controller struct {
}

type HomeController interface {
	Home(response http.ResponseWriter, request *http.Request)
}

func NewHomeController() HomeController {
	return &controller{}
}

// Home godoc
// @Summary Hello World
// @Description Hello World
// @Tags Description
// @Success 200 {object} HomeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router / [get]
func (*controller) Home(w http.ResponseWriter, r *http.Request) {
	//W will send the response
	w.Header().Add("Content-Type", "application/json")
	//Http response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HomeResponse{
		Message: "Hello World",
		Status:  true,
	})
}
