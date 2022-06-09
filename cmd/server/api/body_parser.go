package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func BodyParser(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		c.Response().SetStatusCode(http.StatusBadRequest)
		return err
	}

	return nil
}
