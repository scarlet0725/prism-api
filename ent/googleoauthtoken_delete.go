// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/predicate"
)

// GoogleOauthTokenDelete is the builder for deleting a GoogleOauthToken entity.
type GoogleOauthTokenDelete struct {
	config
	hooks    []Hook
	mutation *GoogleOauthTokenMutation
}

// Where appends a list predicates to the GoogleOauthTokenDelete builder.
func (gotd *GoogleOauthTokenDelete) Where(ps ...predicate.GoogleOauthToken) *GoogleOauthTokenDelete {
	gotd.mutation.Where(ps...)
	return gotd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gotd *GoogleOauthTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GoogleOauthTokenMutation](ctx, gotd.sqlExec, gotd.mutation, gotd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gotd *GoogleOauthTokenDelete) ExecX(ctx context.Context) int {
	n, err := gotd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gotd *GoogleOauthTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: googleoauthtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleoauthtoken.FieldID,
			},
		},
	}
	if ps := gotd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gotd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gotd.mutation.done = true
	return affected, err
}

// GoogleOauthTokenDeleteOne is the builder for deleting a single GoogleOauthToken entity.
type GoogleOauthTokenDeleteOne struct {
	gotd *GoogleOauthTokenDelete
}

// Exec executes the deletion query.
func (gotdo *GoogleOauthTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := gotdo.gotd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{googleoauthtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gotdo *GoogleOauthTokenDeleteOne) ExecX(ctx context.Context) {
	gotdo.gotd.ExecX(ctx)
}