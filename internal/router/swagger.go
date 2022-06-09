package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mvrilo/go-redoc"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
	"toh-api/cmd/server/api"
)

type swaggerRoutes struct{}

func SwaggerRoutes() swaggerRoutes {
	return swaggerRoutes{}
}

func getSwagger() fiber.Handler {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./api/swagger.yaml",
		SpecPath:    "/api/swagger.yaml",
		DocsPath:    "/swagger",
	}
	return fiberredoc.New(doc)
}

func (u swaggerRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/swagger", getSwagger())
	api.GetPublicStatic("/swagger.yaml", "./swagger.yaml")
}
