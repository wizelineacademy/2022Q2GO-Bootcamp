package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	controller "github.com/luischitala/2022Q2GO-Bootcamp/internal/controller"
	repository "github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
	usecase "github.com/luischitala/2022Q2GO-Bootcamp/internal/usecase"
)

var rcsv = repository.NewCsvRepository()
var cu = usecase.NewCharacterUseCase(rcsv)
var characterController controller.CharacterController = controller.NewCharacterController(cu)

func TestCharactersApi(t *testing.T) {

	// Create GET Request
	request, _ := http.NewRequest("GET", "/charactersApi", nil)

	// Assign HTTP handler
	handler := http.HandlerFunc(characterController.ListCharacterApi)

	// Record HTTP Response
	response := httptest.NewRecorder()

	// Dispatch
	handler.ServeHTTP(response, request)

	// Add Assertions on the HTTP status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Decode the HTTP response
}
