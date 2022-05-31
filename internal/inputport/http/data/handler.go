package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jesusrevilla/capstone/internal/app"
	"github.com/jesusrevilla/capstone/internal/app/data/query"
)

// Handler Data http request handler
type Handler struct {
	dataServices app.DataServices
}

// NewHandler Costructor
func NewHandler(app app.DataServices) *Handler {
	return &Handler{dataServices: app}
}

// GetDataIDURLParam contains the parameter identifier to be parsed by the handler
const GetDataIDURLParam = "dataId"

//Find Returns the data with the provided id
func (c Handler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataID := vars[GetDataIDURLParam]
	id, err := strconv.ParseInt(dataID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	data, err := c.dataServices.Queries.DataHandler.Handle(query.DataRequest{DataId: id})
	if err == nil && data == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
}
