package handlers

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/price"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/product"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/form"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/labstack/echo/v4"
)

const (
	routeNameModelPrice             = "model-price"
	routeNameModelPriceList         = "model-price-list"
	routeNameModelPriceNew          = "model-price-new"
	routeNameModelPriceSubmitNew    = "model-price-submit-new"
	routeNameModelPriceEdit         = "model-price-edit"
	routeNameModelPriceSubmitEdit   = "model-price-submit-edit"
	routeNameModelPriceSubmitDelete = "model-price-submit-delete"
)

type (
	ModelPrice struct {
		*services.TemplateRenderer
		db *ent.Client
	}

	ModelPriceListData struct {
		Prices []*ent.Price
	}

	ModelPriceData struct {
		Price       *ent.Price
		AllProducts []*ent.Product
	}

	ModelPriceForm struct {
		StripeID  string `form:"stripe_id" validate:"required"`
		Amount    int    `form:"amount" validate:"required"`
		Quantity  int    `form:"quantity" validate:"required"`
		Type      string `form:"type" validate:"required"`
		ProductID int    `form:"product_id" validate:"required"`
		form.Submission
	}

	ModelPriceEditParams struct {
		ID int `param:"id"`
	}
)

func NewModelPrice(c *services.Container) *ModelPrice {
	h := &ModelPrice{
		TemplateRenderer: c.TemplateRenderer,
		db:               c.ORM,
	}
	return h
}

func (h *ModelPrice) Routes() RouteMapMeta {
	return RouteMapMeta{
		Name:          "Price",
		MainRouteName: routeNameModelPrice,
		routeMaps: []RouteMap{
			{verb: "GET", path: "/price", name: routeNameModelPrice, handler: h.Price},
			{verb: "GET", path: "/price/list", name: routeNameModelPriceList, handler: h.List},
			{verb: "GET", path: "/price/new", name: routeNameModelPriceNew, handler: h.New},
			{verb: "POST", path: "/price/new", name: routeNameModelPriceSubmitNew, handler: h.SubmitNew},
			{verb: "GET", path: "/price/edit/:id", name: routeNameModelPriceEdit, handler: h.Edit},
			{verb: "PUT", path: "/price/edit/:id", name: routeNameModelPriceSubmitEdit, handler: h.SubmitEdit},
			{verb: "DELETE", path: "/price/delete/:id", name: routeNameModelPriceSubmitDelete, handler: h.SubmitDelete},
		},
	}
}

func (h *ModelPrice) Price(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageModelPrice

	return h.RenderPage(ctx, p)
}

func (h *ModelPrice) List(ctx echo.Context) error {
	prices, err := h.db.Price.
		Query().
		WithProduct().
		Where(price.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelPriceList
	p.Data = ModelPriceListData{
		Prices: prices,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelPrice) New(ctx echo.Context) error {
	products, err := h.db.Product.
		Query().
		Where(product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelPriceForm
	p.Form = form.Get[ModelPriceForm](ctx)
	p.Data = &ModelPriceData{
		AllProducts: products,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelPrice) SubmitNew(ctx echo.Context) error {
	var input ModelPriceForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.New(ctx)
	}

	_, err := h.db.Price.Create().
		SetUser(ctx.Get(context.AuthenticatedUserKey).(*ent.User)).
		SetStripeID(input.StripeID).
		SetAmount(input.Amount).
		SetQuantity(input.Quantity).
		SetType(price.Type(input.Type)).
		SetProductID(input.ProductID).
		Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.Price(ctx)
}

func (h *ModelPrice) Edit(ctx echo.Context) error {
	products, err := h.db.Product.
		Query().
		Where(product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}
	params := new(ModelPriceEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	price, err := h.db.Price.Query().
		WithProduct().
		Where(
			price.ID(params.ID),
			price.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Only(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelPriceForm
	f := form.Get[ModelPriceForm](ctx)
	f.StripeID = price.StripeID
	f.Amount = price.Amount
	f.Quantity = price.Quantity
	f.Type = string(price.Type)
	f.ProductID = price.Edges.Product.ID
	p.Form = f
	p.Data = ModelPriceData{
		Price:       price,
		AllProducts: products,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelPrice) SubmitEdit(ctx echo.Context) error {
	params := new(ModelPriceEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	var input ModelPriceForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.Edit(ctx)
	}

	_, err := h.db.Price.UpdateOneID(params.ID).
		Where(
			price.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		SetStripeID(input.StripeID).
		SetAmount(input.Amount).
		SetQuantity(input.Quantity).
		SetType(price.Type(input.Type)).
		SetProductID(input.ProductID).
		Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.Price(ctx)
}

func (h *ModelPrice) SubmitDelete(ctx echo.Context) error {
	params := new(ModelPriceEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	err := h.db.Price.DeleteOneID(params.ID).
		Where(
			price.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Exec(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.List(ctx)
}
