// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowCreateInput holds the creation input of the Workflow entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Type of the workflow.
	Type string `path:"-" query:"-" json:"type"`
	// Display name is the human readable name that is shown to the user.
	DisplayName string `path:"-" query:"-" json:"displayName"`
	// ID of the project that this workflow belongs to.
	ProjectID object.ID `path:"-" query:"-" json:"projectID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID of the environment that this workflow belongs to.
	EnvironmentID object.ID `path:"-" query:"-" json:"environmentID,omitempty"`
	// ID list of the stages that belong to this workflow.
	WorkflowStageIds []object.ID `path:"-" query:"-" json:"workflowStageIds,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
}

// Model returns the Workflow entity for creating,
// after validating.
func (wci *WorkflowCreateInput) Model() *Workflow {
	if wci == nil {
		return nil
	}

	_w := &Workflow{
		Type:             wci.Type,
		DisplayName:      wci.DisplayName,
		ProjectID:        wci.ProjectID,
		Name:             wci.Name,
		Description:      wci.Description,
		Labels:           wci.Labels,
		EnvironmentID:    wci.EnvironmentID,
		WorkflowStageIds: wci.WorkflowStageIds,
		Parallelism:      wci.Parallelism,
	}

	return _w
}

// Validate checks the WorkflowCreateInput entity.
func (wci *WorkflowCreateInput) Validate() error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	return wci.ValidateWith(wci.inputConfig.Context, wci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowCreateInput entity with the given context and client set.
func (wci *WorkflowCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowCreateInputs holds the creation input item of the Workflow entities.
type WorkflowCreateInputsItem struct {
	// Type of the workflow.
	Type string `path:"-" query:"-" json:"type"`
	// Display name is the human readable name that is shown to the user.
	DisplayName string `path:"-" query:"-" json:"displayName"`
	// ID of the project that this workflow belongs to.
	ProjectID object.ID `path:"-" query:"-" json:"projectID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID of the environment that this workflow belongs to.
	EnvironmentID object.ID `path:"-" query:"-" json:"environmentID,omitempty"`
	// ID list of the stages that belong to this workflow.
	WorkflowStageIds []object.ID `path:"-" query:"-" json:"workflowStageIds,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
}

// ValidateWith checks the WorkflowCreateInputsItem entity with the given context and client set.
func (wci *WorkflowCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowCreateInputs holds the creation input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for creating,
// after validating.
func (wci *WorkflowCreateInputs) Model() []*Workflow {
	if wci == nil || len(wci.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wci.Items))

	for i := range wci.Items {
		_w := &Workflow{
			Type:             wci.Items[i].Type,
			DisplayName:      wci.Items[i].DisplayName,
			ProjectID:        wci.Items[i].ProjectID,
			Name:             wci.Items[i].Name,
			Description:      wci.Items[i].Description,
			Labels:           wci.Items[i].Labels,
			EnvironmentID:    wci.Items[i].EnvironmentID,
			WorkflowStageIds: wci.Items[i].WorkflowStageIds,
			Parallelism:      wci.Items[i].Parallelism,
		}

		_ws[i] = _w
	}

	return _ws
}

// Validate checks the WorkflowCreateInputs entity .
func (wci *WorkflowCreateInputs) Validate() error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	return wci.ValidateWith(wci.inputConfig.Context, wci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowCreateInputs entity with the given context and client set.
func (wci *WorkflowCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if len(wci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wci.Items {
		if wci.Items[i] == nil {
			continue
		}

		if err := wci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowDeleteInput holds the deletion input of the Workflow entity,
// please tags with `path:",inline"` if embedding.
type WorkflowDeleteInput struct {
	WorkflowQueryInput `path:",inline"`
}

// WorkflowDeleteInputs holds the deletion input item of the Workflow entities.
type WorkflowDeleteInputsItem struct {
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// WorkflowDeleteInputs holds the deletion input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for deleting,
// after validating.
func (wdi *WorkflowDeleteInputs) Model() []*Workflow {
	if wdi == nil || len(wdi.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wdi.Items))
	for i := range wdi.Items {
		_ws[i] = &Workflow{
			ID: wdi.Items[i].ID,
		}
	}
	return _ws
}

// IDs returns the ID list of the Workflow entities for deleting,
// after validating.
func (wdi *WorkflowDeleteInputs) IDs() []object.ID {
	if wdi == nil || len(wdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wdi.Items))
	for i := range wdi.Items {
		ids[i] = wdi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowDeleteInputs entity.
func (wdi *WorkflowDeleteInputs) Validate() error {
	if wdi == nil {
		return errors.New("nil receiver")
	}

	return wdi.ValidateWith(wdi.inputConfig.Context, wdi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowDeleteInputs entity with the given context and client set.
func (wdi *WorkflowDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wdi == nil {
		return errors.New("nil receiver")
	}

	if len(wdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	ids := make([]object.ID, 0, len(wdi.Items))
	ors := make([]predicate.Workflow, 0, len(wdi.Items))
	indexers := make(map[any][]int)

	for i := range wdi.Items {
		if wdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wdi.Items[i].ID != "" {
			ids = append(ids, wdi.Items[i].ID)
			ors = append(ors, workflow.ID(wdi.Items[i].ID))
			indexers[wdi.Items[i].ID] = append(indexers[wdi.Items[i].ID], i)
		} else if wdi.Items[i].Name != "" {
			ors = append(ors, workflow.And(
				workflow.Name(wdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", wdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := workflow.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = workflow.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			workflow.FieldID,
			workflow.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			wdi.Items[j].ID = es[i].ID
			wdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// WorkflowQueryInput holds the query input of the Workflow entity,
// please tags with `path:",inline"` if embedding.
type WorkflowQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the Workflow entity.
	Refer *object.Refer `path:"workflow,default=" query:"-" json:"-"`
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Workflow entity for querying,
// after validating.
func (wqi *WorkflowQueryInput) Model() *Workflow {
	if wqi == nil {
		return nil
	}

	return &Workflow{
		ID:   wqi.ID,
		Name: wqi.Name,
	}
}

// Validate checks the WorkflowQueryInput entity.
func (wqi *WorkflowQueryInput) Validate() error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	return wqi.ValidateWith(wqi.inputConfig.Context, wqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowQueryInput entity with the given context and client set.
func (wqi *WorkflowQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	if wqi.Refer != nil && *wqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflow.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	if wqi.Refer != nil {
		if wqi.Refer.IsID() {
			q.Where(
				workflow.ID(wqi.Refer.ID()))
		} else if refers := wqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				workflow.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of workflow")
		}
	} else if wqi.ID != "" {
		q.Where(
			workflow.ID(wqi.ID))
	} else if wqi.Name != "" {
		q.Where(
			workflow.Name(wqi.Name))
	} else {
		return errors.New("invalid identify of workflow")
	}

	q.Select(
		workflow.FieldID,
		workflow.FieldName,
	)

	var e *Workflow
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Workflow)
		}
	}

	wqi.ID = e.ID
	wqi.Name = e.Name
	return nil
}

// WorkflowQueryInputs holds the query input of the Workflow entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the WorkflowQueryInputs entity.
func (wqi *WorkflowQueryInputs) Validate() error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	return wqi.ValidateWith(wqi.inputConfig.Context, wqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowQueryInputs entity with the given context and client set.
func (wqi *WorkflowQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowUpdateInput holds the modification input of the Workflow entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowUpdateInput struct {
	WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Display name is the human readable name that is shown to the user.
	DisplayName string `path:"-" query:"-" json:"displayName,omitempty"`
	// ID list of the stages that belong to this workflow.
	WorkflowStageIds []object.ID `path:"-" query:"-" json:"workflowStageIds,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
}

// Model returns the Workflow entity for modifying,
// after validating.
func (wui *WorkflowUpdateInput) Model() *Workflow {
	if wui == nil {
		return nil
	}

	_w := &Workflow{
		ID:               wui.ID,
		Name:             wui.Name,
		Description:      wui.Description,
		Labels:           wui.Labels,
		DisplayName:      wui.DisplayName,
		WorkflowStageIds: wui.WorkflowStageIds,
		Parallelism:      wui.Parallelism,
	}

	return _w
}

// Validate checks the WorkflowUpdateInput entity.
func (wui *WorkflowUpdateInput) Validate() error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	return wui.ValidateWith(wui.inputConfig.Context, wui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowUpdateInput entity with the given context and client set.
func (wui *WorkflowUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wui.WorkflowQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// WorkflowUpdateInputs holds the modification input item of the Workflow entities.
type WorkflowUpdateInputsItem struct {
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Display name is the human readable name that is shown to the user.
	DisplayName string `path:"-" query:"-" json:"displayName"`
	// ID list of the stages that belong to this workflow.
	WorkflowStageIds []object.ID `path:"-" query:"-" json:"workflowStageIds"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism"`
}

// ValidateWith checks the WorkflowUpdateInputsItem entity with the given context and client set.
func (wui *WorkflowUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowUpdateInputs holds the modification input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for modifying,
// after validating.
func (wui *WorkflowUpdateInputs) Model() []*Workflow {
	if wui == nil || len(wui.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wui.Items))

	for i := range wui.Items {
		_w := &Workflow{
			ID:               wui.Items[i].ID,
			Name:             wui.Items[i].Name,
			Description:      wui.Items[i].Description,
			Labels:           wui.Items[i].Labels,
			DisplayName:      wui.Items[i].DisplayName,
			WorkflowStageIds: wui.Items[i].WorkflowStageIds,
			Parallelism:      wui.Items[i].Parallelism,
		}

		_ws[i] = _w
	}

	return _ws
}

// IDs returns the ID list of the Workflow entities for modifying,
// after validating.
func (wui *WorkflowUpdateInputs) IDs() []object.ID {
	if wui == nil || len(wui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wui.Items))
	for i := range wui.Items {
		ids[i] = wui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowUpdateInputs entity.
func (wui *WorkflowUpdateInputs) Validate() error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	return wui.ValidateWith(wui.inputConfig.Context, wui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowUpdateInputs entity with the given context and client set.
func (wui *WorkflowUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	if len(wui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	ids := make([]object.ID, 0, len(wui.Items))
	ors := make([]predicate.Workflow, 0, len(wui.Items))
	indexers := make(map[any][]int)

	for i := range wui.Items {
		if wui.Items[i] == nil {
			return errors.New("nil item")
		}

		if wui.Items[i].ID != "" {
			ids = append(ids, wui.Items[i].ID)
			ors = append(ors, workflow.ID(wui.Items[i].ID))
			indexers[wui.Items[i].ID] = append(indexers[wui.Items[i].ID], i)
		} else if wui.Items[i].Name != "" {
			ors = append(ors, workflow.And(
				workflow.Name(wui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", wui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := workflow.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = workflow.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			workflow.FieldID,
			workflow.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			wui.Items[j].ID = es[i].ID
			wui.Items[j].Name = es[i].Name
		}
	}

	for i := range wui.Items {
		if err := wui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowOutput holds the output of the Workflow entity.
type WorkflowOutput struct {
	ID               object.ID         `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	CreateTime       *time.Time        `json:"createTime,omitempty"`
	UpdateTime       *time.Time        `json:"updateTime,omitempty"`
	Status           status.Status     `json:"status,omitempty"`
	ProjectID        object.ID         `json:"projectID,omitempty"`
	EnvironmentID    object.ID         `json:"environmentID,omitempty"`
	DisplayName      string            `json:"displayName,omitempty"`
	Type             string            `json:"type,omitempty"`
	WorkflowStageIds []object.ID       `json:"workflowStageIds,omitempty"`
	Parallelism      int               `json:"parallelism,omitempty"`
}

// View returns the output of Workflow entity.
func (_w *Workflow) View() *WorkflowOutput {
	return ExposeWorkflow(_w)
}

// View returns the output of Workflow entities.
func (_ws Workflows) View() []*WorkflowOutput {
	return ExposeWorkflows(_ws)
}

// ExposeWorkflow converts the Workflow to WorkflowOutput.
func ExposeWorkflow(_w *Workflow) *WorkflowOutput {
	if _w == nil {
		return nil
	}

	wo := &WorkflowOutput{
		ID:               _w.ID,
		Name:             _w.Name,
		Description:      _w.Description,
		Labels:           _w.Labels,
		CreateTime:       _w.CreateTime,
		UpdateTime:       _w.UpdateTime,
		Status:           _w.Status,
		ProjectID:        _w.ProjectID,
		EnvironmentID:    _w.EnvironmentID,
		DisplayName:      _w.DisplayName,
		Type:             _w.Type,
		WorkflowStageIds: _w.WorkflowStageIds,
		Parallelism:      _w.Parallelism,
	}

	return wo
}

// ExposeWorkflows converts the Workflow slice to WorkflowOutput pointer slice.
func ExposeWorkflows(_ws []*Workflow) []*WorkflowOutput {
	if len(_ws) == 0 {
		return nil
	}

	wos := make([]*WorkflowOutput, len(_ws))
	for i := range _ws {
		wos[i] = ExposeWorkflow(_ws[i])
	}
	return wos
}
