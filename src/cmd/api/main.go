package main

import (
    "log"
	"fmt"
	"net/http"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/controller"
)

func main() {
    // Create a mux for routing incoming requests
    m := http.NewServeMux()

    //  URLs will be handled by a function
    m.HandleFunc("/", controller.GetGreetingHandler)
	m.HandleFunc("/getcsvdata", controller.GetCSVHandler)
	m.HandleFunc("/getexternalapidata", controller.GetExternalApiHandler)

	fmt.Println("Server Listening at port 8000")

    // Create a server listening on port 8000
    s := &http.Server{
        Addr:    ":8000",
        Handler: m,
    }
	
    // Continue to process new requests until an error occurs
    log.Fatal(s.ListenAndServe())
 
}