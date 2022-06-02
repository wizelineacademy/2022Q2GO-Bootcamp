// Package routes The Owl House API
//
// Documentation for the Owl House API
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package routes

import (
	"toh-api/pkg/parser"
)

type routeAbstract interface {
	RegisterRoutes(*parser.ApiService)
}

type routeService struct {
	api   *parser.ApiService
	route routeAbstract
}

func New(api *parser.ApiService, routes routeAbstract) routeService {
	return routeService{
		api:   api,
		route: routes,
	}
}

func (r routeService) Register() {
	r.route.RegisterRoutes(r.api)
}
