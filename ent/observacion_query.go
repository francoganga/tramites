// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/francoganga/go_reno2/ent/observacion"
	"github.com/francoganga/go_reno2/ent/predicate"
)

// ObservacionQuery is the builder for querying Observacion entities.
type ObservacionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Observacion
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ObservacionQuery builder.
func (oq *ObservacionQuery) Where(ps ...predicate.Observacion) *ObservacionQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *ObservacionQuery) Limit(limit int) *ObservacionQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *ObservacionQuery) Offset(offset int) *ObservacionQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *ObservacionQuery) Unique(unique bool) *ObservacionQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *ObservacionQuery) Order(o ...OrderFunc) *ObservacionQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first Observacion entity from the query.
// Returns a *NotFoundError when no Observacion was found.
func (oq *ObservacionQuery) First(ctx context.Context) (*Observacion, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{observacion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *ObservacionQuery) FirstX(ctx context.Context) *Observacion {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Observacion ID from the query.
// Returns a *NotFoundError when no Observacion ID was found.
func (oq *ObservacionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{observacion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *ObservacionQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Observacion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Observacion entity is found.
// Returns a *NotFoundError when no Observacion entities are found.
func (oq *ObservacionQuery) Only(ctx context.Context) (*Observacion, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{observacion.Label}
	default:
		return nil, &NotSingularError{observacion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *ObservacionQuery) OnlyX(ctx context.Context) *Observacion {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Observacion ID in the query.
// Returns a *NotSingularError when more than one Observacion ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *ObservacionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{observacion.Label}
	default:
		err = &NotSingularError{observacion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *ObservacionQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Observacions.
func (oq *ObservacionQuery) All(ctx context.Context) ([]*Observacion, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *ObservacionQuery) AllX(ctx context.Context) []*Observacion {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Observacion IDs.
func (oq *ObservacionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(observacion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *ObservacionQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *ObservacionQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *ObservacionQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *ObservacionQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *ObservacionQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ObservacionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *ObservacionQuery) Clone() *ObservacionQuery {
	if oq == nil {
		return nil
	}
	return &ObservacionQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		predicates: append([]predicate.Observacion{}, oq.predicates...),
		// clone intermediate query.
		sql:    oq.sql.Clone(),
		path:   oq.path,
		unique: oq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (oq *ObservacionQuery) GroupBy(field string, fields ...string) *ObservacionGroupBy {
	grbuild := &ObservacionGroupBy{config: oq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(ctx), nil
	}
	grbuild.label = observacion.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (oq *ObservacionQuery) Select(fields ...string) *ObservacionSelect {
	oq.fields = append(oq.fields, fields...)
	selbuild := &ObservacionSelect{ObservacionQuery: oq}
	selbuild.label = observacion.Label
	selbuild.flds, selbuild.scan = &oq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ObservacionSelect configured with the given aggregations.
func (oq *ObservacionQuery) Aggregate(fns ...AggregateFunc) *ObservacionSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *ObservacionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oq.fields {
		if !observacion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *ObservacionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Observacion, error) {
	var (
		nodes = []*Observacion{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Observacion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Observacion{config: oq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oq *ObservacionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	_spec.Node.Columns = oq.fields
	if len(oq.fields) > 0 {
		_spec.Unique = oq.unique != nil && *oq.unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *ObservacionQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (oq *ObservacionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   observacion.Table,
			Columns: observacion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observacion.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, observacion.FieldID)
		for i := range fields {
			if fields[i] != observacion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *ObservacionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(observacion.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = observacion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.unique != nil && *oq.unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ObservacionGroupBy is the group-by builder for Observacion entities.
type ObservacionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *ObservacionGroupBy) Aggregate(fns ...AggregateFunc) *ObservacionGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogb *ObservacionGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

func (ogb *ObservacionGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ogb.fields {
		if !observacion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *ObservacionGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql.Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
		for _, f := range ogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ogb.fields...)...)
}

// ObservacionSelect is the builder for selecting fields of Observacion entities.
type ObservacionSelect struct {
	*ObservacionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *ObservacionSelect) Aggregate(fns ...AggregateFunc) *ObservacionSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *ObservacionSelect) Scan(ctx context.Context, v any) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	os.sql = os.ObservacionQuery.sqlQuery(ctx)
	return os.sqlScan(ctx, v)
}

func (os *ObservacionSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(os.sql))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		os.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		os.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := os.sql.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
