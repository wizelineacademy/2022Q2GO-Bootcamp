package app

import (
	"github.com/jesusrevilla/capstone/internal/app/data/query"
	"github.com/jesusrevilla/capstone/internal/domain/data"
)

// Queries Contains all available query handlers of this app
type Queries struct {
	DataHandler   query.DataRequestHandler
	CoffeeHandler query.CoffeeRequestHandler
}

// DataServices Contains the grouped queries of the app layer
type DataServices struct {
	Queries Queries
}

type CoffeeServices struct {
	Queries Queries
}

// Services contains all exposed services of the application layer
type Services struct {
	DataServices   DataServices
	CoffeeServices CoffeeServices
}

// NewServices Bootstraps Application Layer dependencies
func NewServices(dataRepo data.Repository) Services {
	return Services{
		DataServices: DataServices{
			Queries: Queries{
				DataHandler: query.NewDataRequestHandler(dataRepo),
			},
		},
	}
}

// NewCoffeeService
func NewCoffeeService(coffeeRepo data.RepoCoffee) Services {
	return Services{
		CoffeeServices: CoffeeServices{
			Queries: Queries{
				CoffeeHandler: query.NewCoffeeRequestHandler(coffeeRepo),
			},
		},
	}
}
