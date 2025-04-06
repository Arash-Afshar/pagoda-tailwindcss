package handlers

import (
	"fmt"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/middleware"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/redirect"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/labstack/echo/v4"
)

// TODO: payment instead of subscription
// TODO: webhook and handling events

const (
	routeNameStripe         = "stripe"
	routeNameStripeCheckout = "stripe-checkout"
	routeNameStripeSuccess  = "stripe-success"
	routeNameStripeCancel   = "stripe-cancel"
)

type (
	Stripe struct {
		*services.TemplateRenderer
		Stripe *services.StripeClient
		Cache  *services.CacheClient
	}

	stripeSuccessData struct {
		SessionID string `query:"session_id"`
	}
)

func init() {
	Register(new(Stripe))
}

func (h *Stripe) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.Stripe = c.Stripe
	h.Cache = c.Cache
	return nil
}

func (h *Stripe) Routes(g *echo.Group) {
	g.GET("/stripe", h.Home, middleware.RequireAuthentication()).Name = routeNameStripe
	g.GET("/stripe/checkout", h.Checkout, middleware.RequireAuthentication()).Name = routeNameStripeCheckout
	g.GET("/stripe/success", h.Success, middleware.RequireAuthentication()).Name = routeNameStripeSuccess
	g.GET("/stripe/cancel", h.Cancel, middleware.RequireAuthentication()).Name = routeNameStripeCancel
}

func (h *Stripe) Home(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageStripe
	p.Metatags.Description = "Welcome to the homepage."
	p.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	p.Pager = page.NewPager(ctx, 4)
	p.Data = "Hello, World!"

	return h.RenderPage(ctx, p)
}

func (h *Stripe) Checkout(ctx echo.Context) error {
	user := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	customer, err := h.Stripe.GetCustomer(ctx.Request().Context(), h.Cache, user.ID, user.Email)
	if err != nil {
		return err
	}

	// TODO: figure out how to get the correct url
	successUrl := "http://localhost:8000/stripe/success?session_id={CHECKOUT_SESSION_ID}"
	cancelUrl := "http://localhost:8000/stripe/cancel"
	// TODO: set the price id (dv vs prod)
	session, err := h.Stripe.CheckoutSession(ctx.Request().Context(), successUrl, cancelUrl, customer.ID, "price_123", 1)
	if err != nil {
		return err
	}
	fmt.Println(session.URL)

	// Redirect to the Stripe checkout page
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return redirect.New(ctx).
		URL(session.URL).
		Go()
}

func (h *Stripe) Success(ctx echo.Context) error {
	user := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	var data stripeSuccessData
	if err := ctx.Bind(&data); err != nil {
		return err
	}

	err := h.Stripe.Success(ctx.Request().Context(), h.Cache, user.ID, data.SessionID)
	if err != nil {
		return err
	}

	customer, err := h.Stripe.GetCustomer(ctx.Request().Context(), h.Cache, user.ID, user.Email)
	if err != nil {
		return err
	}

	paymentsData, err := h.Stripe.GetStripeDataFromKV(ctx.Request().Context(), h.Cache, customer.ID)
	if err != nil {
		return err
	}

	d := paymentsData[data.SessionID]

	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageStripeSuccess
	p.Metatags.Description = "Payment successful."
	p.Data = d.PaymentID + ":" + d.PriceID

	return h.RenderPage(ctx, p)
}

func (h *Stripe) Cancel(ctx echo.Context) error {
	ctx.Logger().Info("Stripe checkout cancelled")
	return redirect.New(ctx).
		Route(routeNameStripe).
		Go()
}
