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
    page.Name = "tramite/show"

    return c.RenderPage(ctx, page)
}
