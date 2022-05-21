package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"toh-api/cmd/server/api"
	routes "toh-api/cmd/server/api/router"
)

func Init() {
	app := fiber.New()

	newApi := api.New(app)

	routes.New(&newApi, routes.CharacterRoutes()).Register()
	routes.New(&newApi, routes.HomeRoutes()).Register()
	newApi.Listen(os.Getenv("API_PORT"))
}
