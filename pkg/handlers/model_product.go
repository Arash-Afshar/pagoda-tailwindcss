package handlers

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
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
	routeNameModelProduct             = "model-product"
	routeNameModelProductList         = "model-product-list"
	routeNameModelProductNew          = "model-product-new"
	routeNameModelProductSubmitNew    = "model-product-submit-new"
	routeNameModelProductEdit         = "model-product-edit"
	routeNameModelProductSubmitEdit   = "model-product-submit-edit"
	routeNameModelProductSubmitDelete = "model-product-submit-delete"
)

type (
	ModelProduct struct {
		*services.TemplateRenderer
		db *ent.Client
	}

	ModelProductListData struct {
		Products []*ent.Product
	}

	ModelProductData struct {
		Product *ent.Product
	}

	ModelProductForm struct {
		StripeID string `form:"stripe_id" validate:"required"`
		Name     string `form:"name" validate:"required"`
		form.Submission
	}

	ModelProductEditParams struct {
		ID int `param:"id"`
	}
)

func NewModelProduct(c *services.Container) *ModelProduct {
	h := &ModelProduct{
		TemplateRenderer: c.TemplateRenderer,
		db:               c.ORM,
	}
	return h
}

func (h *ModelProduct) Routes() RouteMapMeta {
	return RouteMapMeta{
		Name:          "Product",
		MainRouteName: routeNameModelProduct,
		routeMaps: []RouteMap{
			{verb: "GET", path: "/product", name: routeNameModelProduct, handler: h.Product},
			{verb: "GET", path: "/product/list", name: routeNameModelProductList, handler: h.List},
			{verb: "GET", path: "/product/new", name: routeNameModelProductNew, handler: h.New},
			{verb: "POST", path: "/product/new", name: routeNameModelProductSubmitNew, handler: h.SubmitNew},
			{verb: "GET", path: "/product/edit/:id", name: routeNameModelProductEdit, handler: h.Edit},
			{verb: "PUT", path: "/product/edit/:id", name: routeNameModelProductSubmitEdit, handler: h.SubmitEdit},
			{verb: "DELETE", path: "/product/delete/:id", name: routeNameModelProductSubmitDelete, handler: h.SubmitDelete},
		},
	}
}

func (h *ModelProduct) Product(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageModelProduct

	return h.RenderPage(ctx, p)
}

func (h *ModelProduct) List(ctx echo.Context) error {
	products, err := h.db.Product.
		Query().
		Where(product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelProductList
	p.Data = ModelProductListData{
		Products: products,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelProduct) New(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelProductForm
	p.Form = form.Get[ModelProductForm](ctx)
	p.Data = &ModelProductData{}

	return h.RenderPage(ctx, p)
}

func (h *ModelProduct) SubmitNew(ctx echo.Context) error {
	var input ModelProductForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.New(ctx)
	}

	_, err := h.db.Product.Create().
		SetUser(ctx.Get(context.AuthenticatedUserKey).(*ent.User)).
		SetStripeID(input.StripeID).
		SetName(input.Name).
		Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.Product(ctx)
}

func (h *ModelProduct) Edit(ctx echo.Context) error {
	params := new(ModelProductEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	product, err := h.db.Product.Query().
		Where(
			product.ID(params.ID),
			product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Only(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelProductForm
	f := form.Get[ModelProductForm](ctx)
	f.StripeID = product.StripeID
	f.Name = product.Name
	p.Form = f
	p.Data = ModelProductData{
		Product: product,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelProduct) SubmitEdit(ctx echo.Context) error {
	params := new(ModelProductEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	var input ModelProductForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.Edit(ctx)
	}

	_, err := h.db.Product.UpdateOneID(params.ID).
		Where(
			product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		SetStripeID(input.StripeID).
		SetName(input.Name).
		Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.Product(ctx)
}

func (h *ModelProduct) SubmitDelete(ctx echo.Context) error {
	params := new(ModelProductEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	err := h.db.Product.DeleteOneID(params.ID).
		Where(
			product.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Exec(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.List(ctx)
}
