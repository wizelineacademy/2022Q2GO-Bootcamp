package routes

import (
	"github.com/gofiber/fiber/v2"
	"toh-api/cmd/server/api"
)

type homeRoutes struct{}

func HomeRoutes() homeRoutes {
	return homeRoutes{}
}

type Home struct {
	Msg string
}

func (u homeRoutes) getHome(ctx *fiber.Ctx) error {
	home := Home{
		Msg: "Hello world",
	}
	return ctx.JSON(home)
}

func (u homeRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/", u.getHome)
}
