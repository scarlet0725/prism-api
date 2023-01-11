// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/scarlet0725/prism-api/ent/googleoauthtoken"
	"github.com/scarlet0725/prism-api/ent/predicate"
	"github.com/scarlet0725/prism-api/ent/user"
)

// GoogleOauthTokenQuery is the builder for querying GoogleOauthToken entities.
type GoogleOauthTokenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.GoogleOauthToken
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoogleOauthTokenQuery builder.
func (gotq *GoogleOauthTokenQuery) Where(ps ...predicate.GoogleOauthToken) *GoogleOauthTokenQuery {
	gotq.predicates = append(gotq.predicates, ps...)
	return gotq
}

// Limit the number of records to be returned by this query.
func (gotq *GoogleOauthTokenQuery) Limit(limit int) *GoogleOauthTokenQuery {
	gotq.limit = &limit
	return gotq
}

// Offset to start from.
func (gotq *GoogleOauthTokenQuery) Offset(offset int) *GoogleOauthTokenQuery {
	gotq.offset = &offset
	return gotq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gotq *GoogleOauthTokenQuery) Unique(unique bool) *GoogleOauthTokenQuery {
	gotq.unique = &unique
	return gotq
}

// Order specifies how the records should be ordered.
func (gotq *GoogleOauthTokenQuery) Order(o ...OrderFunc) *GoogleOauthTokenQuery {
	gotq.order = append(gotq.order, o...)
	return gotq
}

