package handlers

import (
	"net/http"

	"github.com/Arash-Afshar/pagoda-tailwindcss/config"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/middleware"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/gorilla/sessions"
	echomw "github.com/labstack/echo/v4/middleware"
)

// BuildRouter builds the router
func BuildRouter(c *services.Container) error {
	// Static files with proper cache control
	// funcmap.File() should be used in templates to append a cache key to the URL in order to break cache
	// after each server restart
	c.Web.Group("", middleware.CacheControl(c.Config.Cache.Expiration.StaticFile)).
		Static(config.StaticPrefix, config.StaticDir)

	// Non-static file route group
	g := c.Web.Group("")

	// Force HTTPS, if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
		middleware.Session(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		middleware.LoadAuthenticatedUser(c.Auth),
		middleware.ServeCachedPage(c.TemplateRenderer),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	// Error handler
	err := Error{c.TemplateRenderer}
	c.Web.HTTPErrorHandler = err.Page

	// Initialize and register all handlers
	for _, h := range GetHandlers() {
		if err := h.Init(c); err != nil {
			return err
		}

		h.Routes(g)
	}

	return nil
}
