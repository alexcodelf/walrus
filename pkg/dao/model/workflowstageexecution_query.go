// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// WorkflowStageExecutionQuery is the builder for querying WorkflowStageExecution entities.
type WorkflowStageExecutionQuery struct {
	config
	ctx                   *QueryContext
	order                 []workflowstageexecution.OrderOption
	inters                []Interceptor
	predicates            []predicate.WorkflowStageExecution
	withProject           *ProjectQuery
	withSteps             *WorkflowStepExecutionQuery
	withStage             *WorkflowStageQuery
	withWorkflowExecution *WorkflowExecutionQuery
	modifiers             []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkflowStageExecutionQuery builder.
func (wseq *WorkflowStageExecutionQuery) Where(ps ...predicate.WorkflowStageExecution) *WorkflowStageExecutionQuery {
	wseq.predicates = append(wseq.predicates, ps...)
	return wseq
}

// Limit the number of records to be returned by this query.
func (wseq *WorkflowStageExecutionQuery) Limit(limit int) *WorkflowStageExecutionQuery {
	wseq.ctx.Limit = &limit
	return wseq
}

// Offset to start from.
func (wseq *WorkflowStageExecutionQuery) Offset(offset int) *WorkflowStageExecutionQuery {
	wseq.ctx.Offset = &offset
	return wseq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wseq *WorkflowStageExecutionQuery) Unique(unique bool) *WorkflowStageExecutionQuery {
	wseq.ctx.Unique = &unique
	return wseq
}

// Order specifies how the records should be ordered.
func (wseq *WorkflowStageExecutionQuery) Order(o ...workflowstageexecution.OrderOption) *WorkflowStageExecutionQuery {
	wseq.order = append(wseq.order, o...)
	return wseq
}

// QueryProject chains the current query on the "project" edge.
func (wseq *WorkflowStageExecutionQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: wseq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wseq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wseq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowstageexecution.Table, workflowstageexecution.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflowstageexecution.ProjectTable, workflowstageexecution.ProjectColumn),
		)
		schemaConfig := wseq.schemaConfig
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.WorkflowStageExecution
		fromU = sqlgraph.SetNeighbors(wseq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySteps chains the current query on the "steps" edge.
func (wseq *WorkflowStageExecutionQuery) QuerySteps() *WorkflowStepExecutionQuery {
	query := (&WorkflowStepExecutionClient{config: wseq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wseq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wseq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowstageexecution.Table, workflowstageexecution.FieldID, selector),
			sqlgraph.To(workflowstepexecution.Table, workflowstepexecution.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workflowstageexecution.StepsTable, workflowstageexecution.StepsColumn),
		)
		schemaConfig := wseq.schemaConfig
		step.To.Schema = schemaConfig.WorkflowStepExecution
		step.Edge.Schema = schemaConfig.WorkflowStepExecution
		fromU = sqlgraph.SetNeighbors(wseq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStage chains the current query on the "stage" edge.
func (wseq *WorkflowStageExecutionQuery) QueryStage() *WorkflowStageQuery {
	query := (&WorkflowStageClient{config: wseq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wseq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wseq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowstageexecution.Table, workflowstageexecution.FieldID, selector),
			sqlgraph.To(workflowstage.Table, workflowstage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflowstageexecution.StageTable, workflowstageexecution.StageColumn),
		)
		schemaConfig := wseq.schemaConfig
		step.To.Schema = schemaConfig.WorkflowStage
		step.Edge.Schema = schemaConfig.WorkflowStageExecution
		fromU = sqlgraph.SetNeighbors(wseq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflowExecution chains the current query on the "workflow_execution" edge.
func (wseq *WorkflowStageExecutionQuery) QueryWorkflowExecution() *WorkflowExecutionQuery {
	query := (&WorkflowExecutionClient{config: wseq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wseq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wseq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workflowstageexecution.Table, workflowstageexecution.FieldID, selector),
			sqlgraph.To(workflowexecution.Table, workflowexecution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workflowstageexecution.WorkflowExecutionTable, workflowstageexecution.WorkflowExecutionColumn),
		)
		schemaConfig := wseq.schemaConfig
		step.To.Schema = schemaConfig.WorkflowExecution
		step.Edge.Schema = schemaConfig.WorkflowStageExecution
		fromU = sqlgraph.SetNeighbors(wseq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkflowStageExecution entity from the query.
// Returns a *NotFoundError when no WorkflowStageExecution was found.
func (wseq *WorkflowStageExecutionQuery) First(ctx context.Context) (*WorkflowStageExecution, error) {
	nodes, err := wseq.Limit(1).All(setContextOp(ctx, wseq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workflowstageexecution.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) FirstX(ctx context.Context) *WorkflowStageExecution {
	node, err := wseq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkflowStageExecution ID from the query.
// Returns a *NotFoundError when no WorkflowStageExecution ID was found.
func (wseq *WorkflowStageExecutionQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = wseq.Limit(1).IDs(setContextOp(ctx, wseq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workflowstageexecution.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := wseq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkflowStageExecution entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WorkflowStageExecution entity is found.
// Returns a *NotFoundError when no WorkflowStageExecution entities are found.
func (wseq *WorkflowStageExecutionQuery) Only(ctx context.Context) (*WorkflowStageExecution, error) {
	nodes, err := wseq.Limit(2).All(setContextOp(ctx, wseq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workflowstageexecution.Label}
	default:
		return nil, &NotSingularError{workflowstageexecution.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) OnlyX(ctx context.Context) *WorkflowStageExecution {
	node, err := wseq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkflowStageExecution ID in the query.
// Returns a *NotSingularError when more than one WorkflowStageExecution ID is found.
// Returns a *NotFoundError when no entities are found.
func (wseq *WorkflowStageExecutionQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = wseq.Limit(2).IDs(setContextOp(ctx, wseq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workflowstageexecution.Label}
	default:
		err = &NotSingularError{workflowstageexecution.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := wseq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkflowStageExecutions.
func (wseq *WorkflowStageExecutionQuery) All(ctx context.Context) ([]*WorkflowStageExecution, error) {
	ctx = setContextOp(ctx, wseq.ctx, "All")
	if err := wseq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WorkflowStageExecution, *WorkflowStageExecutionQuery]()
	return withInterceptors[[]*WorkflowStageExecution](ctx, wseq, qr, wseq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) AllX(ctx context.Context) []*WorkflowStageExecution {
	nodes, err := wseq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkflowStageExecution IDs.
func (wseq *WorkflowStageExecutionQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if wseq.ctx.Unique == nil && wseq.path != nil {
		wseq.Unique(true)
	}
	ctx = setContextOp(ctx, wseq.ctx, "IDs")
	if err = wseq.Select(workflowstageexecution.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := wseq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wseq *WorkflowStageExecutionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wseq.ctx, "Count")
	if err := wseq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wseq, querierCount[*WorkflowStageExecutionQuery](), wseq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) CountX(ctx context.Context) int {
	count, err := wseq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wseq *WorkflowStageExecutionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wseq.ctx, "Exist")
	switch _, err := wseq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wseq *WorkflowStageExecutionQuery) ExistX(ctx context.Context) bool {
	exist, err := wseq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkflowStageExecutionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wseq *WorkflowStageExecutionQuery) Clone() *WorkflowStageExecutionQuery {
	if wseq == nil {
		return nil
	}
	return &WorkflowStageExecutionQuery{
		config:                wseq.config,
		ctx:                   wseq.ctx.Clone(),
		order:                 append([]workflowstageexecution.OrderOption{}, wseq.order...),
		inters:                append([]Interceptor{}, wseq.inters...),
		predicates:            append([]predicate.WorkflowStageExecution{}, wseq.predicates...),
		withProject:           wseq.withProject.Clone(),
		withSteps:             wseq.withSteps.Clone(),
		withStage:             wseq.withStage.Clone(),
		withWorkflowExecution: wseq.withWorkflowExecution.Clone(),
		// clone intermediate query.
		sql:  wseq.sql.Clone(),
		path: wseq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (wseq *WorkflowStageExecutionQuery) WithProject(opts ...func(*ProjectQuery)) *WorkflowStageExecutionQuery {
	query := (&ProjectClient{config: wseq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wseq.withProject = query
	return wseq
}

// WithSteps tells the query-builder to eager-load the nodes that are connected to
// the "steps" edge. The optional arguments are used to configure the query builder of the edge.
func (wseq *WorkflowStageExecutionQuery) WithSteps(opts ...func(*WorkflowStepExecutionQuery)) *WorkflowStageExecutionQuery {
	query := (&WorkflowStepExecutionClient{config: wseq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wseq.withSteps = query
	return wseq
}

// WithStage tells the query-builder to eager-load the nodes that are connected to
// the "stage" edge. The optional arguments are used to configure the query builder of the edge.
func (wseq *WorkflowStageExecutionQuery) WithStage(opts ...func(*WorkflowStageQuery)) *WorkflowStageExecutionQuery {
	query := (&WorkflowStageClient{config: wseq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wseq.withStage = query
	return wseq
}

// WithWorkflowExecution tells the query-builder to eager-load the nodes that are connected to
// the "workflow_execution" edge. The optional arguments are used to configure the query builder of the edge.
func (wseq *WorkflowStageExecutionQuery) WithWorkflowExecution(opts ...func(*WorkflowExecutionQuery)) *WorkflowStageExecutionQuery {
	query := (&WorkflowExecutionClient{config: wseq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wseq.withWorkflowExecution = query
	return wseq
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
//	client.WorkflowStageExecution.Query().
//		GroupBy(workflowstageexecution.FieldName).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (wseq *WorkflowStageExecutionQuery) GroupBy(field string, fields ...string) *WorkflowStageExecutionGroupBy {
	wseq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WorkflowStageExecutionGroupBy{build: wseq}
	grbuild.flds = &wseq.ctx.Fields
	grbuild.label = workflowstageexecution.Label
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
//	client.WorkflowStageExecution.Query().
//		Select(workflowstageexecution.FieldName).
//		Scan(ctx, &v)
func (wseq *WorkflowStageExecutionQuery) Select(fields ...string) *WorkflowStageExecutionSelect {
	wseq.ctx.Fields = append(wseq.ctx.Fields, fields...)
	sbuild := &WorkflowStageExecutionSelect{WorkflowStageExecutionQuery: wseq}
	sbuild.label = workflowstageexecution.Label
	sbuild.flds, sbuild.scan = &wseq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WorkflowStageExecutionSelect configured with the given aggregations.
func (wseq *WorkflowStageExecutionQuery) Aggregate(fns ...AggregateFunc) *WorkflowStageExecutionSelect {
	return wseq.Select().Aggregate(fns...)
}

func (wseq *WorkflowStageExecutionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wseq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wseq); err != nil {
				return err
			}
		}
	}
	for _, f := range wseq.ctx.Fields {
		if !workflowstageexecution.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if wseq.path != nil {
		prev, err := wseq.path(ctx)
		if err != nil {
			return err
		}
		wseq.sql = prev
	}
	return nil
}

func (wseq *WorkflowStageExecutionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WorkflowStageExecution, error) {
	var (
		nodes       = []*WorkflowStageExecution{}
		_spec       = wseq.querySpec()
		loadedTypes = [4]bool{
			wseq.withProject != nil,
			wseq.withSteps != nil,
			wseq.withStage != nil,
			wseq.withWorkflowExecution != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WorkflowStageExecution).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WorkflowStageExecution{config: wseq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = wseq.schemaConfig.WorkflowStageExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseq.schemaConfig)
	if len(wseq.modifiers) > 0 {
		_spec.Modifiers = wseq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wseq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wseq.withProject; query != nil {
		if err := wseq.loadProject(ctx, query, nodes, nil,
			func(n *WorkflowStageExecution, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := wseq.withSteps; query != nil {
		if err := wseq.loadSteps(ctx, query, nodes,
			func(n *WorkflowStageExecution) { n.Edges.Steps = []*WorkflowStepExecution{} },
			func(n *WorkflowStageExecution, e *WorkflowStepExecution) { n.Edges.Steps = append(n.Edges.Steps, e) }); err != nil {
			return nil, err
		}
	}
	if query := wseq.withStage; query != nil {
		if err := wseq.loadStage(ctx, query, nodes, nil,
			func(n *WorkflowStageExecution, e *WorkflowStage) { n.Edges.Stage = e }); err != nil {
			return nil, err
		}
	}
	if query := wseq.withWorkflowExecution; query != nil {
		if err := wseq.loadWorkflowExecution(ctx, query, nodes, nil,
			func(n *WorkflowStageExecution, e *WorkflowExecution) { n.Edges.WorkflowExecution = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wseq *WorkflowStageExecutionQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*WorkflowStageExecution, init func(*WorkflowStageExecution), assign func(*WorkflowStageExecution, *Project)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*WorkflowStageExecution)
	for i := range nodes {
		fk := nodes[i].ProjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wseq *WorkflowStageExecutionQuery) loadSteps(ctx context.Context, query *WorkflowStepExecutionQuery, nodes []*WorkflowStageExecution, init func(*WorkflowStageExecution), assign func(*WorkflowStageExecution, *WorkflowStepExecution)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[object.ID]*WorkflowStageExecution)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workflowstepexecution.FieldWorkflowStageExecutionID)
	}
	query.Where(predicate.WorkflowStepExecution(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(workflowstageexecution.StepsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WorkflowStageExecutionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "workflow_stage_execution_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wseq *WorkflowStageExecutionQuery) loadStage(ctx context.Context, query *WorkflowStageQuery, nodes []*WorkflowStageExecution, init func(*WorkflowStageExecution), assign func(*WorkflowStageExecution, *WorkflowStage)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*WorkflowStageExecution)
	for i := range nodes {
		fk := nodes[i].StageID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflowstage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "stage_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wseq *WorkflowStageExecutionQuery) loadWorkflowExecution(ctx context.Context, query *WorkflowExecutionQuery, nodes []*WorkflowStageExecution, init func(*WorkflowStageExecution), assign func(*WorkflowStageExecution, *WorkflowExecution)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*WorkflowStageExecution)
	for i := range nodes {
		fk := nodes[i].WorkflowExecutionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(workflowexecution.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "workflow_execution_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wseq *WorkflowStageExecutionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wseq.querySpec()
	_spec.Node.Schema = wseq.schemaConfig.WorkflowStageExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseq.schemaConfig)
	if len(wseq.modifiers) > 0 {
		_spec.Modifiers = wseq.modifiers
	}
	_spec.Node.Columns = wseq.ctx.Fields
	if len(wseq.ctx.Fields) > 0 {
		_spec.Unique = wseq.ctx.Unique != nil && *wseq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wseq.driver, _spec)
}

func (wseq *WorkflowStageExecutionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(workflowstageexecution.Table, workflowstageexecution.Columns, sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString))
	_spec.From = wseq.sql
	if unique := wseq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wseq.path != nil {
		_spec.Unique = true
	}
	if fields := wseq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowstageexecution.FieldID)
		for i := range fields {
			if fields[i] != workflowstageexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wseq.withProject != nil {
			_spec.Node.AddColumnOnce(workflowstageexecution.FieldProjectID)
		}
		if wseq.withStage != nil {
			_spec.Node.AddColumnOnce(workflowstageexecution.FieldStageID)
		}
		if wseq.withWorkflowExecution != nil {
			_spec.Node.AddColumnOnce(workflowstageexecution.FieldWorkflowExecutionID)
		}
	}
	if ps := wseq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wseq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wseq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wseq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wseq *WorkflowStageExecutionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wseq.driver.Dialect())
	t1 := builder.Table(workflowstageexecution.Table)
	columns := wseq.ctx.Fields
	if len(columns) == 0 {
		columns = workflowstageexecution.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wseq.sql != nil {
		selector = wseq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wseq.ctx.Unique != nil && *wseq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(wseq.schemaConfig.WorkflowStageExecution)
	ctx = internal.NewSchemaConfigContext(ctx, wseq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range wseq.modifiers {
		m(selector)
	}
	for _, p := range wseq.predicates {
		p(selector)
	}
	for _, p := range wseq.order {
		p(selector)
	}
	if offset := wseq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wseq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wseq *WorkflowStageExecutionQuery) ForUpdate(opts ...sql.LockOption) *WorkflowStageExecutionQuery {
	if wseq.driver.Dialect() == dialect.Postgres {
		wseq.Unique(false)
	}
	wseq.modifiers = append(wseq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wseq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wseq *WorkflowStageExecutionQuery) ForShare(opts ...sql.LockOption) *WorkflowStageExecutionQuery {
	if wseq.driver.Dialect() == dialect.Postgres {
		wseq.Unique(false)
	}
	wseq.modifiers = append(wseq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wseq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wseq *WorkflowStageExecutionQuery) Modify(modifiers ...func(s *sql.Selector)) *WorkflowStageExecutionSelect {
	wseq.modifiers = append(wseq.modifiers, modifiers...)
	return wseq.Select()
}

// WhereP appends storage-level predicates to the WorkflowStageExecutionQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (wseq *WorkflowStageExecutionQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.WorkflowStageExecution, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.WorkflowStageExecution(ps[i]))
	}
	wseq.predicates = append(wseq.predicates, wps...)
}

// WorkflowStageExecutionGroupBy is the group-by builder for WorkflowStageExecution entities.
type WorkflowStageExecutionGroupBy struct {
	selector
	build *WorkflowStageExecutionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wsegb *WorkflowStageExecutionGroupBy) Aggregate(fns ...AggregateFunc) *WorkflowStageExecutionGroupBy {
	wsegb.fns = append(wsegb.fns, fns...)
	return wsegb
}

// Scan applies the selector query and scans the result into the given value.
func (wsegb *WorkflowStageExecutionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wsegb.build.ctx, "GroupBy")
	if err := wsegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowStageExecutionQuery, *WorkflowStageExecutionGroupBy](ctx, wsegb.build, wsegb, wsegb.build.inters, v)
}

func (wsegb *WorkflowStageExecutionGroupBy) sqlScan(ctx context.Context, root *WorkflowStageExecutionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wsegb.fns))
	for _, fn := range wsegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wsegb.flds)+len(wsegb.fns))
		for _, f := range *wsegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wsegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wsegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WorkflowStageExecutionSelect is the builder for selecting fields of WorkflowStageExecution entities.
type WorkflowStageExecutionSelect struct {
	*WorkflowStageExecutionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wses *WorkflowStageExecutionSelect) Aggregate(fns ...AggregateFunc) *WorkflowStageExecutionSelect {
	wses.fns = append(wses.fns, fns...)
	return wses
}

// Scan applies the selector query and scans the result into the given value.
func (wses *WorkflowStageExecutionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wses.ctx, "Select")
	if err := wses.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WorkflowStageExecutionQuery, *WorkflowStageExecutionSelect](ctx, wses.WorkflowStageExecutionQuery, wses, wses.inters, v)
}

func (wses *WorkflowStageExecutionSelect) sqlScan(ctx context.Context, root *WorkflowStageExecutionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wses.fns))
	for _, fn := range wses.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wses.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wses.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (wses *WorkflowStageExecutionSelect) Modify(modifiers ...func(s *sql.Selector)) *WorkflowStageExecutionSelect {
	wses.modifiers = append(wses.modifiers, modifiers...)
	return wses
}
