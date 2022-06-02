package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	routes2 "toh-api/internal/router"
	"toh-api/pkg/parser"
)

func Init() {
	app := fiber.New()

	newApi := parser.New(app)

	routes2.New(&newApi, routes2.CharacterRoutes()).Register()
	routes2.New(&newApi, routes2.HomeRoutes()).Register()
	routes2.New(&newApi, routes2.SwaggerRoutes()).Register()
	err := newApi.Listen(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
}
