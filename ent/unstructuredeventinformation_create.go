// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/event"
	"github.com/scarlet0725/prism-api/ent/unstructuredeventinformation"
)

// UnStructuredEventInformationCreate is the builder for creating a UnStructuredEventInformation entity.
type UnStructuredEventInformationCreate struct {
	config
	mutation *UnStructuredEventInformationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRyzmuuid sets the "ryzmuuid" field.
func (useic *UnStructuredEventInformationCreate) SetRyzmuuid(s string) *UnStructuredEventInformationCreate {
	useic.mutation.SetRyzmuuid(s)
	return useic
}

// SetVenueName sets the "venue_name" field.
func (useic *UnStructuredEventInformationCreate) SetVenueName(s string) *UnStructuredEventInformationCreate {
	useic.mutation.SetVenueName(s)
	return useic
}

// SetArtistName sets the "artist_name" field.
func (useic *UnStructuredEventInformationCreate) SetArtistName(s string) *UnStructuredEventInformationCreate {
	useic.mutation.SetArtistName(s)
	return useic
}

// SetPrice sets the "price" field.
func (useic *UnStructuredEventInformationCreate) SetPrice(s string) *UnStructuredEventInformationCreate {
	useic.mutation.SetPrice(s)
	return useic
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (useic *UnStructuredEventInformationCreate) SetEventID(id int) *UnStructuredEventInformationCreate {
	useic.mutation.SetEventID(id)
	return useic
}

// SetEvent sets the "event" edge to the Event entity.
func (useic *UnStructuredEventInformationCreate) SetEvent(e *Event) *UnStructuredEventInformationCreate {
	return useic.SetEventID(e.ID)
}

// Mutation returns the UnStructuredEventInformationMutation object of the builder.
func (useic *UnStructuredEventInformationCreate) Mutation() *UnStructuredEventInformationMutation {
	return useic.mutation
}

// Save creates the UnStructuredEventInformation in the database.
func (useic *UnStructuredEventInformationCreate) Save(ctx context.Context) (*UnStructuredEventInformation, error) {
	return withHooks[*UnStructuredEventInformation, UnStructuredEventInformationMutation](ctx, useic.sqlSave, useic.mutation, useic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (useic *UnStructuredEventInformationCreate) SaveX(ctx context.Context) *UnStructuredEventInformation {
	v, err := useic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (useic *UnStructuredEventInformationCreate) Exec(ctx context.Context) error {
	_, err := useic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (useic *UnStructuredEventInformationCreate) ExecX(ctx context.Context) {
	if err := useic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (useic *UnStructuredEventInformationCreate) check() error {
	if _, ok := useic.mutation.Ryzmuuid(); !ok {
		return &ValidationError{Name: "ryzmuuid", err: errors.New(`ent: missing required field "UnStructuredEventInformation.ryzmuuid"`)}
	}
	if _, ok := useic.mutation.VenueName(); !ok {
		return &ValidationError{Name: "venue_name", err: errors.New(`ent: missing required field "UnStructuredEventInformation.venue_name"`)}
	}
	if _, ok := useic.mutation.ArtistName(); !ok {
		return &ValidationError{Name: "artist_name", err: errors.New(`ent: missing required field "UnStructuredEventInformation.artist_name"`)}
	}
	if _, ok := useic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "UnStructuredEventInformation.price"`)}
	}
	if _, ok := useic.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "UnStructuredEventInformation.event"`)}
	}
	return nil
}

func (useic *UnStructuredEventInformationCreate) sqlSave(ctx context.Context) (*UnStructuredEventInformation, error) {
	if err := useic.check(); err != nil {
		return nil, err
	}
	_node, _spec := useic.createSpec()
	if err := sqlgraph.CreateNode(ctx, useic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	useic.mutation.id = &_node.ID
	useic.mutation.done = true
	return _node, nil
}

func (useic *UnStructuredEventInformationCreate) createSpec() (*UnStructuredEventInformation, *sqlgraph.CreateSpec) {
	var (
		_node = &UnStructuredEventInformation{config: useic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: unstructuredeventinformation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unstructuredeventinformation.FieldID,
			},
		}
	)
	_spec.OnConflict = useic.conflict
	if value, ok := useic.mutation.Ryzmuuid(); ok {
		_spec.SetField(unstructuredeventinformation.FieldRyzmuuid, field.TypeString, value)
		_node.Ryzmuuid = value
	}
	if value, ok := useic.mutation.VenueName(); ok {
		_spec.SetField(unstructuredeventinformation.FieldVenueName, field.TypeString, value)
		_node.VenueName = value
	}
	if value, ok := useic.mutation.ArtistName(); ok {
		_spec.SetField(unstructuredeventinformation.FieldArtistName, field.TypeString, value)
		_node.ArtistName = value
	}
	if value, ok := useic.mutation.Price(); ok {
		_spec.SetField(unstructuredeventinformation.FieldPrice, field.TypeString, value)
		_node.Price = value
	}
	if nodes := useic.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unstructuredeventinformation.EventTable,
			Columns: []string{unstructuredeventinformation.EventColumn},
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
		_node.event_un_structured_event_informations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UnStructuredEventInformation.Create().
//		SetRyzmuuid(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UnStructuredEventInformationUpsert) {
//			SetRyzmuuid(v+v).
//		}).
//		Exec(ctx)
func (useic *UnStructuredEventInformationCreate) OnConflict(opts ...sql.ConflictOption) *UnStructuredEventInformationUpsertOne {
	useic.conflict = opts
	return &UnStructuredEventInformationUpsertOne{
		create: useic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (useic *UnStructuredEventInformationCreate) OnConflictColumns(columns ...string) *UnStructuredEventInformationUpsertOne {
	useic.conflict = append(useic.conflict, sql.ConflictColumns(columns...))
	return &UnStructuredEventInformationUpsertOne{
		create: useic,
	}
}

type (
	// UnStructuredEventInformationUpsertOne is the builder for "upsert"-ing
	//  one UnStructuredEventInformation node.
	UnStructuredEventInformationUpsertOne struct {
		create *UnStructuredEventInformationCreate
	}

	// UnStructuredEventInformationUpsert is the "OnConflict" setter.
	UnStructuredEventInformationUpsert struct {
		*sql.UpdateSet
	}
)

// SetRyzmuuid sets the "ryzmuuid" field.
func (u *UnStructuredEventInformationUpsert) SetRyzmuuid(v string) *UnStructuredEventInformationUpsert {
	u.Set(unstructuredeventinformation.FieldRyzmuuid, v)
	return u
}

// UpdateRyzmuuid sets the "ryzmuuid" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsert) UpdateRyzmuuid() *UnStructuredEventInformationUpsert {
	u.SetExcluded(unstructuredeventinformation.FieldRyzmuuid)
	return u
}

// SetVenueName sets the "venue_name" field.
func (u *UnStructuredEventInformationUpsert) SetVenueName(v string) *UnStructuredEventInformationUpsert {
	u.Set(unstructuredeventinformation.FieldVenueName, v)
	return u
}

// UpdateVenueName sets the "venue_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsert) UpdateVenueName() *UnStructuredEventInformationUpsert {
	u.SetExcluded(unstructuredeventinformation.FieldVenueName)
	return u
}

// SetArtistName sets the "artist_name" field.
func (u *UnStructuredEventInformationUpsert) SetArtistName(v string) *UnStructuredEventInformationUpsert {
	u.Set(unstructuredeventinformation.FieldArtistName, v)
	return u
}

// UpdateArtistName sets the "artist_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsert) UpdateArtistName() *UnStructuredEventInformationUpsert {
	u.SetExcluded(unstructuredeventinformation.FieldArtistName)
	return u
}

// SetPrice sets the "price" field.
func (u *UnStructuredEventInformationUpsert) SetPrice(v string) *UnStructuredEventInformationUpsert {
	u.Set(unstructuredeventinformation.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsert) UpdatePrice() *UnStructuredEventInformationUpsert {
	u.SetExcluded(unstructuredeventinformation.FieldPrice)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UnStructuredEventInformationUpsertOne) UpdateNewValues() *UnStructuredEventInformationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UnStructuredEventInformationUpsertOne) Ignore() *UnStructuredEventInformationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UnStructuredEventInformationUpsertOne) DoNothing() *UnStructuredEventInformationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UnStructuredEventInformationCreate.OnConflict
// documentation for more info.
func (u *UnStructuredEventInformationUpsertOne) Update(set func(*UnStructuredEventInformationUpsert)) *UnStructuredEventInformationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UnStructuredEventInformationUpsert{UpdateSet: update})
	}))
	return u
}

