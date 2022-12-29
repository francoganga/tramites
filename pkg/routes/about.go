package routes

import (
	"html/template"

	_ "github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/francoganga/go_reno2/pkg/controller"

	_ "github.com/francoganga/go_reno2/pkg/domain/candidato"
	_ "github.com/francoganga/go_reno2/pkg/domain/dependencia"
	_ "github.com/francoganga/go_reno2/pkg/domain/materia"
	_ "github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/labstack/echo/v4"
)

type (
	about struct {
		controller.Controller
	}

	aboutData struct {
		ShowCacheWarning bool
		FrontendTabs     []aboutTab
		BackendTabs      []aboutTab
	}

	aboutTab struct {
		Title string
		Body  template.HTML
	}
)

func (c *about) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "about"
	page.Title = "About"

	// This page will be cached!
	page.Cache.Enabled = false
	page.Cache.Tags = []string{"page_about", "page:list"}

	// ms := make([]*materia.Materia, 0)

	// a := aggregate.New(candidato.New("franco", "ganga", "asd@mail.com"), 2022, ms, user.New("admin"), &dependencia.Dependencia{Nombre: "ING", AreaSudocu: "aaa", Autorizante: user.New("admin")}, "LAAST")
	//
	// err := repo.Add(ctx.Request().Context(), a)
	//
	// if err != nil {
	//     return err
	// }

	// e, err := c.Container.ORM.Event.Create().SetType("tipo").SetPayload(map[string]string{"asasd": "ddddddddd"}).Save(ctx.Request().Context())
	//
	// if err != nil {
	//     return err
	// }
	//
	// t, err := c.Container.ORM.Tramite.Create().
	// SetAnoPresupuestario(2022).
	// SetCategoria("IEI").
	// SetCreatedAt(time.Now()).
	// SetLink("22222222").
	// AddEvents(e).
	// SetVersion(0).
	// Save(ctx.Request().Context())
	//
	// if err != nil {
	//     return err
	// }
	//
	// fmt.Printf("t=%v", t)

	// A simple example of how the Data field can contain anything you want to send to the templates
	// even though you wouldn't normally send markup like this
	page.Data = aboutData{
		ShowCacheWarning: true,
		FrontendTabs: []aboutTab{
			{
				Title: "HTMX",
				Body:  template.HTML(`Completes HTML as a hypertext by providing attributes to AJAXify anything and much more. Visit <a href="https://htmx.org/">htmx.org</a> to learn more.`),
			},
			{
				Title: "Alpine.js",
				Body:  template.HTML(`Drop-in, Vue-like functionality written directly in your markup. Visit <a href="https://alpinejs.dev/">alpinejs.dev</a> to learn more.`),
			},
			{
				Title: "Bulma",
				Body:  template.HTML(`Ready-to-use frontend components that you can easily combine to build responsive web interfaces with no JavaScript requirements. Visit <a href="https://bulma.io/">bulma.io</a> to learn more.`),
			},
		},
		BackendTabs: []aboutTab{
			{
				Title: "Echo",
				Body:  template.HTML(`High performance, extensible, minimalist Go web framework. Visit <a href="https://echo.labstack.com/">echo.labstack.com</a> to learn more.`),
			},
			{
				Title: "Ent",
				Body:  template.HTML(`Simple, yet powerful ORM for modeling and querying data. Visit <a href="https://entgo.io/">entgo.io</a> to learn more.`),
			},
		},
	}

	return c.RenderPage(ctx, page)
}
