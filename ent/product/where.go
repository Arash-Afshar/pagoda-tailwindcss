// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// StripeID applies equality check predicate on the "stripe_id" field. It's identical to StripeIDEQ.
func StripeID(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStripeID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// StripeIDEQ applies the EQ predicate on the "stripe_id" field.
func StripeIDEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStripeID, v))
}

// StripeIDNEQ applies the NEQ predicate on the "stripe_id" field.
func StripeIDNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStripeID, v))
}

// StripeIDIn applies the In predicate on the "stripe_id" field.
func StripeIDIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldStripeID, vs...))
}

// StripeIDNotIn applies the NotIn predicate on the "stripe_id" field.
func StripeIDNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldStripeID, vs...))
}

// StripeIDGT applies the GT predicate on the "stripe_id" field.
func StripeIDGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldStripeID, v))
}

// StripeIDGTE applies the GTE predicate on the "stripe_id" field.
func StripeIDGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldStripeID, v))
}

// StripeIDLT applies the LT predicate on the "stripe_id" field.
func StripeIDLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldStripeID, v))
}

// StripeIDLTE applies the LTE predicate on the "stripe_id" field.
func StripeIDLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldStripeID, v))
}

// StripeIDContains applies the Contains predicate on the "stripe_id" field.
func StripeIDContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldStripeID, v))
}

// StripeIDHasPrefix applies the HasPrefix predicate on the "stripe_id" field.
func StripeIDHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldStripeID, v))
}

// StripeIDHasSuffix applies the HasSuffix predicate on the "stripe_id" field.
func StripeIDHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldStripeID, v))
}

// StripeIDEqualFold applies the EqualFold predicate on the "stripe_id" field.
func StripeIDEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldStripeID, v))
}

// StripeIDContainsFold applies the ContainsFold predicate on the "stripe_id" field.
func StripeIDContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldStripeID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrices applies the HasEdge predicate on the "prices" edge.
func HasPrices() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PricesTable, PricesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPricesWith applies the HasEdge predicate on the "prices" edge with a given conditions (other predicates).
func HasPricesWith(preds ...predicate.Price) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newPricesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
