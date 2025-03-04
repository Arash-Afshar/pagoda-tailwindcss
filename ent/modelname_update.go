// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/modelname"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/predicate"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/user"
)

// ModelNameUpdate is the builder for updating ModelName entities.
type ModelNameUpdate struct {
	config
	hooks    []Hook
	mutation *ModelNameMutation
}

// Where appends a list predicates to the ModelNameUpdate builder.
func (mnu *ModelNameUpdate) Where(ps ...predicate.ModelName) *ModelNameUpdate {
	mnu.mutation.Where(ps...)
	return mnu
}

// SetFieldName sets the "field_name" field.
func (mnu *ModelNameUpdate) SetFieldName(s string) *ModelNameUpdate {
	mnu.mutation.SetFieldName(s)
	return mnu
}

// SetNillableFieldName sets the "field_name" field if the given value is not nil.
func (mnu *ModelNameUpdate) SetNillableFieldName(s *string) *ModelNameUpdate {
	if s != nil {
		mnu.SetFieldName(*s)
	}
	return mnu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mnu *ModelNameUpdate) SetUserID(id int) *ModelNameUpdate {
	mnu.mutation.SetUserID(id)
	return mnu
}

// SetUser sets the "user" edge to the User entity.
func (mnu *ModelNameUpdate) SetUser(u *User) *ModelNameUpdate {
	return mnu.SetUserID(u.ID)
}

// Mutation returns the ModelNameMutation object of the builder.
func (mnu *ModelNameUpdate) Mutation() *ModelNameMutation {
	return mnu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mnu *ModelNameUpdate) ClearUser() *ModelNameUpdate {
	mnu.mutation.ClearUser()
	return mnu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mnu *ModelNameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mnu.sqlSave, mnu.mutation, mnu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mnu *ModelNameUpdate) SaveX(ctx context.Context) int {
	affected, err := mnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mnu *ModelNameUpdate) Exec(ctx context.Context) error {
	_, err := mnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mnu *ModelNameUpdate) ExecX(ctx context.Context) {
	if err := mnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mnu *ModelNameUpdate) check() error {
	if v, ok := mnu.mutation.FieldName(); ok {
		if err := modelname.FieldNameValidator(v); err != nil {
			return &ValidationError{Name: "field_name", err: fmt.Errorf(`ent: validator failed for field "ModelName.field_name": %w`, err)}
		}
	}
	if mnu.mutation.UserCleared() && len(mnu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ModelName.user"`)
	}
	return nil
}

func (mnu *ModelNameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mnu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelname.Table, modelname.Columns, sqlgraph.NewFieldSpec(modelname.FieldID, field.TypeInt))
	if ps := mnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mnu.mutation.FieldName(); ok {
		_spec.SetField(modelname.FieldFieldName, field.TypeString, value)
	}
	if mnu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelname.UserTable,
			Columns: []string{modelname.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mnu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelname.UserTable,
			Columns: []string{modelname.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelname.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mnu.mutation.done = true
	return n, nil
}

// ModelNameUpdateOne is the builder for updating a single ModelName entity.
type ModelNameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ModelNameMutation
}

// SetFieldName sets the "field_name" field.
func (mnuo *ModelNameUpdateOne) SetFieldName(s string) *ModelNameUpdateOne {
	mnuo.mutation.SetFieldName(s)
	return mnuo
}

// SetNillableFieldName sets the "field_name" field if the given value is not nil.
func (mnuo *ModelNameUpdateOne) SetNillableFieldName(s *string) *ModelNameUpdateOne {
	if s != nil {
		mnuo.SetFieldName(*s)
	}
	return mnuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mnuo *ModelNameUpdateOne) SetUserID(id int) *ModelNameUpdateOne {
	mnuo.mutation.SetUserID(id)
	return mnuo
}

// SetUser sets the "user" edge to the User entity.
func (mnuo *ModelNameUpdateOne) SetUser(u *User) *ModelNameUpdateOne {
	return mnuo.SetUserID(u.ID)
}

// Mutation returns the ModelNameMutation object of the builder.
func (mnuo *ModelNameUpdateOne) Mutation() *ModelNameMutation {
	return mnuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mnuo *ModelNameUpdateOne) ClearUser() *ModelNameUpdateOne {
	mnuo.mutation.ClearUser()
	return mnuo
}

// Where appends a list predicates to the ModelNameUpdate builder.
func (mnuo *ModelNameUpdateOne) Where(ps ...predicate.ModelName) *ModelNameUpdateOne {
	mnuo.mutation.Where(ps...)
	return mnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mnuo *ModelNameUpdateOne) Select(field string, fields ...string) *ModelNameUpdateOne {
	mnuo.fields = append([]string{field}, fields...)
	return mnuo
}

// Save executes the query and returns the updated ModelName entity.
func (mnuo *ModelNameUpdateOne) Save(ctx context.Context) (*ModelName, error) {
	return withHooks(ctx, mnuo.sqlSave, mnuo.mutation, mnuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mnuo *ModelNameUpdateOne) SaveX(ctx context.Context) *ModelName {
	node, err := mnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mnuo *ModelNameUpdateOne) Exec(ctx context.Context) error {
	_, err := mnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mnuo *ModelNameUpdateOne) ExecX(ctx context.Context) {
	if err := mnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mnuo *ModelNameUpdateOne) check() error {
	if v, ok := mnuo.mutation.FieldName(); ok {
		if err := modelname.FieldNameValidator(v); err != nil {
			return &ValidationError{Name: "field_name", err: fmt.Errorf(`ent: validator failed for field "ModelName.field_name": %w`, err)}
		}
	}
	if mnuo.mutation.UserCleared() && len(mnuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ModelName.user"`)
	}
	return nil
}

func (mnuo *ModelNameUpdateOne) sqlSave(ctx context.Context) (_node *ModelName, err error) {
	if err := mnuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(modelname.Table, modelname.Columns, sqlgraph.NewFieldSpec(modelname.FieldID, field.TypeInt))
	id, ok := mnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ModelName.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, modelname.FieldID)
		for _, f := range fields {
			if !modelname.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != modelname.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mnuo.mutation.FieldName(); ok {
		_spec.SetField(modelname.FieldFieldName, field.TypeString, value)
	}
	if mnuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelname.UserTable,
			Columns: []string{modelname.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mnuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   modelname.UserTable,
			Columns: []string{modelname.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ModelName{config: mnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{modelname.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mnuo.mutation.done = true
	return _node, nil
}
