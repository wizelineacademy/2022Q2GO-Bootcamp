package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type muxRouter struct {
	r *mux.Router
}

func NewMuxRouter(r *mux.Router) Router {
	return &muxRouter{
		r,
	}
}

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.r.HandleFunc(uri, f).Methods("POST")

}

func (m *muxRouter) SERVE(port string) {
	// CORS middleware implementation
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5050", "https://rickandmortyapi.com/api/character"},
	})
	handler := cors.Handler(m.r)
	fmt.Println(handler)
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, handler)
}
