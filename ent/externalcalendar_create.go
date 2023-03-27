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
	"github.com/scarlet0725/prism-api/ent/externalcalendar"
	"github.com/scarlet0725/prism-api/ent/user"
)

// ExternalCalendarCreate is the builder for creating a ExternalCalendar entity.
type ExternalCalendarCreate struct {
	config
	mutation *ExternalCalendarMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ecc *ExternalCalendarCreate) SetName(s string) *ExternalCalendarCreate {
	ecc.mutation.SetName(s)
	return ecc
}

// SetDescription sets the "description" field.
func (ecc *ExternalCalendarCreate) SetDescription(s string) *ExternalCalendarCreate {
	ecc.mutation.SetDescription(s)
	return ecc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ecc *ExternalCalendarCreate) SetNillableDescription(s *string) *ExternalCalendarCreate {
	if s != nil {
		ecc.SetDescription(*s)
	}
	return ecc
}

// SetCalendarID sets the "calendar_id" field.
func (ecc *ExternalCalendarCreate) SetCalendarID(s string) *ExternalCalendarCreate {
	ecc.mutation.SetCalendarID(s)
	return ecc
}

// SetSourceType sets the "source_type" field.
func (ecc *ExternalCalendarCreate) SetSourceType(s string) *ExternalCalendarCreate {
	ecc.mutation.SetSourceType(s)
	return ecc
}

// SetCreatedAt sets the "created_at" field.
func (ecc *ExternalCalendarCreate) SetCreatedAt(t time.Time) *ExternalCalendarCreate {
	ecc.mutation.SetCreatedAt(t)
	return ecc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ecc *ExternalCalendarCreate) SetNillableCreatedAt(t *time.Time) *ExternalCalendarCreate {
	if t != nil {
		ecc.SetCreatedAt(*t)
	}
	return ecc
}

// SetUpdatedAt sets the "updated_at" field.
func (ecc *ExternalCalendarCreate) SetUpdatedAt(t time.Time) *ExternalCalendarCreate {
	ecc.mutation.SetUpdatedAt(t)
	return ecc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ecc *ExternalCalendarCreate) SetNillableUpdatedAt(t *time.Time) *ExternalCalendarCreate {
	if t != nil {
		ecc.SetUpdatedAt(*t)
	}
	return ecc
}

// SetDeletedAt sets the "deleted_at" field.
func (ecc *ExternalCalendarCreate) SetDeletedAt(t time.Time) *ExternalCalendarCreate {
	ecc.mutation.SetDeletedAt(t)
	return ecc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ecc *ExternalCalendarCreate) SetNillableDeletedAt(t *time.Time) *ExternalCalendarCreate {
	if t != nil {
		ecc.SetDeletedAt(*t)
	}
	return ecc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ecc *ExternalCalendarCreate) SetUserID(id int) *ExternalCalendarCreate {
	ecc.mutation.SetUserID(id)
	return ecc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ecc *ExternalCalendarCreate) SetNillableUserID(id *int) *ExternalCalendarCreate {
	if id != nil {
		ecc = ecc.SetUserID(*id)
	}
	return ecc
}

// SetUser sets the "user" edge to the User entity.
func (ecc *ExternalCalendarCreate) SetUser(u *User) *ExternalCalendarCreate {
	return ecc.SetUserID(u.ID)
}

// Mutation returns the ExternalCalendarMutation object of the builder.
func (ecc *ExternalCalendarCreate) Mutation() *ExternalCalendarMutation {
	return ecc.mutation
}

// Save creates the ExternalCalendar in the database.
func (ecc *ExternalCalendarCreate) Save(ctx context.Context) (*ExternalCalendar, error) {
	ecc.defaults()
	return withHooks[*ExternalCalendar, ExternalCalendarMutation](ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *ExternalCalendarCreate) SaveX(ctx context.Context) *ExternalCalendar {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *ExternalCalendarCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *ExternalCalendarCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ecc *ExternalCalendarCreate) defaults() {
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		v := externalcalendar.DefaultCreatedAt()
		ecc.mutation.SetCreatedAt(v)
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		v := externalcalendar.DefaultUpdatedAt()
		ecc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *ExternalCalendarCreate) check() error {
	if _, ok := ecc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ExternalCalendar.name"`)}
	}
	if v, ok := ecc.mutation.Name(); ok {
		if err := externalcalendar.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ExternalCalendar.name": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.CalendarID(); !ok {
		return &ValidationError{Name: "calendar_id", err: errors.New(`ent: missing required field "ExternalCalendar.calendar_id"`)}
	}
	if v, ok := ecc.mutation.CalendarID(); ok {
		if err := externalcalendar.CalendarIDValidator(v); err != nil {
			return &ValidationError{Name: "calendar_id", err: fmt.Errorf(`ent: validator failed for field "ExternalCalendar.calendar_id": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.SourceType(); !ok {
		return &ValidationError{Name: "source_type", err: errors.New(`ent: missing required field "ExternalCalendar.source_type"`)}
	}
	if v, ok := ecc.mutation.SourceType(); ok {
		if err := externalcalendar.SourceTypeValidator(v); err != nil {
			return &ValidationError{Name: "source_type", err: fmt.Errorf(`ent: validator failed for field "ExternalCalendar.source_type": %w`, err)}
		}
	}
	if _, ok := ecc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ExternalCalendar.created_at"`)}
	}
	if _, ok := ecc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ExternalCalendar.updated_at"`)}
	}
	return nil
}

