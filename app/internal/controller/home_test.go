package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var homeController HomeController = NewHomeController()

func TestHome(t *testing.T) {
	// Create GET Request
	request, _ := http.NewRequest("GET", "/", nil)

	// Assign HTTP handler
	handler := http.HandlerFunc(homeController.Home)

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
