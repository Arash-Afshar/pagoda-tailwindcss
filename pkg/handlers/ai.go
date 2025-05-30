package handlers

import (
	"time"

	"github.com/mikestefanello/backlite"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/ai"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/context"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/form"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/middleware"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/page"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/tasks"
	"github.com/Arash-Afshar/pagoda-tailwindcss/templates"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	routeNameAITask       = "ai-task"
	routeNameAITaskSubmit = "ai-task.submit"
)

type (
	AITask struct {
		tasks *backlite.Client
		ais   services.AIClients
		db    *ent.Client
		*services.TemplateRenderer
	}

	aiTaskData struct {
		AIClientList     []string
		SelectedAIClient string
		Status           string
		Prompt           string
		Result           string
	}

	aiTaskForm struct {
		Prompt       string `form:"prompt" validate:"required"`
		AIClientName string `form:"ai_client_name" validate:"required"`
		form.Submission
	}
)

func init() {
	Register(new(AITask))
}

func (h *AITask) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.tasks = c.Tasks
	h.ais = c.AIs
	h.db = c.ORM
	return nil
}

func (h *AITask) Routes(g *echo.Group) {
	g.GET("/ai-task", h.Page, middleware.RequireAuthentication()).Name = routeNameAITask
	g.POST("/ai-task", h.Submit, middleware.RequireAuthentication()).Name = routeNameAITaskSubmit
}

func (h *AITask) latestAITask(ctx echo.Context) (*ent.AI, error) {
	aiTask, err := h.db.AI.
		Query().
		Where(ai.HasUserWith(user.ID(ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID))).
		Order(ent.Desc(ai.FieldCreatedAt)).
		First(ctx.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, fail(err, "unable to get ai task")
		}
	}

	return aiTask, nil
}

func (h *AITask) Page(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageAITask
	p.Title = "Create an AI task"
	p.Form = form.Get[aiTaskForm](ctx)

	aiTask, err := h.latestAITask(ctx)
	if err != nil {
		return fail(err, "unable to get ai task")
	}

	data := aiTaskData{
		AIClientList:     h.ais.GetClientList(),
		Status:           "",
		Result:           "",
		Prompt:           "",
		SelectedAIClient: "",
	}

	if aiTask != nil {
		data.Status = string(aiTask.Status)
		data.Result = string(aiTask.Result)
		data.Prompt = aiTask.Prompt
		data.SelectedAIClient = aiTask.AiClientName
	}

	p.Data = data

	return h.RenderPage(ctx, p)
}

func (h *AITask) Submit(ctx echo.Context) error {
	var input aiTaskForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.Page(ctx)
	default:
		return fail(err, "unable to submit ai task")
	}

	aiTask, err := h.latestAITask(ctx)
	if err != nil {
		return fail(err, "unable to get ai task")
	}
	if aiTask != nil {
		if aiTask.Status == ai.StatusRunning {
			return fail(err, "an ai task is already running")
		}
	}

	// Since the task does not exist or is not running, we can create a new task
	aiTask, err = h.db.AI.Create().
		SetAiClientName(input.AIClientName).
		SetPrompt(input.Prompt).
		SetUser(ctx.Get(context.AuthenticatedUserKey).(*ent.User)).
		Save(ctx.Request().Context())
	if err != nil {
		return fail(err, "unable to create ai task")
	}

	if h.ais.GetClient(input.AIClientName) == nil {
		return fail(err, "invalid ai client")
	}

	err = h.tasks.
		Add(tasks.AITask{
			Prompt:       input.Prompt,
			TaskID:       aiTask.ID,
			UserID:       ctx.Get(context.AuthenticatedUserKey).(*ent.User).ID,
			AIClientName: input.AIClientName,
		}).
		Wait(time.Duration(1) * time.Second).
		Save()
	if err != nil {
		return fail(err, "unable to create a task")
	}

	return h.Page(ctx)
}
