package handlers

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/form"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/log"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	routeNameClickToEditMain   = "click-to-edit.main"
	routeNameClickToEditEdit   = "click-to-edit.edit"
	routeNameClickToEditSubmit = "click-to-edit.submit"
)

type (
	ClickToEdit struct {
		*services.TemplateRenderer
		orm *ent.Client
	}

	clickToEditForm struct {
		Name  string `form:"name" validate:"required"`
		Email string `form:"email" validate:"required,email"`
		form.Submission
	}

	clickToEditData struct {
		Mode  string
		Name  string
		Email string
	}
)

func init() {
	Register(new(ClickToEdit))
}

func (h *ClickToEdit) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.orm = c.ORM
	return nil
}

func (h *ClickToEdit) Routes(g *echo.Group) {
	g.GET("/examples/click-to-edit", h.MainPage).Name = routeNameClickToEditMain
	g.GET("/examples/click-to-edit/edit", h.EditablePage).Name = routeNameClickToEditEdit
	g.PUT("/examples/click-to-edit", h.Submit).Name = routeNameClickToEditSubmit
}

func (h *ClickToEdit) MainPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageClickToEdit
	p.Title = "htmx example: Click to Edit"
	p.Data = clickToEditData{
		Mode:  "view",
		Name:  p.AuthUser.Name,
		Email: p.AuthUser.Email,
	}

	return h.RenderPage(ctx, p)
}

func (h *ClickToEdit) EditablePage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageClickToEdit
	p.Title = "htmx example: Click to Edit"
	form := form.Get[clickToEditForm](ctx)
	form.Name = p.AuthUser.Name
	form.Email = p.AuthUser.Email
	p.Form = form
	p.Data = clickToEditData{
		Mode: "edit",
	}

	return h.RenderPage(ctx, p)
}

func (h *ClickToEdit) Submit(ctx echo.Context) error {
	var input clickToEditForm

	err := form.Submit(ctx, &input)
	switch err.(type) {
	case nil:
		log.Ctx(ctx).Info("received form data", "name", input.Name, "email", input.Email)
		if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
			updatedUser, err := h.orm.User.UpdateOneID(u.(*ent.User).ID).SetName(input.Name).SetEmail(input.Email).Save(ctx.Request().Context())
			if err != nil {
				return err
			}
			ctx.Set(context.AuthenticatedUserKey, updatedUser)
			return h.MainPage(ctx)
		}
		return err
	case validator.ValidationErrors:
		return h.MainPage(ctx)
	default:
		return err
	}

}
