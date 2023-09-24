// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowStageCreateInput holds the creation input of the WorkflowStage entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// IDs of the workflow steps that belong to this workflow stage.
	WorkflowStepIds []object.ID `path:"-" query:"-" json:"workflowStepIds"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Duration of the workflow stage.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// ID list of the workflow stages that this workflow stage depends on.
	Dependencies []object.ID `path:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the WorkflowStage entity for creating,
// after validating.
func (wsci *WorkflowStageCreateInput) Model() *WorkflowStage {
	if wsci == nil {
		return nil
	}

	_ws := &WorkflowStage{
		WorkflowStepIds: wsci.WorkflowStepIds,
		Name:            wsci.Name,
		Description:     wsci.Description,
		Labels:          wsci.Labels,
		Duration:        wsci.Duration,
		Dependencies:    wsci.Dependencies,
	}

	return _ws
}

// Validate checks the WorkflowStageCreateInput entity.
func (wsci *WorkflowStageCreateInput) Validate() error {
	if wsci == nil {
		return errors.New("nil receiver")
	}

	return wsci.ValidateWith(wsci.inputConfig.Context, wsci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageCreateInput entity with the given context and client set.
func (wsci *WorkflowStageCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStageCreateInputs holds the creation input item of the WorkflowStage entities.
type WorkflowStageCreateInputsItem struct {
	// IDs of the workflow steps that belong to this workflow stage.
	WorkflowStepIds []object.ID `path:"-" query:"-" json:"workflowStepIds"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Duration of the workflow stage.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// ID list of the workflow stages that this workflow stage depends on.
	Dependencies []object.ID `path:"-" query:"-" json:"dependencies,omitempty"`
}

// ValidateWith checks the WorkflowStageCreateInputsItem entity with the given context and client set.
func (wsci *WorkflowStageCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStageCreateInputs holds the creation input of the WorkflowStage entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStage entities for creating,
// after validating.
func (wsci *WorkflowStageCreateInputs) Model() []*WorkflowStage {
	if wsci == nil || len(wsci.Items) == 0 {
		return nil
	}

	_wss := make([]*WorkflowStage, len(wsci.Items))

	for i := range wsci.Items {
		_ws := &WorkflowStage{
			WorkflowStepIds: wsci.Items[i].WorkflowStepIds,
			Name:            wsci.Items[i].Name,
			Description:     wsci.Items[i].Description,
			Labels:          wsci.Items[i].Labels,
			Duration:        wsci.Items[i].Duration,
			Dependencies:    wsci.Items[i].Dependencies,
		}

		_wss[i] = _ws
	}

	return _wss
}

// Validate checks the WorkflowStageCreateInputs entity .
func (wsci *WorkflowStageCreateInputs) Validate() error {
	if wsci == nil {
		return errors.New("nil receiver")
	}

	return wsci.ValidateWith(wsci.inputConfig.Context, wsci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageCreateInputs entity with the given context and client set.
func (wsci *WorkflowStageCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsci == nil {
		return errors.New("nil receiver")
	}

	if len(wsci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wsci.Items {
		if wsci.Items[i] == nil {
			continue
		}

		if err := wsci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStageDeleteInput holds the deletion input of the WorkflowStage entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStageDeleteInput struct {
	WorkflowStageQueryInput `path:",inline"`
}

// WorkflowStageDeleteInputs holds the deletion input item of the WorkflowStage entities.
type WorkflowStageDeleteInputsItem struct {
	// ID of the WorkflowStage entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// WorkflowStageDeleteInputs holds the deletion input of the WorkflowStage entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStage entities for deleting,
// after validating.
func (wsdi *WorkflowStageDeleteInputs) Model() []*WorkflowStage {
	if wsdi == nil || len(wsdi.Items) == 0 {
		return nil
	}

	_wss := make([]*WorkflowStage, len(wsdi.Items))
	for i := range wsdi.Items {
		_wss[i] = &WorkflowStage{
			ID: wsdi.Items[i].ID,
		}
	}
	return _wss
}

// IDs returns the ID list of the WorkflowStage entities for deleting,
// after validating.
func (wsdi *WorkflowStageDeleteInputs) IDs() []object.ID {
	if wsdi == nil || len(wsdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wsdi.Items))
	for i := range wsdi.Items {
		ids[i] = wsdi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStageDeleteInputs entity.
func (wsdi *WorkflowStageDeleteInputs) Validate() error {
	if wsdi == nil {
		return errors.New("nil receiver")
	}

	return wsdi.ValidateWith(wsdi.inputConfig.Context, wsdi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageDeleteInputs entity with the given context and client set.
func (wsdi *WorkflowStageDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsdi == nil {
		return errors.New("nil receiver")
	}

	if len(wsdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStages().Query()

	ids := make([]object.ID, 0, len(wsdi.Items))

	for i := range wsdi.Items {
		if wsdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wsdi.Items[i].ID != "" {
			ids = append(ids, wsdi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstage.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// WorkflowStageQueryInput holds the query input of the WorkflowStage entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStageQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the WorkflowStage entity.
	Refer *object.Refer `path:"workflowstage,default=" query:"-" json:"-"`
	// ID of the WorkflowStage entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the WorkflowStage entity for querying,
// after validating.
func (wsqi *WorkflowStageQueryInput) Model() *WorkflowStage {
	if wsqi == nil {
		return nil
	}

	return &WorkflowStage{
		ID: wsqi.ID,
	}
}

// Validate checks the WorkflowStageQueryInput entity.
func (wsqi *WorkflowStageQueryInput) Validate() error {
	if wsqi == nil {
		return errors.New("nil receiver")
	}

	return wsqi.ValidateWith(wsqi.inputConfig.Context, wsqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageQueryInput entity with the given context and client set.
func (wsqi *WorkflowStageQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsqi == nil {
		return errors.New("nil receiver")
	}

	if wsqi.Refer != nil && *wsqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflowstage.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStages().Query()

	if wsqi.Refer != nil {
		if wsqi.Refer.IsID() {
			q.Where(
				workflowstage.ID(wsqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of workflowstage")
		}
	} else if wsqi.ID != "" {
		q.Where(
			workflowstage.ID(wsqi.ID))
	} else {
		return errors.New("invalid identify of workflowstage")
	}

	q.Select(
		workflowstage.FieldID,
	)

	var e *WorkflowStage
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
			e = cv.(*WorkflowStage)
		}
	}

	wsqi.ID = e.ID
	return nil
}

// WorkflowStageQueryInputs holds the query input of the WorkflowStage entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowStageQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the WorkflowStageQueryInputs entity.
func (wsqi *WorkflowStageQueryInputs) Validate() error {
	if wsqi == nil {
		return errors.New("nil receiver")
	}

	return wsqi.ValidateWith(wsqi.inputConfig.Context, wsqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageQueryInputs entity with the given context and client set.
func (wsqi *WorkflowStageQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStageUpdateInput holds the modification input of the WorkflowStage entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageUpdateInput struct {
	WorkflowStageQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Duration of the workflow stage.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// ID list of the workflow stages that this workflow stage depends on.
	Dependencies []object.ID `path:"-" query:"-" json:"dependencies,omitempty"`
}

// Model returns the WorkflowStage entity for modifying,
// after validating.
func (wsui *WorkflowStageUpdateInput) Model() *WorkflowStage {
	if wsui == nil {
		return nil
	}

	_ws := &WorkflowStage{
		ID:           wsui.ID,
		Description:  wsui.Description,
		Labels:       wsui.Labels,
		Duration:     wsui.Duration,
		Dependencies: wsui.Dependencies,
	}

	return _ws
}

// Validate checks the WorkflowStageUpdateInput entity.
func (wsui *WorkflowStageUpdateInput) Validate() error {
	if wsui == nil {
		return errors.New("nil receiver")
	}

	return wsui.ValidateWith(wsui.inputConfig.Context, wsui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageUpdateInput entity with the given context and client set.
func (wsui *WorkflowStageUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wsui.WorkflowStageQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// WorkflowStageUpdateInputs holds the modification input item of the WorkflowStage entities.
type WorkflowStageUpdateInputsItem struct {
	// ID of the WorkflowStage entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Duration of the workflow stage.
	Duration int `path:"-" query:"-" json:"duration"`
	// ID list of the workflow stages that this workflow stage depends on.
	Dependencies []object.ID `path:"-" query:"-" json:"dependencies"`
}

// ValidateWith checks the WorkflowStageUpdateInputsItem entity with the given context and client set.
func (wsui *WorkflowStageUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStageUpdateInputs holds the modification input of the WorkflowStage entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStageUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStageUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStage entities for modifying,
// after validating.
func (wsui *WorkflowStageUpdateInputs) Model() []*WorkflowStage {
	if wsui == nil || len(wsui.Items) == 0 {
		return nil
	}

	_wss := make([]*WorkflowStage, len(wsui.Items))

	for i := range wsui.Items {
		_ws := &WorkflowStage{
			ID:           wsui.Items[i].ID,
			Description:  wsui.Items[i].Description,
			Labels:       wsui.Items[i].Labels,
			Duration:     wsui.Items[i].Duration,
			Dependencies: wsui.Items[i].Dependencies,
		}

		_wss[i] = _ws
	}

	return _wss
}

// IDs returns the ID list of the WorkflowStage entities for modifying,
// after validating.
func (wsui *WorkflowStageUpdateInputs) IDs() []object.ID {
	if wsui == nil || len(wsui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wsui.Items))
	for i := range wsui.Items {
		ids[i] = wsui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStageUpdateInputs entity.
func (wsui *WorkflowStageUpdateInputs) Validate() error {
	if wsui == nil {
		return errors.New("nil receiver")
	}

	return wsui.ValidateWith(wsui.inputConfig.Context, wsui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStageUpdateInputs entity with the given context and client set.
func (wsui *WorkflowStageUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsui == nil {
		return errors.New("nil receiver")
	}

	if len(wsui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStages().Query()

	ids := make([]object.ID, 0, len(wsui.Items))

	for i := range wsui.Items {
		if wsui.Items[i] == nil {
			return errors.New("nil item")
		}

		if wsui.Items[i].ID != "" {
			ids = append(ids, wsui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstage.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range wsui.Items {
		if err := wsui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStageOutput holds the output of the WorkflowStage entity.
type WorkflowStageOutput struct {
	ID              object.ID         `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	CreateTime      *time.Time        `json:"createTime,omitempty"`
	UpdateTime      *time.Time        `json:"updateTime,omitempty"`
	Status          status.Status     `json:"status,omitempty"`
	WorkflowStepIds []object.ID       `json:"workflowStepIds,omitempty"`
	Duration        int               `json:"duration,omitempty"`
	Dependencies    []object.ID       `json:"dependencies,omitempty"`
}

// View returns the output of WorkflowStage entity.
func (_ws *WorkflowStage) View() *WorkflowStageOutput {
	return ExposeWorkflowStage(_ws)
}

// View returns the output of WorkflowStage entities.
func (_wss WorkflowStages) View() []*WorkflowStageOutput {
	return ExposeWorkflowStages(_wss)
}

// ExposeWorkflowStage converts the WorkflowStage to WorkflowStageOutput.
func ExposeWorkflowStage(_ws *WorkflowStage) *WorkflowStageOutput {
	if _ws == nil {
		return nil
	}

	wso := &WorkflowStageOutput{
		ID:              _ws.ID,
		Name:            _ws.Name,
		Description:     _ws.Description,
		Labels:          _ws.Labels,
		CreateTime:      _ws.CreateTime,
		UpdateTime:      _ws.UpdateTime,
		Status:          _ws.Status,
		WorkflowStepIds: _ws.WorkflowStepIds,
		Duration:        _ws.Duration,
		Dependencies:    _ws.Dependencies,
	}

	return wso
}

// ExposeWorkflowStages converts the WorkflowStage slice to WorkflowStageOutput pointer slice.
func ExposeWorkflowStages(_wss []*WorkflowStage) []*WorkflowStageOutput {
	if len(_wss) == 0 {
		return nil
	}

	wsos := make([]*WorkflowStageOutput, len(_wss))
	for i := range _wss {
		wsos[i] = ExposeWorkflowStage(_wss[i])
	}
	return wsos
}
