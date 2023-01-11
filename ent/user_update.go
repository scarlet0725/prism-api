// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/event"
	"github.com/scarlet0725/prism-api/ent/googleoauthstate"
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/predicate"
	"github.com/scarlet0725/prism-api/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UserUpdate) SetUserID(s string) *UserUpdate {
	uu.mutation.SetUserID(s)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(b []byte) *UserUpdate {
	uu.mutation.SetPassword(b)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of the "first_name" field.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of the "last_name" field.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// SetIsAdminVerified sets the "is_admin_verified" field.
func (uu *UserUpdate) SetIsAdminVerified(b bool) *UserUpdate {
	uu.mutation.SetIsAdminVerified(b)
	return uu
}

// SetNillableIsAdminVerified sets the "is_admin_verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsAdminVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsAdminVerified(*b)
	}
	return uu
}

// SetDeleteProtected sets the "delete_protected" field.
func (uu *UserUpdate) SetDeleteProtected(b bool) *UserUpdate {
	uu.mutation.SetDeleteProtected(b)
	return uu
}

// SetNillableDeleteProtected sets the "delete_protected" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeleteProtected(b *bool) *UserUpdate {
	if b != nil {
		uu.SetDeleteProtected(*b)
	}
	return uu
}

// SetAPIKey sets the "api_key" field.
func (uu *UserUpdate) SetAPIKey(s string) *UserUpdate {
	uu.mutation.SetAPIKey(s)
	return uu
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAPIKey(s *string) *UserUpdate {
	if s != nil {
		uu.SetAPIKey(*s)
	}
	return uu
}

// ClearAPIKey clears the value of the "api_key" field.
func (uu *UserUpdate) ClearAPIKey() *UserUpdate {
	uu.mutation.ClearAPIKey()
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetGoogleOauthTokensID sets the "google_oauth_tokens" edge to the GoogleOauthToken entity by ID.
func (uu *UserUpdate) SetGoogleOauthTokensID(id int) *UserUpdate {
	uu.mutation.SetGoogleOauthTokensID(id)
	return uu
}

// SetNillableGoogleOauthTokensID sets the "google_oauth_tokens" edge to the GoogleOauthToken entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableGoogleOauthTokensID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetGoogleOauthTokensID(*id)
	}
	return uu
}

// SetGoogleOauthTokens sets the "google_oauth_tokens" edge to the GoogleOauthToken entity.
func (uu *UserUpdate) SetGoogleOauthTokens(g *GoogleOauthToken) *UserUpdate {
	return uu.SetGoogleOauthTokensID(g.ID)
}

// SetGoogleOauthStatesID sets the "google_oauth_states" edge to the GoogleOauthState entity by ID.
func (uu *UserUpdate) SetGoogleOauthStatesID(id int) *UserUpdate {
	uu.mutation.SetGoogleOauthStatesID(id)
	return uu
}

// SetNillableGoogleOauthStatesID sets the "google_oauth_states" edge to the GoogleOauthState entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableGoogleOauthStatesID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetGoogleOauthStatesID(*id)
	}
	return uu
}

