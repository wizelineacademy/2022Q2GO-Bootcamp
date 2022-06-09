package query

import (
	"github.com/jesusrevilla/capstone/internal/domain/data"
)

// Get Coffee results
type CoffeeResult struct {
	Id          int64
	Name        string
	Description string
	Origin      string
	Variety     string
}

// CoffeeRequestHandler provides an interface to handle a CoffeeRequest and return a *CoffeeResult
type CoffeeRequestHandler interface {
	Handle() ([]CoffeeResult, error)
}

type coffeeRequestHandler struct {
	repo data.RepoCoffee
}

// Coffee Handler Constructor
func NewCoffeeRequestHandler(repo data.RepoCoffee) CoffeeRequestHandler {
	return coffeeRequestHandler{repo: repo}
}

// Handles CoffeeRequest query
func (hc coffeeRequestHandler) Handle() ([]CoffeeResult, error) {
	res, err := hc.repo.GetCoffee()
	if err != nil {
		return nil, err
	}
	var result []CoffeeResult
	for _, coffee := range res {
		result = append(result, CoffeeResult{
			Id:          coffee.Id,
			Name:        coffee.Name,
			Description: coffee.Description,
			Origin:      coffee.Origin,
			Variety:     coffee.Variety,
		})
	}
	return result, err
}
