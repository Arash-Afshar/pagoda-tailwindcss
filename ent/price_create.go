// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/price"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/product"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
)

// PriceCreate is the builder for creating a Price entity.
type PriceCreate struct {
	config
	mutation *PriceMutation
	hooks    []Hook
}

// SetStripeID sets the "stripe_id" field.
func (pc *PriceCreate) SetStripeID(s string) *PriceCreate {
	pc.mutation.SetStripeID(s)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PriceCreate) SetAmount(i int) *PriceCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetQuantity sets the "quantity" field.
func (pc *PriceCreate) SetQuantity(i int) *PriceCreate {
	pc.mutation.SetQuantity(i)
	return pc
}

// SetType sets the "type" field.
func (pc *PriceCreate) SetType(pr price.Type) *PriceCreate {
	pc.mutation.SetType(pr)
	return pc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pc *PriceCreate) SetNillableType(pr *price.Type) *PriceCreate {
	if pr != nil {
		pc.SetType(*pr)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PriceCreate) SetCreatedAt(t time.Time) *PriceCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PriceCreate) SetNillableCreatedAt(t *time.Time) *PriceCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PriceCreate) SetUserID(id int) *PriceCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PriceCreate) SetUser(u *User) *PriceCreate {
	return pc.SetUserID(u.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (pc *PriceCreate) SetProductID(id int) *PriceCreate {
	pc.mutation.SetProductID(id)
	return pc
}

// SetProduct sets the "product" edge to the Product entity.
func (pc *PriceCreate) SetProduct(p *Product) *PriceCreate {
	return pc.SetProductID(p.ID)
}

// Mutation returns the PriceMutation object of the builder.
func (pc *PriceCreate) Mutation() *PriceMutation {
	return pc.mutation
}

// Save creates the Price in the database.
func (pc *PriceCreate) Save(ctx context.Context) (*Price, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PriceCreate) SaveX(ctx context.Context) *Price {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PriceCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PriceCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PriceCreate) defaults() {
	if _, ok := pc.mutation.GetType(); !ok {
		v := price.DefaultType
		pc.mutation.SetType(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := price.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PriceCreate) check() error {
	if _, ok := pc.mutation.StripeID(); !ok {
		return &ValidationError{Name: "stripe_id", err: errors.New(`ent: missing required field "Price.stripe_id"`)}
	}
	if v, ok := pc.mutation.StripeID(); ok {
		if err := price.StripeIDValidator(v); err != nil {
			return &ValidationError{Name: "stripe_id", err: fmt.Errorf(`ent: validator failed for field "Price.stripe_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Price.amount"`)}
	}
	if v, ok := pc.mutation.Amount(); ok {
		if err := price.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Price.amount": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Price.quantity"`)}
	}
	if v, ok := pc.mutation.Quantity(); ok {
		if err := price.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "Price.quantity": %w`, err)}
		}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Price.type"`)}
	}
	if v, ok := pc.mutation.GetType(); ok {
		if err := price.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Price.type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Price.created_at"`)}
	}
	if len(pc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Price.user"`)}
	}
	if len(pc.mutation.ProductIDs()) == 0 {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "Price.product"`)}
	}
	return nil
}

func (pc *PriceCreate) sqlSave(ctx context.Context) (*Price, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PriceCreate) createSpec() (*Price, *sqlgraph.CreateSpec) {
	var (
		_node = &Price{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(price.Table, sqlgraph.NewFieldSpec(price.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.StripeID(); ok {
		_spec.SetField(price.FieldStripeID, field.TypeString, value)
		_node.StripeID = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(price.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.Quantity(); ok {
		_spec.SetField(price.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(price.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(price.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   price.UserTable,
			Columns: []string{price.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_prices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   price.ProductTable,
			Columns: []string{price.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_prices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PriceCreateBulk is the builder for creating many Price entities in bulk.
type PriceCreateBulk struct {
	config
	err      error
	builders []*PriceCreate
}

// Save creates the Price entities in the database.
func (pcb *PriceCreateBulk) Save(ctx context.Context) ([]*Price, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Price, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PriceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PriceCreateBulk) SaveX(ctx context.Context) []*Price {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PriceCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PriceCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
