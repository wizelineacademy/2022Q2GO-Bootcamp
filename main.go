package main

import (
	"log"

	"github.com/esvarez/go-api/api"
)

func main() {
	log.Printf("api is running in port %s...\n", api.Port)
	api.Start()
}