// QueryUser chains the current query on the "user" edge.
func (gotq *GoogleOauthTokenQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: gotq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gotq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(googleoauthtoken.Table, googleoauthtoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, googleoauthtoken.UserTable, googleoauthtoken.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(gotq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GoogleOauthToken entity from the query.
// Returns a *NotFoundError when no GoogleOauthToken was found.
func (gotq *GoogleOauthTokenQuery) First(ctx context.Context) (*GoogleOauthToken, error) {
	nodes, err := gotq.Limit(1).All(newQueryContext(ctx, TypeGoogleOauthToken, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{googleoauthtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) FirstX(ctx context.Context) *GoogleOauthToken {
	node, err := gotq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoogleOauthToken ID from the query.
// Returns a *NotFoundError when no GoogleOauthToken ID was found.
func (gotq *GoogleOauthTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gotq.Limit(1).IDs(newQueryContext(ctx, TypeGoogleOauthToken, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{googleoauthtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := gotq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoogleOauthToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoogleOauthToken entity is found.
// Returns a *NotFoundError when no GoogleOauthToken entities are found.
func (gotq *GoogleOauthTokenQuery) Only(ctx context.Context) (*GoogleOauthToken, error) {
	nodes, err := gotq.Limit(2).All(newQueryContext(ctx, TypeGoogleOauthToken, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{googleoauthtoken.Label}
	default:
		return nil, &NotSingularError{googleoauthtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) OnlyX(ctx context.Context) *GoogleOauthToken {
	node, err := gotq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoogleOauthToken ID in the query.
// Returns a *NotSingularError when more than one GoogleOauthToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (gotq *GoogleOauthTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gotq.Limit(2).IDs(newQueryContext(ctx, TypeGoogleOauthToken, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{googleoauthtoken.Label}
	default:
		err = &NotSingularError{googleoauthtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := gotq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoogleOauthTokens.
func (gotq *GoogleOauthTokenQuery) All(ctx context.Context) ([]*GoogleOauthToken, error) {
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "All")
	if err := gotq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GoogleOauthToken, *GoogleOauthTokenQuery]()
	return withInterceptors[[]*GoogleOauthToken](ctx, gotq, qr, gotq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) AllX(ctx context.Context) []*GoogleOauthToken {
	nodes, err := gotq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoogleOauthToken IDs.
func (gotq *GoogleOauthTokenQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "IDs")
	if err := gotq.Select(googleoauthtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := gotq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gotq *GoogleOauthTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "Count")
	if err := gotq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gotq, querierCount[*GoogleOauthTokenQuery](), gotq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) CountX(ctx context.Context) int {
	count, err := gotq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gotq *GoogleOauthTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "Exist")
	switch _, err := gotq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gotq *GoogleOauthTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := gotq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoogleOauthTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gotq *GoogleOauthTokenQuery) Clone() *GoogleOauthTokenQuery {
	if gotq == nil {
		return nil
	}
	return &GoogleOauthTokenQuery{
		config:     gotq.config,
		limit:      gotq.limit,
		offset:     gotq.offset,
		order:      append([]OrderFunc{}, gotq.order...),
		inters:     append([]Interceptor{}, gotq.inters...),
		predicates: append([]predicate.GoogleOauthToken{}, gotq.predicates...),
		withUser:   gotq.withUser.Clone(),
		// clone intermediate query.
		sql:    gotq.sql.Clone(),
		path:   gotq.path,
		unique: gotq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (gotq *GoogleOauthTokenQuery) WithUser(opts ...func(*UserQuery)) *GoogleOauthTokenQuery {
	query := (&UserClient{config: gotq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gotq.withUser = query
	return gotq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RefreshToken string `json:"refresh_token,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoogleOauthToken.Query().
//		GroupBy(googleoauthtoken.FieldRefreshToken).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gotq *GoogleOauthTokenQuery) GroupBy(field string, fields ...string) *GoogleOauthTokenGroupBy {
	gotq.fields = append([]string{field}, fields...)
	grbuild := &GoogleOauthTokenGroupBy{build: gotq}
	grbuild.flds = &gotq.fields
	grbuild.label = googleoauthtoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RefreshToken string `json:"refresh_token,omitempty"`
//	}
//
//	client.GoogleOauthToken.Query().
//		Select(googleoauthtoken.FieldRefreshToken).
//		Scan(ctx, &v)
func (gotq *GoogleOauthTokenQuery) Select(fields ...string) *GoogleOauthTokenSelect {
	gotq.fields = append(gotq.fields, fields...)
	sbuild := &GoogleOauthTokenSelect{GoogleOauthTokenQuery: gotq}
	sbuild.label = googleoauthtoken.Label
	sbuild.flds, sbuild.scan = &gotq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoogleOauthTokenSelect configured with the given aggregations.
func (gotq *GoogleOauthTokenQuery) Aggregate(fns ...AggregateFunc) *GoogleOauthTokenSelect {
	return gotq.Select().Aggregate(fns...)
}

func (gotq *GoogleOauthTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gotq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gotq); err != nil {
				return err
			}
		}
	}
	for _, f := range gotq.fields {
		if !googleoauthtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gotq.path != nil {
		prev, err := gotq.path(ctx)
		if err != nil {
			return err
		}
		gotq.sql = prev
	}
	return nil
}

func (gotq *GoogleOauthTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoogleOauthToken, error) {
	var (
		nodes       = []*GoogleOauthToken{}
		withFKs     = gotq.withFKs
		_spec       = gotq.querySpec()
		loadedTypes = [1]bool{
			gotq.withUser != nil,
		}
	)
	if gotq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, googleoauthtoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GoogleOauthToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GoogleOauthToken{config: gotq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gotq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gotq.withUser; query != nil {
		if err := gotq.loadUser(ctx, query, nodes, nil,
			func(n *GoogleOauthToken, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gotq *GoogleOauthTokenQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GoogleOauthToken, init func(*GoogleOauthToken), assign func(*GoogleOauthToken, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GoogleOauthToken)
	for i := range nodes {
		if nodes[i].user_id == nil {
			continue
		}
		fk := *nodes[i].user_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gotq *GoogleOauthTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gotq.querySpec()
	_spec.Node.Columns = gotq.fields
	if len(gotq.fields) > 0 {
		_spec.Unique = gotq.unique != nil && *gotq.unique
	}
	return sqlgraph.CountNodes(ctx, gotq.driver, _spec)
}

func (gotq *GoogleOauthTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   googleoauthtoken.Table,
			Columns: googleoauthtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: googleoauthtoken.FieldID,
			},
		},
		From:   gotq.sql,
		Unique: true,
	}
	if unique := gotq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gotq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, googleoauthtoken.FieldID)
		for i := range fields {
			if fields[i] != googleoauthtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gotq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gotq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gotq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gotq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gotq *GoogleOauthTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gotq.driver.Dialect())
	t1 := builder.Table(googleoauthtoken.Table)
	columns := gotq.fields
	if len(columns) == 0 {
		columns = googleoauthtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gotq.sql != nil {
		selector = gotq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gotq.unique != nil && *gotq.unique {
		selector.Distinct()
	}
	for _, p := range gotq.predicates {
		p(selector)
	}
	for _, p := range gotq.order {
		p(selector)
	}
	if offset := gotq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gotq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoogleOauthTokenGroupBy is the group-by builder for GoogleOauthToken entities.
type GoogleOauthTokenGroupBy struct {
	selector
	build *GoogleOauthTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gotgb *GoogleOauthTokenGroupBy) Aggregate(fns ...AggregateFunc) *GoogleOauthTokenGroupBy {
	gotgb.fns = append(gotgb.fns, fns...)
	return gotgb
}

// Scan applies the selector query and scans the result into the given value.
func (gotgb *GoogleOauthTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "GroupBy")
	if err := gotgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoogleOauthTokenQuery, *GoogleOauthTokenGroupBy](ctx, gotgb.build, gotgb, gotgb.build.inters, v)
}

func (gotgb *GoogleOauthTokenGroupBy) sqlScan(ctx context.Context, root *GoogleOauthTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gotgb.fns))
	for _, fn := range gotgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gotgb.flds)+len(gotgb.fns))
		for _, f := range *gotgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gotgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gotgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoogleOauthTokenSelect is the builder for selecting fields of GoogleOauthToken entities.
type GoogleOauthTokenSelect struct {
	*GoogleOauthTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gots *GoogleOauthTokenSelect) Aggregate(fns ...AggregateFunc) *GoogleOauthTokenSelect {
	gots.fns = append(gots.fns, fns...)
	return gots
}

// Scan applies the selector query and scans the result into the given value.
func (gots *GoogleOauthTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGoogleOauthToken, "Select")
	if err := gots.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoogleOauthTokenQuery, *GoogleOauthTokenSelect](ctx, gots.GoogleOauthTokenQuery, gots, gots.inters, v)
}

func (gots *GoogleOauthTokenSelect) sqlScan(ctx context.Context, root *GoogleOauthTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gots.fns))
	for _, fn := range gots.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gots.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
