package tasks

import (
	"github.com/Arash-Afshar/pagoda-tailwindcss/pkg/services"
)

// Register registers all task queues with the task client
func Register(c *services.Container) {
	c.Tasks.Register(NewExampleTaskQueue(c))
	c.Tasks.Register(NewAITaskQueue(c))
}
