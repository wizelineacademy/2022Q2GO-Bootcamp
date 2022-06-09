package main

// Most of the code taken form https://github.com/pkritiotis/go-climb-clean-architecture-example as reference

import (
	"github.com/jesusrevilla/capstone/internal/app"
	"github.com/jesusrevilla/capstone/internal/inputport"
	"github.com/jesusrevilla/capstone/internal/interfaceadapter"
)

func main() {
	interfaceAdapterServices := interfaceadapter.NewServices()
	//appServices := app.NewServices(interfaceAdapterServices.DataRepository)
	appCoffeeServices := app.NewCoffeeService(interfaceAdapterServices.CoffeeRepository)
	//inputPortServices := inputport.NewServices(appServices)
	inputPortServices := inputport.NewServices(appCoffeeServices)
	inputPortServices.Server.ListenAndServe(":8080")
}
