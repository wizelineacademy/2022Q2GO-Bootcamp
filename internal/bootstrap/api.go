package bootstrap

import "github.com/labstack/echo/v4"

type app struct {
	server *echo.Echo
}

func ServeAPI() {
	a := new(app)
	a.setupServer()
	a.start()
}

func (a *app) start() {
	a.server.Start(":5000")
}

func (a *app) setupServer() {
	a.server = echo.New()
	baseGroup := a.server.Group("/api/v1")
	setAPIRoute(baseGroup)
}

func setAPIRoute(g *echo.Group) {

}
