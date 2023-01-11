// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/googleoauthstate"
	"github.com/scarlet0725/prism-api/ent/user"
)

// GoogleOauthStateCreate is the builder for creating a GoogleOauthState entity.
type GoogleOauthStateCreate struct {
	config
	mutation *GoogleOauthStateMutation
	hooks    []Hook
}

// SetState sets the "state" field.
func (gosc *GoogleOauthStateCreate) SetState(s string) *GoogleOauthStateCreate {
	gosc.mutation.SetState(s)
	return gosc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gosc *GoogleOauthStateCreate) SetUserID(id int) *GoogleOauthStateCreate {
	gosc.mutation.SetUserID(id)
	return gosc
}

// SetUser sets the "user" edge to the User entity.
func (gosc *GoogleOauthStateCreate) SetUser(u *User) *GoogleOauthStateCreate {
	return gosc.SetUserID(u.ID)
}

// Mutation returns the GoogleOauthStateMutation object of the builder.
func (gosc *GoogleOauthStateCreate) Mutation() *GoogleOauthStateMutation {
	return gosc.mutation
}

// Save creates the GoogleOauthState in the database.
func (gosc *GoogleOauthStateCreate) Save(ctx context.Context) (*GoogleOauthState, error) {
	return withHooks[*GoogleOauthState, GoogleOauthStateMutation](ctx, gosc.sqlSave, gosc.mutation, gosc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gosc *GoogleOauthStateCreate) SaveX(ctx context.Context) *GoogleOauthState {
	v, err := gosc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gosc *GoogleOauthStateCreate) Exec(ctx context.Context) error {
	_, err := gosc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gosc *GoogleOauthStateCreate) ExecX(ctx context.Context) {
	if err := gosc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gosc *GoogleOauthStateCreate) check() error {
	if _, ok := gosc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "GoogleOauthState.state"`)}
	}
	if v, ok := gosc.mutation.State(); ok {
		if err := googleoauthstate.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "GoogleOauthState.state": %w`, err)}
		}
	}
	if _, ok := gosc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GoogleOauthState.user"`)}
	}
	return nil
}

func (gosc *GoogleOauthStateCreate) sqlSave(ctx context.Context) (*GoogleOauthState, error) {
	if err := gosc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gosc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gosc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gosc.mutation.id = &_node.ID
	gosc.mutation.done = true
	return _node, nil
}

func (gosc *GoogleOauthStateCreate) createSpec() (*GoogleOauthState, *sqlgraph.CreateSpec) {
	var (
		_node = &GoogleOauthState{config: gosc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: googleoauthstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleoauthstate.FieldID,
			},
		}
	)
	if value, ok := gosc.mutation.State(); ok {
		_spec.SetField(googleoauthstate.FieldState, field.TypeString, value)
		_node.State = value
	}
	if nodes := gosc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   googleoauthstate.UserTable,
			Columns: []string{googleoauthstate.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GoogleOauthStateCreateBulk is the builder for creating many GoogleOauthState entities in bulk.
type GoogleOauthStateCreateBulk struct {
	config
	builders []*GoogleOauthStateCreate
}

// Save creates the GoogleOauthState entities in the database.
func (goscb *GoogleOauthStateCreateBulk) Save(ctx context.Context) ([]*GoogleOauthState, error) {
	specs := make([]*sqlgraph.CreateSpec, len(goscb.builders))
	nodes := make([]*GoogleOauthState, len(goscb.builders))
	mutators := make([]Mutator, len(goscb.builders))
	for i := range goscb.builders {
		func(i int, root context.Context) {
			builder := goscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoogleOauthStateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, goscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, goscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, goscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (goscb *GoogleOauthStateCreateBulk) SaveX(ctx context.Context) []*GoogleOauthState {
	v, err := goscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (goscb *GoogleOauthStateCreateBulk) Exec(ctx context.Context) error {
	_, err := goscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goscb *GoogleOauthStateCreateBulk) ExecX(ctx context.Context) {
	if err := goscb.Exec(ctx); err != nil {
		panic(err)
	}
}
