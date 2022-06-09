package api

import "github.com/gofiber/fiber/v2"

type AuthService interface {
	AuthorizeToken(token string) error
}

type ApiService struct {
	api *fiber.App
}

func New(api *fiber.App) ApiService {
	return ApiService{
		api: api,
	}
}

func (s ApiService) Listen(addr string) error {
	return s.api.Listen(addr)
}

func (s ApiService) GetPublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Get(route, func(c *fiber.Ctx) error { return handler(c) })
}

func (s ApiService) PostPublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Post(route, func(c *fiber.Ctx) error { return handler(c) })
}

func (s ApiService) PutPublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Put(route, func(c *fiber.Ctx) error { return handler(c) })
}

func (s ApiService) DeletePublic(route string, handler func(*fiber.Ctx) error) {
	s.api.Delete(route, func(c *fiber.Ctx) error { return handler(c) })
}

func (s ApiService) GetPublicStatic(prefix, root string) {
	s.api.Static(prefix, root)
}
