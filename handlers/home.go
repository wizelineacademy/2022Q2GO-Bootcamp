package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TanZng/toh-api/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(writter http.ResponseWriter, request *http.Request) {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusOK)
		json.NewEncoder(writter).Encode(HomeResponse{
			Message: "Hello there!",
			Status:  http.StatusText(http.StatusOK),
		})
	}
}
