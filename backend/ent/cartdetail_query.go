// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"selfregi/ent/cart"
	"selfregi/ent/cartdetail"
	"selfregi/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartDetailQuery is the builder for querying CartDetail entities.
type CartDetailQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CartDetail
	withCart   *CartQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CartDetailQuery builder.
func (cdq *CartDetailQuery) Where(ps ...predicate.CartDetail) *CartDetailQuery {
	cdq.predicates = append(cdq.predicates, ps...)
	return cdq
}

// Limit adds a limit step to the query.
func (cdq *CartDetailQuery) Limit(limit int) *CartDetailQuery {
	cdq.limit = &limit
	return cdq
}

// Offset adds an offset step to the query.
func (cdq *CartDetailQuery) Offset(offset int) *CartDetailQuery {
	cdq.offset = &offset
	return cdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cdq *CartDetailQuery) Unique(unique bool) *CartDetailQuery {
	cdq.unique = &unique
	return cdq
}

// Order adds an order step to the query.
func (cdq *CartDetailQuery) Order(o ...OrderFunc) *CartDetailQuery {
	cdq.order = append(cdq.order, o...)
	return cdq
}

// QueryCart chains the current query on the "cart" edge.
func (cdq *CartDetailQuery) QueryCart() *CartQuery {
	query := &CartQuery{config: cdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cartdetail.Table, cartdetail.FieldID, selector),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartdetail.CartTable, cartdetail.CartColumn),
		)
		fromU = sqlgraph.SetNeighbors(cdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CartDetail entity from the query.
// Returns a *NotFoundError when no CartDetail was found.
func (cdq *CartDetailQuery) First(ctx context.Context) (*CartDetail, error) {
	nodes, err := cdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cartdetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cdq *CartDetailQuery) FirstX(ctx context.Context) *CartDetail {
	node, err := cdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CartDetail ID from the query.
// Returns a *NotFoundError when no CartDetail ID was found.
func (cdq *CartDetailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cartdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cdq *CartDetailQuery) FirstIDX(ctx context.Context) int {
	id, err := cdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CartDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CartDetail entity is found.
// Returns a *NotFoundError when no CartDetail entities are found.
func (cdq *CartDetailQuery) Only(ctx context.Context) (*CartDetail, error) {
	nodes, err := cdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cartdetail.Label}
	default:
		return nil, &NotSingularError{cartdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cdq *CartDetailQuery) OnlyX(ctx context.Context) *CartDetail {
	node, err := cdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CartDetail ID in the query.
// Returns a *NotSingularError when more than one CartDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (cdq *CartDetailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cartdetail.Label}
	default:
		err = &NotSingularError{cartdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cdq *CartDetailQuery) OnlyIDX(ctx context.Context) int {
	id, err := cdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CartDetails.
func (cdq *CartDetailQuery) All(ctx context.Context) ([]*CartDetail, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cdq *CartDetailQuery) AllX(ctx context.Context) []*CartDetail {
	nodes, err := cdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CartDetail IDs.
func (cdq *CartDetailQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cdq.Select(cartdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cdq *CartDetailQuery) IDsX(ctx context.Context) []int {
	ids, err := cdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cdq *CartDetailQuery) Count(ctx context.Context) (int, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cdq *CartDetailQuery) CountX(ctx context.Context) int {
	count, err := cdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cdq *CartDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cdq *CartDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := cdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CartDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cdq *CartDetailQuery) Clone() *CartDetailQuery {
	if cdq == nil {
		return nil
	}
	return &CartDetailQuery{
		config:     cdq.config,
		limit:      cdq.limit,
		offset:     cdq.offset,
		order:      append([]OrderFunc{}, cdq.order...),
		predicates: append([]predicate.CartDetail{}, cdq.predicates...),
		withCart:   cdq.withCart.Clone(),
		// clone intermediate query.
		sql:    cdq.sql.Clone(),
		path:   cdq.path,
		unique: cdq.unique,
	}
}

// WithCart tells the query-builder to eager-load the nodes that are connected to
// the "cart" edge. The optional arguments are used to configure the query builder of the edge.
func (cdq *CartDetailQuery) WithCart(opts ...func(*CartQuery)) *CartDetailQuery {
	query := &CartQuery{config: cdq.config}
	for _, opt := range opts {
		opt(query)
	}
	cdq.withCart = query
	return cdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CartDetail.Query().
//		GroupBy(cartdetail.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cdq *CartDetailQuery) GroupBy(field string, fields ...string) *CartDetailGroupBy {
	grbuild := &CartDetailGroupBy{config: cdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cdq.sqlQuery(ctx), nil
	}
	grbuild.label = cartdetail.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.CartDetail.Query().
//		Select(cartdetail.FieldCreatedAt).
//		Scan(ctx, &v)
func (cdq *CartDetailQuery) Select(fields ...string) *CartDetailSelect {
	cdq.fields = append(cdq.fields, fields...)
	selbuild := &CartDetailSelect{CartDetailQuery: cdq}
	selbuild.label = cartdetail.Label
	selbuild.flds, selbuild.scan = &cdq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a CartDetailSelect configured with the given aggregations.
func (cdq *CartDetailQuery) Aggregate(fns ...AggregateFunc) *CartDetailSelect {
	return cdq.Select().Aggregate(fns...)
}

func (cdq *CartDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cdq.fields {
		if !cartdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cdq.path != nil {
		prev, err := cdq.path(ctx)
		if err != nil {
			return err
		}
		cdq.sql = prev
	}
	return nil
}

func (cdq *CartDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CartDetail, error) {
	var (
		nodes       = []*CartDetail{}
		withFKs     = cdq.withFKs
		_spec       = cdq.querySpec()
		loadedTypes = [1]bool{
			cdq.withCart != nil,
		}
	)
	if cdq.withCart != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cartdetail.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CartDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CartDetail{config: cdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cdq.withCart; query != nil {
		if err := cdq.loadCart(ctx, query, nodes, nil,
			func(n *CartDetail, e *Cart) { n.Edges.Cart = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cdq *CartDetailQuery) loadCart(ctx context.Context, query *CartQuery, nodes []*CartDetail, init func(*CartDetail), assign func(*CartDetail, *Cart)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CartDetail)
	for i := range nodes {
		if nodes[i].cart_cart_details == nil {
			continue
		}
		fk := *nodes[i].cart_cart_details
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(cart.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cart_cart_details" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cdq *CartDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cdq.querySpec()
	_spec.Node.Columns = cdq.fields
	if len(cdq.fields) > 0 {
		_spec.Unique = cdq.unique != nil && *cdq.unique
	}
	return sqlgraph.CountNodes(ctx, cdq.driver, _spec)
}

func (cdq *CartDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cdq *CartDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cartdetail.Table,
			Columns: cartdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cartdetail.FieldID,
			},
		},
		From:   cdq.sql,
		Unique: true,
	}
	if unique := cdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartdetail.FieldID)
		for i := range fields {
			if fields[i] != cartdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cdq *CartDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cdq.driver.Dialect())
	t1 := builder.Table(cartdetail.Table)
	columns := cdq.fields
	if len(columns) == 0 {
		columns = cartdetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cdq.sql != nil {
		selector = cdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cdq.unique != nil && *cdq.unique {
		selector.Distinct()
	}
	for _, p := range cdq.predicates {
		p(selector)
	}
	for _, p := range cdq.order {
		p(selector)
	}
	if offset := cdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CartDetailGroupBy is the group-by builder for CartDetail entities.
type CartDetailGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cdgb *CartDetailGroupBy) Aggregate(fns ...AggregateFunc) *CartDetailGroupBy {
	cdgb.fns = append(cdgb.fns, fns...)
	return cdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cdgb *CartDetailGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cdgb.path(ctx)
	if err != nil {
		return err
	}
	cdgb.sql = query
	return cdgb.sqlScan(ctx, v)
}

func (cdgb *CartDetailGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cdgb.fields {
		if !cartdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cdgb *CartDetailGroupBy) sqlQuery() *sql.Selector {
	selector := cdgb.sql.Select()
	aggregation := make([]string, 0, len(cdgb.fns))
	for _, fn := range cdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cdgb.fields)+len(cdgb.fns))
		for _, f := range cdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cdgb.fields...)...)
}

// CartDetailSelect is the builder for selecting fields of CartDetail entities.
type CartDetailSelect struct {
	*CartDetailQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cds *CartDetailSelect) Aggregate(fns ...AggregateFunc) *CartDetailSelect {
	cds.fns = append(cds.fns, fns...)
	return cds
}

// Scan applies the selector query and scans the result into the given value.
func (cds *CartDetailSelect) Scan(ctx context.Context, v any) error {
	if err := cds.prepareQuery(ctx); err != nil {
		return err
	}
	cds.sql = cds.CartDetailQuery.sqlQuery(ctx)
	return cds.sqlScan(ctx, v)
}

func (cds *CartDetailSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(cds.fns))
	for _, fn := range cds.fns {
		aggregation = append(aggregation, fn(cds.sql))
	}
	switch n := len(*cds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		cds.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		cds.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := cds.sql.Query()
	if err := cds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
