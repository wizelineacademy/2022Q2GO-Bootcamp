package web

import (
	"encoding/json"
	"net/http"
)

type ResponseAPI struct {
	Status int `json:"status"`
	Result any `json:"result"`
}

func Success(result any, status int) *ResponseAPI {
	return &ResponseAPI{
		Status: status,
		Result: result,
	}
}

func (r ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
