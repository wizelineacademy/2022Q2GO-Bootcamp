package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luischitala/2022Q2GO-Bootcamp/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	//W will send the response
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//Http response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Hello World",
			Status:  true,
		})
	}
}
