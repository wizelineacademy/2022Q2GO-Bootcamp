package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"toh-api/cmd/server/api"
	router "toh-api/internal/router"
)

func Init() {
	app := fiber.New()

	newApi := api.New(app)

	router.New(&newApi, router.CharacterRoutes()).Register()
	router.New(&newApi, router.HomeRoutes()).Register()
	router.New(&newApi, router.SwaggerRoutes()).Register()
	err := newApi.Listen(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
}