// SetGoogleOauthStates sets the "google_oauth_states" edge to the GoogleOauthState entity.
func (uu *UserUpdate) SetGoogleOauthStates(g *GoogleOauthState) *UserUpdate {
	return uu.SetGoogleOauthStatesID(g.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (uu *UserUpdate) AddEventIDs(ids ...int) *UserUpdate {
	uu.mutation.AddEventIDs(ids...)
	return uu
}

// AddEvents adds the "events" edges to the Event entity.
func (uu *UserUpdate) AddEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddEventIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGoogleOauthTokens clears the "google_oauth_tokens" edge to the GoogleOauthToken entity.
func (uu *UserUpdate) ClearGoogleOauthTokens() *UserUpdate {
	uu.mutation.ClearGoogleOauthTokens()
	return uu
}

// ClearGoogleOauthStates clears the "google_oauth_states" edge to the GoogleOauthState entity.
func (uu *UserUpdate) ClearGoogleOauthStates() *UserUpdate {
	uu.mutation.ClearGoogleOauthStates()
	return uu
}

// ClearEvents clears all "events" edges to the Event entity.
func (uu *UserUpdate) ClearEvents() *UserUpdate {
	uu.mutation.ClearEvents()
	return uu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (uu *UserUpdate) RemoveEventIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveEventIDs(ids...)
	return uu
}

// RemoveEvents removes "events" edges to Event entities.
func (uu *UserUpdate) RemoveEvents(e ...*Event) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UserID(); ok {
		if err := user.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "User.user_id": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeBytes, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uu.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uu.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uu.mutation.IsAdminVerified(); ok {
		_spec.SetField(user.FieldIsAdminVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.DeleteProtected(); ok {
		_spec.SetField(user.FieldDeleteProtected, field.TypeBool, value)
	}
	if value, ok := uu.mutation.APIKey(); ok {
		_spec.SetField(user.FieldAPIKey, field.TypeString, value)
	}
	if uu.mutation.APIKeyCleared() {
		_spec.ClearField(user.FieldAPIKey, field.TypeString)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if uu.mutation.GoogleOauthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthTokensTable,
			Columns: []string{user.GoogleOauthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthtoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GoogleOauthTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthTokensTable,
			Columns: []string{user.GoogleOauthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthtoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GoogleOauthStatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthStatesTable,
			Columns: []string{user.GoogleOauthStatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthstate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GoogleOauthStatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthStatesTable,
			Columns: []string{user.GoogleOauthStatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthstate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !uu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserID sets the "user_id" field.
func (uuo *UserUpdateOne) SetUserID(s string) *UserUpdateOne {
	uuo.mutation.SetUserID(s)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(b []byte) *UserUpdateOne {
	uuo.mutation.SetPassword(b)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of the "last_name" field.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// SetIsAdminVerified sets the "is_admin_verified" field.
func (uuo *UserUpdateOne) SetIsAdminVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetIsAdminVerified(b)
	return uuo
}

// SetNillableIsAdminVerified sets the "is_admin_verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsAdminVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsAdminVerified(*b)
	}
	return uuo
}

// SetDeleteProtected sets the "delete_protected" field.
func (uuo *UserUpdateOne) SetDeleteProtected(b bool) *UserUpdateOne {
	uuo.mutation.SetDeleteProtected(b)
	return uuo
}

// SetNillableDeleteProtected sets the "delete_protected" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeleteProtected(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetDeleteProtected(*b)
	}
	return uuo
}

// SetAPIKey sets the "api_key" field.
func (uuo *UserUpdateOne) SetAPIKey(s string) *UserUpdateOne {
	uuo.mutation.SetAPIKey(s)
	return uuo
}

// SetNillableAPIKey sets the "api_key" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAPIKey(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAPIKey(*s)
	}
	return uuo
}

// ClearAPIKey clears the value of the "api_key" field.
func (uuo *UserUpdateOne) ClearAPIKey() *UserUpdateOne {
	uuo.mutation.ClearAPIKey()
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetGoogleOauthTokensID sets the "google_oauth_tokens" edge to the GoogleOauthToken entity by ID.
func (uuo *UserUpdateOne) SetGoogleOauthTokensID(id int) *UserUpdateOne {
	uuo.mutation.SetGoogleOauthTokensID(id)
	return uuo
}

// SetNillableGoogleOauthTokensID sets the "google_oauth_tokens" edge to the GoogleOauthToken entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGoogleOauthTokensID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetGoogleOauthTokensID(*id)
	}
	return uuo
}

// SetGoogleOauthTokens sets the "google_oauth_tokens" edge to the GoogleOauthToken entity.
func (uuo *UserUpdateOne) SetGoogleOauthTokens(g *GoogleOauthToken) *UserUpdateOne {
	return uuo.SetGoogleOauthTokensID(g.ID)
}

// SetGoogleOauthStatesID sets the "google_oauth_states" edge to the GoogleOauthState entity by ID.
func (uuo *UserUpdateOne) SetGoogleOauthStatesID(id int) *UserUpdateOne {
	uuo.mutation.SetGoogleOauthStatesID(id)
	return uuo
}

// SetNillableGoogleOauthStatesID sets the "google_oauth_states" edge to the GoogleOauthState entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGoogleOauthStatesID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetGoogleOauthStatesID(*id)
	}
	return uuo
}

// SetGoogleOauthStates sets the "google_oauth_states" edge to the GoogleOauthState entity.
func (uuo *UserUpdateOne) SetGoogleOauthStates(g *GoogleOauthState) *UserUpdateOne {
	return uuo.SetGoogleOauthStatesID(g.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (uuo *UserUpdateOne) AddEventIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddEventIDs(ids...)
	return uuo
}

// AddEvents adds the "events" edges to the Event entity.
func (uuo *UserUpdateOne) AddEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddEventIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGoogleOauthTokens clears the "google_oauth_tokens" edge to the GoogleOauthToken entity.
func (uuo *UserUpdateOne) ClearGoogleOauthTokens() *UserUpdateOne {
	uuo.mutation.ClearGoogleOauthTokens()
	return uuo
}

// ClearGoogleOauthStates clears the "google_oauth_states" edge to the GoogleOauthState entity.
func (uuo *UserUpdateOne) ClearGoogleOauthStates() *UserUpdateOne {
	uuo.mutation.ClearGoogleOauthStates()
	return uuo
}

// ClearEvents clears all "events" edges to the Event entity.
func (uuo *UserUpdateOne) ClearEvents() *UserUpdateOne {
	uuo.mutation.ClearEvents()
	return uuo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (uuo *UserUpdateOne) RemoveEventIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveEventIDs(ids...)
	return uuo
}

// RemoveEvents removes "events" edges to Event entities.
func (uuo *UserUpdateOne) RemoveEvents(e ...*Event) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveEventIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UserID(); ok {
		if err := user.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "User.user_id": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserID(); ok {
		_spec.SetField(user.FieldUserID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeBytes, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uuo.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uuo.mutation.IsAdminVerified(); ok {
		_spec.SetField(user.FieldIsAdminVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.DeleteProtected(); ok {
		_spec.SetField(user.FieldDeleteProtected, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.APIKey(); ok {
		_spec.SetField(user.FieldAPIKey, field.TypeString, value)
	}
	if uuo.mutation.APIKeyCleared() {
		_spec.ClearField(user.FieldAPIKey, field.TypeString)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if uuo.mutation.GoogleOauthTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthTokensTable,
			Columns: []string{user.GoogleOauthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthtoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GoogleOauthTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthTokensTable,
			Columns: []string{user.GoogleOauthTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthtoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GoogleOauthStatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthStatesTable,
			Columns: []string{user.GoogleOauthStatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthstate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GoogleOauthStatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.GoogleOauthStatesTable,
			Columns: []string{user.GoogleOauthStatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: googleoauthstate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !uuo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
