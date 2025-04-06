package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PriceType string

const (
	PriceTypeOneTime PriceType = "one-time"
	PriceTypeMonthly PriceType = "monthly"
	PriceTypeAnnual  PriceType = "annual"
)

// Price holds the schema definition for the Price entity.
type Price struct {
	ent.Schema
}

// Fields of the Price.
func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.String("stripe_id").
			NotEmpty(),
		field.Int("amount").
			Positive(),
		field.Int("quantity").
			Positive(),
		field.Enum("type").
			Values(string(PriceTypeOneTime), string(PriceTypeMonthly), string(PriceTypeAnnual)).
			Default(string(PriceTypeOneTime)),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("Prices").
			Unique().
			Required(),
		edge.From("product", Product.Type).
			Ref("prices").
			Unique().
			Required(),
	}
}
