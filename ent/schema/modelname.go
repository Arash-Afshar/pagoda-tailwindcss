package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ModelName holds the schema definition for the ModelName entity.
type ModelName struct {
	ent.Schema
}

// Fields of the ModelName.
func (ModelName) Fields() []ent.Field {
	return []ent.Field{
		field.String("field_name").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the ModelName.
func (ModelName) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("ModelNames").
			Unique().
			Required(),
	}
}
