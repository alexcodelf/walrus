// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowsteptemplate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// WorkflowStepTemplateQuery is the builder for querying WorkflowStepTemplate entities.
type WorkflowStepTemplateQuery struct {
	config
	ctx        *QueryContext
	order      []workflowsteptemplate.OrderOption
	inters     []Interceptor
	predicates []predicate.WorkflowStepTemplate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkflowStepTemplateQuery builder.
func (wstq *WorkflowStepTemplateQuery) Where(ps ...predicate.WorkflowStepTemplate) *WorkflowStepTemplateQuery {
	wstq.predicates = append(wstq.predicates, ps...)
	return wstq
}

// Limit the number of records to be returned by this query.
func (wstq *WorkflowStepTemplateQuery) Limit(limit int) *WorkflowStepTemplateQuery {
	wstq.ctx.Limit = &limit
	return wstq
}

// Offset to start from.
func (wstq *WorkflowStepTemplateQuery) Offset(offset int) *WorkflowStepTemplateQuery {
	wstq.ctx.Offset = &offset
	return wstq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wstq *WorkflowStepTemplateQuery) Unique(unique bool) *WorkflowStepTemplateQuery {
	wstq.ctx.Unique = &unique
	return wstq
}

// Order specifies how the records should be ordered.
func (wstq *WorkflowStepTemplateQuery) Order(o ...workflowsteptemplate.OrderOption) *WorkflowStepTemplateQuery {
	wstq.order = append(wstq.order, o...)
	return wstq
}

// First returns the first WorkflowStepTemplate entity from the query.
// Returns a *NotFoundError when no WorkflowStepTemplate was found.
func (wstq *WorkflowStepTemplateQuery) First(ctx context.Context) (*WorkflowStepTemplate, error) {
	nodes, err := wstq.Limit(1).All(setContextOp(ctx, wstq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workflowsteptemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) FirstX(ctx context.Context) *WorkflowStepTemplate {
	node, err := wstq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkflowStepTemplate ID from the query.
// Returns a *NotFoundError when no WorkflowStepTemplate ID was found.
func (wstq *WorkflowStepTemplateQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = wstq.Limit(1).IDs(setContextOp(ctx, wstq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workflowsteptemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := wstq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkflowStepTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkflowStepTemplate entity is found.
// Returns a *NotFoundError when no WorkflowStepTemplate entities are found.
func (wstq *WorkflowStepTemplateQuery) Only(ctx context.Context) (*WorkflowStepTemplate, error) {
	nodes, err := wstq.Limit(2).All(setContextOp(ctx, wstq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workflowsteptemplate.Label}
	default:
		return nil, &NotSingularError{workflowsteptemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) OnlyX(ctx context.Context) *WorkflowStepTemplate {
	node, err := wstq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkflowStepTemplate ID in the query.
// Returns a *NotSingularError when more than one WorkflowStepTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (wstq *WorkflowStepTemplateQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = wstq.Limit(2).IDs(setContextOp(ctx, wstq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workflowsteptemplate.Label}
	default:
		err = &NotSingularError{workflowsteptemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := wstq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkflowStepTemplates.
func (wstq *WorkflowStepTemplateQuery) All(ctx context.Context) ([]*WorkflowStepTemplate, error) {
	ctx = setContextOp(ctx, wstq.ctx, "All")
	if err := wstq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkflowStepTemplate, *WorkflowStepTemplateQuery]()
	return withInterceptors[[]*WorkflowStepTemplate](ctx, wstq, qr, wstq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) AllX(ctx context.Context) []*WorkflowStepTemplate {
	nodes, err := wstq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkflowStepTemplate IDs.
func (wstq *WorkflowStepTemplateQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if wstq.ctx.Unique == nil && wstq.path != nil {
		wstq.Unique(true)
	}
	ctx = setContextOp(ctx, wstq.ctx, "IDs")
	if err = wstq.Select(workflowsteptemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := wstq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wstq *WorkflowStepTemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wstq.ctx, "Count")
	if err := wstq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wstq, querierCount[*WorkflowStepTemplateQuery](), wstq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) CountX(ctx context.Context) int {
	count, err := wstq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wstq *WorkflowStepTemplateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wstq.ctx, "Exist")
	switch _, err := wstq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wstq *WorkflowStepTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := wstq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkflowStepTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wstq *WorkflowStepTemplateQuery) Clone() *WorkflowStepTemplateQuery {
	if wstq == nil {
		return nil
	}
	return &WorkflowStepTemplateQuery{
		config:     wstq.config,
		ctx:        wstq.ctx.Clone(),
		order:      append([]workflowsteptemplate.OrderOption{}, wstq.order...),
		inters:     append([]Interceptor{}, wstq.inters...),
		predicates: append([]predicate.WorkflowStepTemplate{}, wstq.predicates...),
		// clone intermediate query.
		sql:  wstq.sql.Clone(),
		path: wstq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkflowStepTemplate.Query().
//		GroupBy(workflowsteptemplate.FieldName).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (wstq *WorkflowStepTemplateQuery) GroupBy(field string, fields ...string) *WorkflowStepTemplateGroupBy {
	wstq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkflowStepTemplateGroupBy{build: wstq}
	grbuild.flds = &wstq.ctx.Fields
	grbuild.label = workflowsteptemplate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WorkflowStepTemplate.Query().
//		Select(workflowsteptemplate.FieldName).
//		Scan(ctx, &v)
func (wstq *WorkflowStepTemplateQuery) Select(fields ...string) *WorkflowStepTemplateSelect {
	wstq.ctx.Fields = append(wstq.ctx.Fields, fields...)
	sbuild := &WorkflowStepTemplateSelect{WorkflowStepTemplateQuery: wstq}
	sbuild.label = workflowsteptemplate.Label
	sbuild.flds, sbuild.scan = &wstq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkflowStepTemplateSelect configured with the given aggregations.
func (wstq *WorkflowStepTemplateQuery) Aggregate(fns ...AggregateFunc) *WorkflowStepTemplateSelect {
	return wstq.Select().Aggregate(fns...)
}

func (wstq *WorkflowStepTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wstq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wstq); err != nil {
				return err
			}
		}
	}
	for _, f := range wstq.ctx.Fields {
		if !workflowsteptemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if wstq.path != nil {
		prev, err := wstq.path(ctx)
		if err != nil {
			return err
		}
		wstq.sql = prev
	}
	return nil
}

func (wstq *WorkflowStepTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkflowStepTemplate, error) {
	var (
		nodes = []*WorkflowStepTemplate{}
		_spec = wstq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkflowStepTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkflowStepTemplate{config: wstq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = wstq.schemaConfig.WorkflowStepTemplate
	ctx = internal.NewSchemaConfigContext(ctx, wstq.schemaConfig)
	if len(wstq.modifiers) > 0 {
		_spec.Modifiers = wstq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wstq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wstq *WorkflowStepTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wstq.querySpec()
	_spec.Node.Schema = wstq.schemaConfig.WorkflowStepTemplate
	ctx = internal.NewSchemaConfigContext(ctx, wstq.schemaConfig)
	if len(wstq.modifiers) > 0 {
		_spec.Modifiers = wstq.modifiers
	}
	_spec.Node.Columns = wstq.ctx.Fields
	if len(wstq.ctx.Fields) > 0 {
		_spec.Unique = wstq.ctx.Unique != nil && *wstq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wstq.driver, _spec)
}

func (wstq *WorkflowStepTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workflowsteptemplate.Table, workflowsteptemplate.Columns, sqlgraph.NewFieldSpec(workflowsteptemplate.FieldID, field.TypeString))
	_spec.From = wstq.sql
	if unique := wstq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wstq.path != nil {
		_spec.Unique = true
	}
	if fields := wstq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowsteptemplate.FieldID)
		for i := range fields {
			if fields[i] != workflowsteptemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wstq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wstq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wstq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wstq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wstq *WorkflowStepTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wstq.driver.Dialect())
	t1 := builder.Table(workflowsteptemplate.Table)
	columns := wstq.ctx.Fields
	if len(columns) == 0 {
		columns = workflowsteptemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wstq.sql != nil {
		selector = wstq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wstq.ctx.Unique != nil && *wstq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(wstq.schemaConfig.WorkflowStepTemplate)
	ctx = internal.NewSchemaConfigContext(ctx, wstq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range wstq.modifiers {
		m(selector)
	}
	for _, p := range wstq.predicates {
		p(selector)
	}
	for _, p := range wstq.order {
		p(selector)
	}
	if offset := wstq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wstq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wstq *WorkflowStepTemplateQuery) ForUpdate(opts ...sql.LockOption) *WorkflowStepTemplateQuery {
	if wstq.driver.Dialect() == dialect.Postgres {
		wstq.Unique(false)
	}
	wstq.modifiers = append(wstq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wstq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wstq *WorkflowStepTemplateQuery) ForShare(opts ...sql.LockOption) *WorkflowStepTemplateQuery {
	if wstq.driver.Dialect() == dialect.Postgres {
		wstq.Unique(false)
	}
	wstq.modifiers = append(wstq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wstq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wstq *WorkflowStepTemplateQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkflowStepTemplateSelect {
	wstq.modifiers = append(wstq.modifiers, modifiers...)
	return wstq.Select()
}

// WhereP appends storage-level predicates to the WorkflowStepTemplateQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (wstq *WorkflowStepTemplateQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.WorkflowStepTemplate, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.WorkflowStepTemplate(ps[i]))
	}
	wstq.predicates = append(wstq.predicates, wps...)
}

// WorkflowStepTemplateGroupBy is the group-by builder for WorkflowStepTemplate entities.
type WorkflowStepTemplateGroupBy struct {
	selector
	build *WorkflowStepTemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wstgb *WorkflowStepTemplateGroupBy) Aggregate(fns ...AggregateFunc) *WorkflowStepTemplateGroupBy {
	wstgb.fns = append(wstgb.fns, fns...)
	return wstgb
}

// Scan applies the selector query and scans the result into the given value.
func (wstgb *WorkflowStepTemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wstgb.build.ctx, "GroupBy")
	if err := wstgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowStepTemplateQuery, *WorkflowStepTemplateGroupBy](ctx, wstgb.build, wstgb, wstgb.build.inters, v)
}

func (wstgb *WorkflowStepTemplateGroupBy) sqlScan(ctx context.Context, root *WorkflowStepTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wstgb.fns))
	for _, fn := range wstgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wstgb.flds)+len(wstgb.fns))
		for _, f := range *wstgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wstgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wstgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkflowStepTemplateSelect is the builder for selecting fields of WorkflowStepTemplate entities.
type WorkflowStepTemplateSelect struct {
	*WorkflowStepTemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wsts *WorkflowStepTemplateSelect) Aggregate(fns ...AggregateFunc) *WorkflowStepTemplateSelect {
	wsts.fns = append(wsts.fns, fns...)
	return wsts
}

// Scan applies the selector query and scans the result into the given value.
func (wsts *WorkflowStepTemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wsts.ctx, "Select")
	if err := wsts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowStepTemplateQuery, *WorkflowStepTemplateSelect](ctx, wsts.WorkflowStepTemplateQuery, wsts, wsts.inters, v)
}

func (wsts *WorkflowStepTemplateSelect) sqlScan(ctx context.Context, root *WorkflowStepTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wsts.fns))
	for _, fn := range wsts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wsts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wsts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wsts *WorkflowStepTemplateSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkflowStepTemplateSelect {
	wsts.modifiers = append(wsts.modifiers, modifiers...)
	return wsts
}