func (ecc *ExternalCalendarCreate) sqlSave(ctx context.Context) (*ExternalCalendar, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *ExternalCalendarCreate) createSpec() (*ExternalCalendar, *sqlgraph.CreateSpec) {
	var (
		_node = &ExternalCalendar{config: ecc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: externalcalendar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: externalcalendar.FieldID,
			},
		}
	)
	_spec.OnConflict = ecc.conflict
	if value, ok := ecc.mutation.Name(); ok {
		_spec.SetField(externalcalendar.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ecc.mutation.Description(); ok {
		_spec.SetField(externalcalendar.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ecc.mutation.CalendarID(); ok {
		_spec.SetField(externalcalendar.FieldCalendarID, field.TypeString, value)
		_node.CalendarID = value
	}
	if value, ok := ecc.mutation.SourceType(); ok {
		_spec.SetField(externalcalendar.FieldSourceType, field.TypeString, value)
		_node.SourceType = value
	}
	if value, ok := ecc.mutation.CreatedAt(); ok {
		_spec.SetField(externalcalendar.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ecc.mutation.UpdatedAt(); ok {
		_spec.SetField(externalcalendar.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ecc.mutation.DeletedAt(); ok {
		_spec.SetField(externalcalendar.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := ecc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   externalcalendar.UserTable,
			Columns: []string{externalcalendar.UserColumn},
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExternalCalendar.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExternalCalendarUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ecc *ExternalCalendarCreate) OnConflict(opts ...sql.ConflictOption) *ExternalCalendarUpsertOne {
	ecc.conflict = opts
	return &ExternalCalendarUpsertOne{
		create: ecc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecc *ExternalCalendarCreate) OnConflictColumns(columns ...string) *ExternalCalendarUpsertOne {
	ecc.conflict = append(ecc.conflict, sql.ConflictColumns(columns...))
	return &ExternalCalendarUpsertOne{
		create: ecc,
	}
}

type (
	// ExternalCalendarUpsertOne is the builder for "upsert"-ing
	//  one ExternalCalendar node.
	ExternalCalendarUpsertOne struct {
		create *ExternalCalendarCreate
	}

	// ExternalCalendarUpsert is the "OnConflict" setter.
	ExternalCalendarUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ExternalCalendarUpsert) SetName(v string) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateName() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ExternalCalendarUpsert) SetDescription(v string) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateDescription() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *ExternalCalendarUpsert) ClearDescription() *ExternalCalendarUpsert {
	u.SetNull(externalcalendar.FieldDescription)
	return u
}

// SetCalendarID sets the "calendar_id" field.
func (u *ExternalCalendarUpsert) SetCalendarID(v string) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldCalendarID, v)
	return u
}

// UpdateCalendarID sets the "calendar_id" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateCalendarID() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldCalendarID)
	return u
}

// SetSourceType sets the "source_type" field.
func (u *ExternalCalendarUpsert) SetSourceType(v string) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldSourceType, v)
	return u
}

// UpdateSourceType sets the "source_type" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateSourceType() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldSourceType)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExternalCalendarUpsert) SetUpdatedAt(v time.Time) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateUpdatedAt() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExternalCalendarUpsert) SetDeletedAt(v time.Time) *ExternalCalendarUpsert {
	u.Set(externalcalendar.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsert) UpdateDeletedAt() *ExternalCalendarUpsert {
	u.SetExcluded(externalcalendar.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ExternalCalendarUpsert) ClearDeletedAt() *ExternalCalendarUpsert {
	u.SetNull(externalcalendar.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ExternalCalendarUpsertOne) UpdateNewValues() *ExternalCalendarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(externalcalendar.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ExternalCalendarUpsertOne) Ignore() *ExternalCalendarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExternalCalendarUpsertOne) DoNothing() *ExternalCalendarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExternalCalendarCreate.OnConflict
// documentation for more info.
func (u *ExternalCalendarUpsertOne) Update(set func(*ExternalCalendarUpsert)) *ExternalCalendarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExternalCalendarUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ExternalCalendarUpsertOne) SetName(v string) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateName() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ExternalCalendarUpsertOne) SetDescription(v string) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateDescription() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ExternalCalendarUpsertOne) ClearDescription() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.ClearDescription()
	})
}

