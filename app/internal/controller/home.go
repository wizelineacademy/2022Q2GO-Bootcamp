package controller

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
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
// @Tags General
// @Accept  json
// @Produce  json
// @Success 200 {object} HomeResponse
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
