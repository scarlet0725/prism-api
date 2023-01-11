// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/googleoauthstate"
	"github.com/scarlet0725/prism-api/ent/predicate"
)

// GoogleOauthStateDelete is the builder for deleting a GoogleOauthState entity.
type GoogleOauthStateDelete struct {
	config
	hooks    []Hook
	mutation *GoogleOauthStateMutation
}

// Where appends a list predicates to the GoogleOauthStateDelete builder.
func (gosd *GoogleOauthStateDelete) Where(ps ...predicate.GoogleOauthState) *GoogleOauthStateDelete {
	gosd.mutation.Where(ps...)
	return gosd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gosd *GoogleOauthStateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GoogleOauthStateMutation](ctx, gosd.sqlExec, gosd.mutation, gosd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gosd *GoogleOauthStateDelete) ExecX(ctx context.Context) int {
	n, err := gosd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gosd *GoogleOauthStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: googleoauthstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleoauthstate.FieldID,
			},
		},
	}
	if ps := gosd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gosd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gosd.mutation.done = true
	return affected, err
}

// GoogleOauthStateDeleteOne is the builder for deleting a single GoogleOauthState entity.
type GoogleOauthStateDeleteOne struct {
	gosd *GoogleOauthStateDelete
}

// Exec executes the deletion query.
func (gosdo *GoogleOauthStateDeleteOne) Exec(ctx context.Context) error {
	n, err := gosdo.gosd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{googleoauthstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gosdo *GoogleOauthStateDeleteOne) ExecX(ctx context.Context) {
	gosdo.gosd.ExecX(ctx)
}
