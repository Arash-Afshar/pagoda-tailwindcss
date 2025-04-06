package handlers

import (
	"fmt"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/product"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/middleware"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/redirect"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/labstack/echo/v4"
)

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
		db     *ent.Client
	}

	stripeSuccessParams struct {
		SessionID string `query:"session_id"`
	}

	stripeSuccessData struct {
		Product   string
		Amounts   []int
		Quanities []int
	}
)

func init() {
	Register(new(Stripe))
}

func (h *Stripe) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.Stripe = c.Stripe
	h.Cache = c.Cache
	h.db = c.ORM
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

	return h.RenderPage(ctx, p)
}

func (h *Stripe) Checkout(ctx echo.Context) error {
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	customer, err := h.Stripe.GetCustomer(ctx.Request().Context(), h.Cache, u.ID, u.Email)
	if err != nil {
		return err
	}

	// TODO: figure out how to get the correct url
	successUrl := "http://localhost:8000/stripe/success?session_id={CHECKOUT_SESSION_ID}"
	cancelUrl := "http://localhost:8000/stripe/cancel"

	products, err := h.db.Product.
		Query().
		WithPrices().
		Where(product.HasUserWith(user.ID(u.ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}
	priceId := products[0].Edges.Prices[0].StripeID
	quantity := products[0].Edges.Prices[0].Quantity
	session, err := h.Stripe.CheckoutSession(ctx.Request().Context(), successUrl, cancelUrl, customer.ID, priceId, quantity)
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
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	var params stripeSuccessParams
	if err := ctx.Bind(&params); err != nil {
		return err
	}

	products, err := h.db.Product.
		Query().
		WithPrices().
		Where(product.HasUserWith(user.ID(u.ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}
	data := stripeSuccessData{
		Product: products[0].Name,
	}

	if params.SessionID != "" {
		err := h.Stripe.Success(ctx.Request().Context(), h.Cache, u.ID, params.SessionID)
		if err != nil {
			return err
		}
	}

	customer, err := h.Stripe.GetCustomer(ctx.Request().Context(), h.Cache, u.ID, u.Email)
	if err != nil {
		return err
	}

	paymentsData, err := h.Stripe.GetStripeDataFromKV(ctx.Request().Context(), h.Cache, customer.ID)
	if err != nil {
		return err
	}

	for _, d := range paymentsData {
		amount := 0
		quantity := 0
		for _, price := range products[0].Edges.Prices {
			if price.StripeID == d.PriceID {
				amount = price.Amount
				quantity = price.Quantity
				break
			}
		}
		data.Amounts = append(data.Amounts, amount)
		data.Quanities = append(data.Quanities, quantity)
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageStripeSuccess
	p.Data = data

	return h.RenderPage(ctx, p)
}

func (h *Stripe) Cancel(ctx echo.Context) error {
	ctx.Logger().Info("Stripe checkout cancelled")
	return redirect.New(ctx).
		Route(routeNameStripe).
		Go()
}