// SetRyzmuuid sets the "ryzmuuid" field.
func (u *UnStructuredEventInformationUpsertOne) SetRyzmuuid(v string) *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetRyzmuuid(v)
	})
}

// UpdateRyzmuuid sets the "ryzmuuid" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertOne) UpdateRyzmuuid() *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateRyzmuuid()
	})
}

// SetVenueName sets the "venue_name" field.
func (u *UnStructuredEventInformationUpsertOne) SetVenueName(v string) *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetVenueName(v)
	})
}

// UpdateVenueName sets the "venue_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertOne) UpdateVenueName() *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateVenueName()
	})
}

// SetArtistName sets the "artist_name" field.
func (u *UnStructuredEventInformationUpsertOne) SetArtistName(v string) *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetArtistName(v)
	})
}

// UpdateArtistName sets the "artist_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertOne) UpdateArtistName() *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateArtistName()
	})
}

// SetPrice sets the "price" field.
func (u *UnStructuredEventInformationUpsertOne) SetPrice(v string) *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertOne) UpdatePrice() *UnStructuredEventInformationUpsertOne {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *UnStructuredEventInformationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UnStructuredEventInformationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UnStructuredEventInformationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UnStructuredEventInformationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UnStructuredEventInformationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UnStructuredEventInformationCreateBulk is the builder for creating many UnStructuredEventInformation entities in bulk.
type UnStructuredEventInformationCreateBulk struct {
	config
	builders []*UnStructuredEventInformationCreate
	conflict []sql.ConflictOption
}

// Save creates the UnStructuredEventInformation entities in the database.
func (useicb *UnStructuredEventInformationCreateBulk) Save(ctx context.Context) ([]*UnStructuredEventInformation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(useicb.builders))
	nodes := make([]*UnStructuredEventInformation, len(useicb.builders))
	mutators := make([]Mutator, len(useicb.builders))
	for i := range useicb.builders {
		func(i int, root context.Context) {
			builder := useicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UnStructuredEventInformationMutation)
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
					_, err = mutators[i+1].Mutate(root, useicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = useicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, useicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, useicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (useicb *UnStructuredEventInformationCreateBulk) SaveX(ctx context.Context) []*UnStructuredEventInformation {
	v, err := useicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (useicb *UnStructuredEventInformationCreateBulk) Exec(ctx context.Context) error {
	_, err := useicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (useicb *UnStructuredEventInformationCreateBulk) ExecX(ctx context.Context) {
	if err := useicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UnStructuredEventInformation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UnStructuredEventInformationUpsert) {
//			SetRyzmuuid(v+v).
//		}).
//		Exec(ctx)
func (useicb *UnStructuredEventInformationCreateBulk) OnConflict(opts ...sql.ConflictOption) *UnStructuredEventInformationUpsertBulk {
	useicb.conflict = opts
	return &UnStructuredEventInformationUpsertBulk{
		create: useicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (useicb *UnStructuredEventInformationCreateBulk) OnConflictColumns(columns ...string) *UnStructuredEventInformationUpsertBulk {
	useicb.conflict = append(useicb.conflict, sql.ConflictColumns(columns...))
	return &UnStructuredEventInformationUpsertBulk{
		create: useicb,
	}
}

// UnStructuredEventInformationUpsertBulk is the builder for "upsert"-ing
// a bulk of UnStructuredEventInformation nodes.
type UnStructuredEventInformationUpsertBulk struct {
	create *UnStructuredEventInformationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UnStructuredEventInformationUpsertBulk) UpdateNewValues() *UnStructuredEventInformationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UnStructuredEventInformation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UnStructuredEventInformationUpsertBulk) Ignore() *UnStructuredEventInformationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UnStructuredEventInformationUpsertBulk) DoNothing() *UnStructuredEventInformationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UnStructuredEventInformationCreateBulk.OnConflict
// documentation for more info.
func (u *UnStructuredEventInformationUpsertBulk) Update(set func(*UnStructuredEventInformationUpsert)) *UnStructuredEventInformationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UnStructuredEventInformationUpsert{UpdateSet: update})
	}))
	return u
}

// SetRyzmuuid sets the "ryzmuuid" field.
func (u *UnStructuredEventInformationUpsertBulk) SetRyzmuuid(v string) *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetRyzmuuid(v)
	})
}

// UpdateRyzmuuid sets the "ryzmuuid" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertBulk) UpdateRyzmuuid() *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateRyzmuuid()
	})
}

// SetVenueName sets the "venue_name" field.
func (u *UnStructuredEventInformationUpsertBulk) SetVenueName(v string) *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetVenueName(v)
	})
}

// UpdateVenueName sets the "venue_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertBulk) UpdateVenueName() *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateVenueName()
	})
}

// SetArtistName sets the "artist_name" field.
func (u *UnStructuredEventInformationUpsertBulk) SetArtistName(v string) *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetArtistName(v)
	})
}

// UpdateArtistName sets the "artist_name" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertBulk) UpdateArtistName() *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdateArtistName()
	})
}

// SetPrice sets the "price" field.
func (u *UnStructuredEventInformationUpsertBulk) SetPrice(v string) *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.SetPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *UnStructuredEventInformationUpsertBulk) UpdatePrice() *UnStructuredEventInformationUpsertBulk {
	return u.Update(func(s *UnStructuredEventInformationUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *UnStructuredEventInformationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UnStructuredEventInformationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UnStructuredEventInformationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UnStructuredEventInformationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}