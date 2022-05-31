package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesusrevilla/capstone/internal/app"
	"github.com/jesusrevilla/capstone/internal/inputport/http/data"
)

// Server Represents the http server running for this service
type Server struct {
	appServices app.Services
	router      *mux.Router
}

// NewServer HTTP Server constructor
func NewServer(appServices app.Services) *Server {
	httpServer := &Server{appServices: appServices}
	httpServer.router = mux.NewRouter()
	httpServer.AddDataHTTPRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

// AddDataHTTPRoutes registers data route handlers
func (httpServer *Server) AddDataHTTPRoutes() {
	const dataHTTPRoutePath = "/findid"
	//Queries
	httpServer.router.HandleFunc(dataHTTPRoutePath+"/{"+data.GetDataIDURLParam+"}", data.NewHandler(httpServer.appServices.DataServices).Find).Methods("GET")
}

// ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
