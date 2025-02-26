package handlers

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/modelname"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/form"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/labstack/echo/v4"
)

const (
	routeNameModelModelName             = "model-modelname"
	routeNameModelModelNameList         = "model-modelname-list"
	routeNameModelModelNameNew          = "model-modelname-new"
	routeNameModelModelNameSubmitNew    = "model-modelname-submit-new"
	routeNameModelModelNameEdit         = "model-modelname-edit"
	routeNameModelModelNameSubmitEdit   = "model-modelname-submit-edit"
	routeNameModelModelNameSubmitDelete = "model-modelname-submit-delete"
)

type (
	ModelModelName struct {
		*services.TemplateRenderer
		db *ent.Client
	}

	ModelModelNameListData struct {
		ModelNames []*ent.ModelName
	}

	ModelModelNameData struct {
		ModelName *ent.ModelName
	}

	ModelModelNameForm struct {
		FieldName string `form:"field_name" validate:"required"`
		form.Submission
	}

	ModelModelNameEditParams struct {
		ID int `param:"id"`
	}
)

func NewModelModelName(c *services.Container) *ModelModelName {
	h := &ModelModelName{
		TemplateRenderer: c.TemplateRenderer,
		db:               c.ORM,
	}
	return h
}

func (h *ModelModelName) Routes() RouteMapMeta {
	return RouteMapMeta{
		Name:          "ModelName",
		MainRouteName: routeNameModelModelName,
		routeMaps: []RouteMap{
			{verb: "GET", path: "/modelname", name: routeNameModelModelName, handler: h.ModelName},
			{verb: "GET", path: "/modelname/list", name: routeNameModelModelNameList, handler: h.List},
			{verb: "GET", path: "/modelname/new", name: routeNameModelModelNameNew, handler: h.New},
			{verb: "POST", path: "/modelname/new", name: routeNameModelModelNameSubmitNew, handler: h.SubmitNew},
			{verb: "GET", path: "/modelname/edit/:id", name: routeNameModelModelNameEdit, handler: h.Edit},
			{verb: "PUT", path: "/modelname/edit/:id", name: routeNameModelModelNameSubmitEdit, handler: h.SubmitEdit},
			{verb: "DELETE", path: "/modelname/delete/:id", name: routeNameModelModelNameSubmitDelete, handler: h.SubmitDelete},
		},
	}
}

func (h *ModelModelName) ModelName(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageModelModelName

	return h.RenderPage(ctx, p)
}

func (h *ModelModelName) List(ctx echo.Context) error {
	modelNames, err := h.db.ModelName.
		Query().
		Where(modelname.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		All(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelModelNameList
	p.Data = ModelModelNameListData{
		ModelNames: modelNames,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelModelName) New(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelModelNameForm
	p.Form = form.Get[ModelModelNameForm](ctx)
	p.Data = &ModelModelNameData{}

	return h.RenderPage(ctx, p)
}

func (h *ModelModelName) SubmitNew(ctx echo.Context) error {
	var input ModelModelNameForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.New(ctx)
	}

	_, err := h.db.ModelName.Create().
		SetUser(ctx.Get(context.AuthenticatedUserKey).(*ent.User)).
		SetFieldName(input.FieldName).
		Save(ctx.Request().Context())

	if err != nil {
		return err
	}

	return h.ModelName(ctx)
}

func (h *ModelModelName) Edit(ctx echo.Context) error {
	params := new(ModelModelNameEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	modelName, err := h.db.ModelName.Query().
		Where(
			modelname.ID(params.ID),
			modelname.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Only(ctx.Request().Context())
	if err != nil {
		return err
	}

	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageModelModelNameForm
	f := form.Get[ModelModelNameForm](ctx)
	f.FieldName = modelName.FieldName
	p.Form = f
	p.Data = ModelModelNameData{
		ModelName: modelName,
	}

	return h.RenderPage(ctx, p)
}

func (h *ModelModelName) SubmitEdit(ctx echo.Context) error {
	params := new(ModelModelNameEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	var input ModelModelNameForm
	if err := form.Submit(ctx, &input); err != nil {
		ctx.Response().Header().Set("HX-Retarget", "#form-container")
		return h.Edit(ctx)
	}

	_, err := h.db.ModelName.Update().
		Where(
			modelname.ID(params.ID),
			modelname.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		SetFieldName(input.FieldName).
		Save(ctx.Request().Context())

	if err != nil {
		return err
	}

	return h.ModelName(ctx)
}

func (h *ModelModelName) SubmitDelete(ctx echo.Context) error {
	params := new(ModelModelNameEditParams)
	if err := ctx.Bind(params); err != nil {
		return err
	}

	err := h.db.ModelName.DeleteOneID(params.ID).
		Where(
			modelname.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID)),
		).
		Exec(ctx.Request().Context())
	if err != nil {
		return err
	}

	return h.List(ctx)
}
