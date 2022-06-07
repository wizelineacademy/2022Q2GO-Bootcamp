package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	router "toh-api/internal/router"
	"toh-api/pkg/parser"
)

func Init() {
	app := fiber.New()

	newApi := parser.New(app)

	router.New(&newApi, router.CharacterRoutes()).Register()
	router.New(&newApi, router.HomeRoutes()).Register()
	router.New(&newApi, router.SwaggerRoutes()).Register()
	err := newApi.Listen(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
}
