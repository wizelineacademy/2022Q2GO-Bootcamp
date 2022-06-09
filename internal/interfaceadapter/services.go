package interfaceadapter

import (
	"github.com/jesusrevilla/capstone/internal/domain/data"
	"github.com/jesusrevilla/capstone/internal/interfaceadapter/storage/file"
)

// Services contains the exposed services of interface adapters
type Services struct {
	DataRepository   data.Repository
	CoffeeRepository data.RepoCoffee
}

// NewServices Instantiates the interface adapter services
func NewServices() Services {
	return Services{
		DataRepository:   file.NewRepo(),
		CoffeeRepository: file.NewRepoC(),
	}
}
