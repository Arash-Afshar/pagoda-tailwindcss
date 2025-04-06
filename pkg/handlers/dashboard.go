package handlers

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/middleware"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/labstack/echo/v4"
)

const (
	routeNameDashboard = "dashboard"
)

type (
	Dashboard struct {
		*services.TemplateRenderer
		db            *ent.Client
		price *ModelPrice
		product *ModelProduct
		modelname     *ModelModelName
		dashboardData *DashboardData
	}

	DashboardData struct {
		RouteMapMetas []RouteMapMeta
	}

	RouteMapMeta struct {
		routeMaps     []RouteMap
		Name          string
		MainRouteName string
	}

	RouteMap struct {
		verb    string
		path    string
		name    string
		handler func(ctx echo.Context) error
	}
)

func init() {
	Register(new(Dashboard))
}

func (h *Dashboard) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.db = c.ORM
	h.price = NewModelPrice(c)
	h.product = NewModelProduct(c)
	h.modelname = NewModelModelName(c)
	h.dashboardData = &DashboardData{}
	return nil
}

func (h *Dashboard) Routes(g *echo.Group) {
	g.GET("/dashboard", h.Home, middleware.RequireAuthentication()).Name = routeNameDashboard
	auth := g.Group("/dashboard", middleware.RequireAuthentication())
	// --------- Sub-Routes --------

	for _, route := range h.price.Routes().routeMaps {
		auth.Add(route.verb, route.path, route.handler).Name = route.name
	}
	h.dashboardData.RouteMapMetas = append(h.dashboardData.RouteMapMetas, h.price.Routes())

	for _, route := range h.product.Routes().routeMaps {
		auth.Add(route.verb, route.path, route.handler).Name = route.name
	}
	h.dashboardData.RouteMapMetas = append(h.dashboardData.RouteMapMetas, h.product.Routes())

	for _, route := range h.modelname.Routes().routeMaps {
		auth.Add(route.verb, route.path, route.handler).Name = route.name
	}
	h.dashboardData.RouteMapMetas = append(h.dashboardData.RouteMapMetas, h.modelname.Routes())
}

func (h *Dashboard) Home(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageDashboard
	p.Data = h.dashboardData

	return h.RenderPage(ctx, p)
}
