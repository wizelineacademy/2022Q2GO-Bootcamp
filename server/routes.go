package server
import "net/http"

func initRoutes(){
	http.HandleFunc("/", index)
}