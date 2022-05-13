package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type server struct {
	*http.Server
}

func newServer(port string, routes *mux.Router) *server {
	s := &server{
		Server: &http.Server{
			Addr:         ":" + port,
			Handler:      routes,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
	return s
}
