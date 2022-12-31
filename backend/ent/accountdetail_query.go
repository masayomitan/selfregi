// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"selfregi/ent/account"
	"selfregi/ent/accountdetail"
	"selfregi/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountDetailQuery is the builder for querying AccountDetail entities.
type AccountDetailQuery struct {
	config
	limit              *int
	offset             *int
	unique             *bool
	order              []OrderFunc
	fields             []string
	predicates         []predicate.AccountDetail
	withManagedAccount *AccountQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountDetailQuery builder.
func (adq *AccountDetailQuery) Where(ps ...predicate.AccountDetail) *AccountDetailQuery {
	adq.predicates = append(adq.predicates, ps...)
	return adq
}

// Limit adds a limit step to the query.
func (adq *AccountDetailQuery) Limit(limit int) *AccountDetailQuery {
	adq.limit = &limit
	return adq
}

// Offset adds an offset step to the query.
func (adq *AccountDetailQuery) Offset(offset int) *AccountDetailQuery {
	adq.offset = &offset
	return adq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (adq *AccountDetailQuery) Unique(unique bool) *AccountDetailQuery {
	adq.unique = &unique
	return adq
}

// Order adds an order step to the query.
func (adq *AccountDetailQuery) Order(o ...OrderFunc) *AccountDetailQuery {
	adq.order = append(adq.order, o...)
	return adq
}

// QueryManagedAccount chains the current query on the "managed_account" edge.
func (adq *AccountDetailQuery) QueryManagedAccount() *AccountQuery {
	query := &AccountQuery{config: adq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountdetail.Table, accountdetail.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountdetail.ManagedAccountTable, accountdetail.ManagedAccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(adq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountDetail entity from the query.
// Returns a *NotFoundError when no AccountDetail was found.
func (adq *AccountDetailQuery) First(ctx context.Context) (*AccountDetail, error) {
	nodes, err := adq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountdetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (adq *AccountDetailQuery) FirstX(ctx context.Context) *AccountDetail {
	node, err := adq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountDetail ID from the query.
// Returns a *NotFoundError when no AccountDetail ID was found.
func (adq *AccountDetailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = adq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountdetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (adq *AccountDetailQuery) FirstIDX(ctx context.Context) int {
	id, err := adq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountDetail entity is found.
// Returns a *NotFoundError when no AccountDetail entities are found.
func (adq *AccountDetailQuery) Only(ctx context.Context) (*AccountDetail, error) {
	nodes, err := adq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountdetail.Label}
	default:
		return nil, &NotSingularError{accountdetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (adq *AccountDetailQuery) OnlyX(ctx context.Context) *AccountDetail {
	node, err := adq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountDetail ID in the query.
// Returns a *NotSingularError when more than one AccountDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (adq *AccountDetailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = adq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountdetail.Label}
	default:
		err = &NotSingularError{accountdetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (adq *AccountDetailQuery) OnlyIDX(ctx context.Context) int {
	id, err := adq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountDetails.
func (adq *AccountDetailQuery) All(ctx context.Context) ([]*AccountDetail, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return adq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (adq *AccountDetailQuery) AllX(ctx context.Context) []*AccountDetail {
	nodes, err := adq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountDetail IDs.
func (adq *AccountDetailQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := adq.Select(accountdetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (adq *AccountDetailQuery) IDsX(ctx context.Context) []int {
	ids, err := adq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (adq *AccountDetailQuery) Count(ctx context.Context) (int, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return adq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (adq *AccountDetailQuery) CountX(ctx context.Context) int {
	count, err := adq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (adq *AccountDetailQuery) Exist(ctx context.Context) (bool, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return adq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (adq *AccountDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := adq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (adq *AccountDetailQuery) Clone() *AccountDetailQuery {
	if adq == nil {
		return nil
	}
	return &AccountDetailQuery{
		config:             adq.config,
		limit:              adq.limit,
		offset:             adq.offset,
		order:              append([]OrderFunc{}, adq.order...),
		predicates:         append([]predicate.AccountDetail{}, adq.predicates...),
		withManagedAccount: adq.withManagedAccount.Clone(),
		// clone intermediate query.
		sql:    adq.sql.Clone(),
		path:   adq.path,
		unique: adq.unique,
	}
}

// WithManagedAccount tells the query-builder to eager-load the nodes that are connected to
// the "managed_account" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AccountDetailQuery) WithManagedAccount(opts ...func(*AccountQuery)) *AccountDetailQuery {
	query := &AccountQuery{config: adq.config}
	for _, opt := range opts {
		opt(query)
	}
	adq.withManagedAccount = query
	return adq
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
//	client.AccountDetail.Query().
//		GroupBy(accountdetail.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (adq *AccountDetailQuery) GroupBy(field string, fields ...string) *AccountDetailGroupBy {
	grbuild := &AccountDetailGroupBy{config: adq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return adq.sqlQuery(ctx), nil
	}
	grbuild.label = accountdetail.Label
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
//	client.AccountDetail.Query().
//		Select(accountdetail.FieldCreatedAt).
//		Scan(ctx, &v)
func (adq *AccountDetailQuery) Select(fields ...string) *AccountDetailSelect {
	adq.fields = append(adq.fields, fields...)
	selbuild := &AccountDetailSelect{AccountDetailQuery: adq}
	selbuild.label = accountdetail.Label
	selbuild.flds, selbuild.scan = &adq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AccountDetailSelect configured with the given aggregations.
func (adq *AccountDetailQuery) Aggregate(fns ...AggregateFunc) *AccountDetailSelect {
	return adq.Select().Aggregate(fns...)
}

func (adq *AccountDetailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range adq.fields {
		if !accountdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if adq.path != nil {
		prev, err := adq.path(ctx)
		if err != nil {
			return err
		}
		adq.sql = prev
	}
	return nil
}

func (adq *AccountDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountDetail, error) {
	var (
		nodes       = []*AccountDetail{}
		withFKs     = adq.withFKs
		_spec       = adq.querySpec()
		loadedTypes = [1]bool{
			adq.withManagedAccount != nil,
		}
	)
	if adq.withManagedAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accountdetail.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountDetail{config: adq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, adq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := adq.withManagedAccount; query != nil {
		if err := adq.loadManagedAccount(ctx, query, nodes, nil,
			func(n *AccountDetail, e *Account) { n.Edges.ManagedAccount = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (adq *AccountDetailQuery) loadManagedAccount(ctx context.Context, query *AccountQuery, nodes []*AccountDetail, init func(*AccountDetail), assign func(*AccountDetail, *Account)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AccountDetail)
	for i := range nodes {
		if nodes[i].account_managed_account_details == nil {
			continue
		}
		fk := *nodes[i].account_managed_account_details
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_managed_account_details" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (adq *AccountDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := adq.querySpec()
	_spec.Node.Columns = adq.fields
	if len(adq.fields) > 0 {
		_spec.Unique = adq.unique != nil && *adq.unique
	}
	return sqlgraph.CountNodes(ctx, adq.driver, _spec)
}

func (adq *AccountDetailQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := adq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (adq *AccountDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountdetail.Table,
			Columns: accountdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountdetail.FieldID,
			},
		},
		From:   adq.sql,
		Unique: true,
	}
	if unique := adq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := adq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountdetail.FieldID)
		for i := range fields {
			if fields[i] != accountdetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := adq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := adq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := adq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := adq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (adq *AccountDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(adq.driver.Dialect())
	t1 := builder.Table(accountdetail.Table)
	columns := adq.fields
	if len(columns) == 0 {
		columns = accountdetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if adq.sql != nil {
		selector = adq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if adq.unique != nil && *adq.unique {
		selector.Distinct()
	}
	for _, p := range adq.predicates {
		p(selector)
	}
	for _, p := range adq.order {
		p(selector)
	}
	if offset := adq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := adq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountDetailGroupBy is the group-by builder for AccountDetail entities.
type AccountDetailGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (adgb *AccountDetailGroupBy) Aggregate(fns ...AggregateFunc) *AccountDetailGroupBy {
	adgb.fns = append(adgb.fns, fns...)
	return adgb
}

// Scan applies the group-by query and scans the result into the given value.
func (adgb *AccountDetailGroupBy) Scan(ctx context.Context, v any) error {
	query, err := adgb.path(ctx)
	if err != nil {
		return err
	}
	adgb.sql = query
	return adgb.sqlScan(ctx, v)
}

func (adgb *AccountDetailGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range adgb.fields {
		if !accountdetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := adgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (adgb *AccountDetailGroupBy) sqlQuery() *sql.Selector {
	selector := adgb.sql.Select()
	aggregation := make([]string, 0, len(adgb.fns))
	for _, fn := range adgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(adgb.fields)+len(adgb.fns))
		for _, f := range adgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(adgb.fields...)...)
}

// AccountDetailSelect is the builder for selecting fields of AccountDetail entities.
type AccountDetailSelect struct {
	*AccountDetailQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ads *AccountDetailSelect) Aggregate(fns ...AggregateFunc) *AccountDetailSelect {
	ads.fns = append(ads.fns, fns...)
	return ads
}

// Scan applies the selector query and scans the result into the given value.
func (ads *AccountDetailSelect) Scan(ctx context.Context, v any) error {
	if err := ads.prepareQuery(ctx); err != nil {
		return err
	}
	ads.sql = ads.AccountDetailQuery.sqlQuery(ctx)
	return ads.sqlScan(ctx, v)
}

func (ads *AccountDetailSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ads.fns))
	for _, fn := range ads.fns {
		aggregation = append(aggregation, fn(ads.sql))
	}
	switch n := len(*ads.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ads.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ads.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ads.sql.Query()
	if err := ads.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
