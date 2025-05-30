package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/mikestefanello/backlite"

	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/ai"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
)

// AITask is an AI implementation of backlite.Task
// This represents the task that can be queued for execution via the task client and should contain everything
// that your queue processor needs to process the task.
type AITask struct {
	Prompt       string
	TaskID       int
	UserID       int
	AIClientName string
}

// Config satisfies the backlite.Task interface by providing configuration for the queue that these items will be
// placed into for execution.
func (t AITask) Config() backlite.QueueConfig {
	return backlite.QueueConfig{
		Name:        "AITask",
		MaxAttempts: 1,
		Timeout:     5 * time.Second,
		Backoff:     10 * time.Second,
		Retention: &backlite.Retention{
			Duration:   24 * time.Hour,
			OnlyFailed: false,
			Data: &backlite.RetainData{
				OnlyFailed: false,
			},
		},
	}
}

// NewAITaskQueue provides a Queue that can process AITask tasks
// The service container is provided so the subscriber can have access to the app dependencies.
// All queues must be registered in the Register() function.
// Whenever an AITask is added to the task client, it will be queued and eventually sent here for execution.

func updateAITask(ctx context.Context, c *services.Container, task AITask, status ai.Status, result string) error {
	_, err := c.ORM.AI.Update().
		Where(ai.HasUserWith(user.ID(task.UserID))).
		Where(ai.ID(task.TaskID)).
		SetStatus(status).
		SetResult(fmt.Appendf(nil, "%s", result)).
		SetCompletedAt(time.Now()).
		Save(ctx)
	return err
}
func NewAITaskQueue(c *services.Container) backlite.Queue {
	return backlite.NewQueue[AITask](func(ctx context.Context, task AITask) error {
		aiClient := c.AIs.GetClient(task.AIClientName)
		if aiClient == nil {
			return updateAITask(ctx, c, task, ai.StatusFailed, fmt.Sprintf("invalid ai client: %s", task.AIClientName))
		}

		result, err := aiClient.GenerateText(ctx, task.Prompt)
		if err != nil {
			return updateAITask(ctx, c, task, ai.StatusFailed, fmt.Sprintf("error generating text: %v", err))
		}

		return updateAITask(ctx, c, task, ai.StatusCompleted, result)
	})
}
