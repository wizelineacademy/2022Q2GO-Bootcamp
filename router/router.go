package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Router interface {
	GET(uri string, f func(http.ResponseWriter, *http.Request))
	POST(uri string, f func(http.ResponseWriter, *http.Request))
	SERVER(port string)
}

type router struct{}

var dispatcher = chi.NewRouter()

func NewRouter() Router {
	return &router{}
}

func (*router) GET(uri string, f func(http.ResponseWriter, *http.Request)) {
	dispatcher.Get(uri, f)
}

func (*router) POST(uri string, f func(http.ResponseWriter, *http.Request)) {
	dispatcher.Post(uri, f)
}

func (*router) SERVER(port string) {
	fmt.Printf("Chi HTTP Server running on port: %v", port)
	http.ListenAndServe(port, dispatcher)
}
