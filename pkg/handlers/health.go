package handlers

import (
	"net/http"

	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/labstack/echo/v4"
)

const (
	routeNameHealth = "health"
)

type (
	Health struct {
		*services.TemplateRenderer
	}
)

func init() {
	Register(new(Health))
}

func (h *Health) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	return nil
}

func (h *Health) Routes(g *echo.Group) {
	g.GET("/up", h.Up).Name = routeNameHealth
}

func (h *Health) Up(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
