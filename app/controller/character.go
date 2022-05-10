package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/usecase"
)

type CharResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type cc struct {
	cu usecase.Character
}

type CharacterController interface {
	ListCharacter(response http.ResponseWriter, request *http.Request)
	ReadCsv(response http.ResponseWriter, request *http.Request)
}

func NewCharacterController(cu usecase.Character) CharacterController {
	return &cc{
		cu,
	}
}

func (c *cc) ListCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	pageStr := r.URL.Query().Get("page")
	var page = uint64(0)
	if pageStr != "" {
		page, err = strconv.ParseUint(pageStr, 10, 53)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	characters, err := repository.ListCharacter(r.Context(), page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func (c *cc) ReadCsv(w http.ResponseWriter, r *http.Request) {
	characters, err := c.cu.ReadCsv()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}