// SetCalendarID sets the "calendar_id" field.
func (u *ExternalCalendarUpsertOne) SetCalendarID(v string) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetCalendarID(v)
	})
}

// UpdateCalendarID sets the "calendar_id" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateCalendarID() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateCalendarID()
	})
}

// SetSourceType sets the "source_type" field.
func (u *ExternalCalendarUpsertOne) SetSourceType(v string) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetSourceType(v)
	})
}

// UpdateSourceType sets the "source_type" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateSourceType() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateSourceType()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExternalCalendarUpsertOne) SetUpdatedAt(v time.Time) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateUpdatedAt() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExternalCalendarUpsertOne) SetDeletedAt(v time.Time) *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsertOne) UpdateDeletedAt() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ExternalCalendarUpsertOne) ClearDeletedAt() *ExternalCalendarUpsertOne {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ExternalCalendarUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExternalCalendarCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExternalCalendarUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ExternalCalendarUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ExternalCalendarUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ExternalCalendarCreateBulk is the builder for creating many ExternalCalendar entities in bulk.
type ExternalCalendarCreateBulk struct {
	config
	builders []*ExternalCalendarCreate
	conflict []sql.ConflictOption
}

// Save creates the ExternalCalendar entities in the database.
func (eccb *ExternalCalendarCreateBulk) Save(ctx context.Context) ([]*ExternalCalendar, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*ExternalCalendar, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExternalCalendarMutation)
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
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *ExternalCalendarCreateBulk) SaveX(ctx context.Context) []*ExternalCalendar {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *ExternalCalendarCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *ExternalCalendarCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExternalCalendar.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExternalCalendarUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (eccb *ExternalCalendarCreateBulk) OnConflict(opts ...sql.ConflictOption) *ExternalCalendarUpsertBulk {
	eccb.conflict = opts
	return &ExternalCalendarUpsertBulk{
		create: eccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eccb *ExternalCalendarCreateBulk) OnConflictColumns(columns ...string) *ExternalCalendarUpsertBulk {
	eccb.conflict = append(eccb.conflict, sql.ConflictColumns(columns...))
	return &ExternalCalendarUpsertBulk{
		create: eccb,
	}
}

// ExternalCalendarUpsertBulk is the builder for "upsert"-ing
// a bulk of ExternalCalendar nodes.
type ExternalCalendarUpsertBulk struct {
	create *ExternalCalendarCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ExternalCalendarUpsertBulk) UpdateNewValues() *ExternalCalendarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(externalcalendar.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ExternalCalendar.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ExternalCalendarUpsertBulk) Ignore() *ExternalCalendarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExternalCalendarUpsertBulk) DoNothing() *ExternalCalendarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExternalCalendarCreateBulk.OnConflict
// documentation for more info.
func (u *ExternalCalendarUpsertBulk) Update(set func(*ExternalCalendarUpsert)) *ExternalCalendarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExternalCalendarUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ExternalCalendarUpsertBulk) SetName(v string) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateName() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ExternalCalendarUpsertBulk) SetDescription(v string) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateDescription() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ExternalCalendarUpsertBulk) ClearDescription() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.ClearDescription()
	})
}

// SetCalendarID sets the "calendar_id" field.
func (u *ExternalCalendarUpsertBulk) SetCalendarID(v string) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetCalendarID(v)
	})
}

// UpdateCalendarID sets the "calendar_id" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateCalendarID() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateCalendarID()
	})
}

// SetSourceType sets the "source_type" field.
func (u *ExternalCalendarUpsertBulk) SetSourceType(v string) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetSourceType(v)
	})
}

// UpdateSourceType sets the "source_type" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateSourceType() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateSourceType()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExternalCalendarUpsertBulk) SetUpdatedAt(v time.Time) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateUpdatedAt() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExternalCalendarUpsertBulk) SetDeletedAt(v time.Time) *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExternalCalendarUpsertBulk) UpdateDeletedAt() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ExternalCalendarUpsertBulk) ClearDeletedAt() *ExternalCalendarUpsertBulk {
	return u.Update(func(s *ExternalCalendarUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *ExternalCalendarUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ExternalCalendarCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExternalCalendarCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExternalCalendarUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}