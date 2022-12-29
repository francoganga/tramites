package routes

import (
	"github.com/francoganga/go_reno2/pkg/controller"
	"github.com/labstack/echo/v4"
)

type tramite struct {
	controller.Controller
}

func (c *tramite) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "tramite/new"

	return c.RenderPage(ctx, page)
}

func (c *tramite) Post(ctx echo.Context) error {

	return ctx.JSON(200, map[string]string{"message": "ok"})
}
