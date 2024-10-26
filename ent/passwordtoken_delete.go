// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/passwordtoken"
	"github.com/Arash-Afshar/pagoda-tailwindcss/ent/predicate"
)

// PasswordTokenDelete is the builder for deleting a PasswordToken entity.
type PasswordTokenDelete struct {
	config
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// Where appends a list predicates to the PasswordTokenDelete builder.
func (ptd *PasswordTokenDelete) Where(ps ...predicate.PasswordToken) *PasswordTokenDelete {
	ptd.mutation.Where(ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *PasswordTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ptd.sqlExec, ptd.mutation, ptd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *PasswordTokenDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *PasswordTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(passwordtoken.Table, sqlgraph.NewFieldSpec(passwordtoken.FieldID, field.TypeInt))
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ptd.mutation.done = true
	return affected, err
}

// PasswordTokenDeleteOne is the builder for deleting a single PasswordToken entity.
type PasswordTokenDeleteOne struct {
	ptd *PasswordTokenDelete
}

// Where appends a list predicates to the PasswordTokenDelete builder.
func (ptdo *PasswordTokenDeleteOne) Where(ps ...predicate.PasswordToken) *PasswordTokenDeleteOne {
	ptdo.ptd.mutation.Where(ps...)
	return ptdo
}

// Exec executes the deletion query.
func (ptdo *PasswordTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{passwordtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *PasswordTokenDeleteOne) ExecX(ctx context.Context) {
	if err := ptdo.Exec(ctx); err != nil {
		panic(err)
	}
}
