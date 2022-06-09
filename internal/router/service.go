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

import "toh-api/cmd/server/api"

type routeAbstract interface {
	RegisterRoutes(*api.ApiService)
}

type routeService struct {
	api   *api.ApiService
	route routeAbstract
}

func New(api *api.ApiService, routes routeAbstract) routeService {
	return routeService{
		api:   api,
		route: routes,
	}
}

func (r routeService) Register() {
	r.route.RegisterRoutes(r.api)
}
