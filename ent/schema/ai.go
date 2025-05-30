package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TaskStatus string

const (
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

// AI holds the schema definition for the AI entity.
type AI struct {
	ent.Schema
}

// Fields of the AI.
func (AI) Fields() []ent.Field {
	return []ent.Field{
		field.String("ai_client_name").
			NotEmpty(),
		field.String("prompt").
			NotEmpty(),
		field.Enum("status").
			Values(string(TaskStatusRunning), string(TaskStatusCompleted), string(TaskStatusFailed)).
			Default(string(TaskStatusRunning)),
		field.Bytes("result").
			Default([]byte{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("completed_at").
			Optional(),
	}
}

// Edges of the AI.
func (AI) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("AIs").
			Unique().
			Required(),
	}
}
