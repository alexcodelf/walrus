// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowStepExecutionCreateInput holds the creation input of the WorkflowStepExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStepExecutionCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// StageExecution indicates to create WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Type of the workflow step execution.
	Type string `path:"-" query:"-" json:"type"`
	// ID of the workflow that this workflow step execution belongs to.
	WorkflowID object.ID `path:"-" query:"-" json:"workflowID"`
	// ID of the workflow execution that this workflow step execution belongs to.
	WorkflowExecutionID object.ID `path:"-" query:"-" json:"workflowExecutionID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Spec of the workflow step execution.
	Spec map[string]interface{} `path:"-" query:"-" json:"spec,omitempty"`
	// Number of times that this workflow step execution has been executed.
	Times int `path:"-" query:"-" json:"times,omitempty"`
	// Duration of the workflow step execution.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Log record of the workflow step execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow step execution. It's the yaml file that defines the workflow step execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// Model returns the WorkflowStepExecution entity for creating,
// after validating.
func (wseci *WorkflowStepExecutionCreateInput) Model() *WorkflowStepExecution {
	if wseci == nil {
		return nil
	}

	_wse := &WorkflowStepExecution{
		Type:                wseci.Type,
		WorkflowID:          wseci.WorkflowID,
		WorkflowExecutionID: wseci.WorkflowExecutionID,
		Name:                wseci.Name,
		Description:         wseci.Description,
		Labels:              wseci.Labels,
		Spec:                wseci.Spec,
		Times:               wseci.Times,
		Duration:            wseci.Duration,
		Record:              wseci.Record,
		Input:               wseci.Input,
	}

	if wseci.StageExecution != nil {
		_wse.WorkflowStageExecutionID = wseci.StageExecution.ID
	}
	if wseci.Project != nil {
		_wse.ProjectID = wseci.Project.ID
	}

	return _wse
}

// Validate checks the WorkflowStepExecutionCreateInput entity.
func (wseci *WorkflowStepExecutionCreateInput) Validate() error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	return wseci.ValidateWith(wseci.inputConfig.Context, wseci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionCreateInput entity with the given context and client set.
func (wseci *WorkflowStepExecutionCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if wseci.Project != nil {
		if err := wseci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}
	// Validate when creating under the StageExecution route.
	if wseci.StageExecution != nil {
		if err := wseci.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStepExecutionCreateInputs holds the creation input item of the WorkflowStepExecution entities.
type WorkflowStepExecutionCreateInputsItem struct {
	// Type of the workflow step execution.
	Type string `path:"-" query:"-" json:"type"`
	// ID of the workflow that this workflow step execution belongs to.
	WorkflowID object.ID `path:"-" query:"-" json:"workflowID"`
	// ID of the workflow execution that this workflow step execution belongs to.
	WorkflowExecutionID object.ID `path:"-" query:"-" json:"workflowExecutionID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Spec of the workflow step execution.
	Spec map[string]interface{} `path:"-" query:"-" json:"spec,omitempty"`
	// Number of times that this workflow step execution has been executed.
	Times int `path:"-" query:"-" json:"times,omitempty"`
	// Duration of the workflow step execution.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Log record of the workflow step execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow step execution. It's the yaml file that defines the workflow step execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// ValidateWith checks the WorkflowStepExecutionCreateInputsItem entity with the given context and client set.
func (wseci *WorkflowStepExecutionCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStepExecutionCreateInputs holds the creation input of the WorkflowStepExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStepExecutionCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// StageExecution indicates to create WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStepExecutionCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStepExecution entities for creating,
// after validating.
func (wseci *WorkflowStepExecutionCreateInputs) Model() []*WorkflowStepExecution {
	if wseci == nil || len(wseci.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStepExecution, len(wseci.Items))

	for i := range wseci.Items {
		_wse := &WorkflowStepExecution{
			Type:                wseci.Items[i].Type,
			WorkflowID:          wseci.Items[i].WorkflowID,
			WorkflowExecutionID: wseci.Items[i].WorkflowExecutionID,
			Name:                wseci.Items[i].Name,
			Description:         wseci.Items[i].Description,
			Labels:              wseci.Items[i].Labels,
			Spec:                wseci.Items[i].Spec,
			Times:               wseci.Items[i].Times,
			Duration:            wseci.Items[i].Duration,
			Record:              wseci.Items[i].Record,
			Input:               wseci.Items[i].Input,
		}

		if wseci.StageExecution != nil {
			_wse.WorkflowStageExecutionID = wseci.StageExecution.ID
		}
		if wseci.Project != nil {
			_wse.ProjectID = wseci.Project.ID
		}

		_wses[i] = _wse
	}

	return _wses
}

// Validate checks the WorkflowStepExecutionCreateInputs entity .
func (wseci *WorkflowStepExecutionCreateInputs) Validate() error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	return wseci.ValidateWith(wseci.inputConfig.Context, wseci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionCreateInputs entity with the given context and client set.
func (wseci *WorkflowStepExecutionCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseci == nil {
		return errors.New("nil receiver")
	}

	if len(wseci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if wseci.Project != nil {
		if err := wseci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseci.Project = nil
			}
		}
	}
	// Validate when creating under the StageExecution route.
	if wseci.StageExecution != nil {
		if err := wseci.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wseci.StageExecution = nil
			}
		}
	}

	for i := range wseci.Items {
		if wseci.Items[i] == nil {
			continue
		}

		if err := wseci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStepExecutionDeleteInput holds the deletion input of the WorkflowStepExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStepExecutionDeleteInput struct {
	WorkflowStepExecutionQueryInput `path:",inline"`
}

// WorkflowStepExecutionDeleteInputs holds the deletion input item of the WorkflowStepExecution entities.
type WorkflowStepExecutionDeleteInputsItem struct {
	// ID of the WorkflowStepExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// WorkflowStepExecutionDeleteInputs holds the deletion input of the WorkflowStepExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStepExecutionDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// StageExecution indicates to delete WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStepExecutionDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStepExecution entities for deleting,
// after validating.
func (wsedi *WorkflowStepExecutionDeleteInputs) Model() []*WorkflowStepExecution {
	if wsedi == nil || len(wsedi.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStepExecution, len(wsedi.Items))
	for i := range wsedi.Items {
		_wses[i] = &WorkflowStepExecution{
			ID: wsedi.Items[i].ID,
		}
	}
	return _wses
}

// IDs returns the ID list of the WorkflowStepExecution entities for deleting,
// after validating.
func (wsedi *WorkflowStepExecutionDeleteInputs) IDs() []object.ID {
	if wsedi == nil || len(wsedi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wsedi.Items))
	for i := range wsedi.Items {
		ids[i] = wsedi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStepExecutionDeleteInputs entity.
func (wsedi *WorkflowStepExecutionDeleteInputs) Validate() error {
	if wsedi == nil {
		return errors.New("nil receiver")
	}

	return wsedi.ValidateWith(wsedi.inputConfig.Context, wsedi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionDeleteInputs entity with the given context and client set.
func (wsedi *WorkflowStepExecutionDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wsedi == nil {
		return errors.New("nil receiver")
	}

	if len(wsedi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStepExecutions().Query()

	// Validate when deleting under the Project route.
	if wsedi.Project != nil {
		if err := wsedi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflowstepexecution.ProjectID(wsedi.Project.ID))
		}
	}

	// Validate when deleting under the StageExecution route.
	if wsedi.StageExecution != nil {
		if err := wsedi.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstepexecution.WorkflowStageExecutionID(wsedi.StageExecution.ID))
		}
	}

	ids := make([]object.ID, 0, len(wsedi.Items))

	for i := range wsedi.Items {
		if wsedi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wsedi.Items[i].ID != "" {
			ids = append(ids, wsedi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstepexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// WorkflowStepExecutionQueryInput holds the query input of the WorkflowStepExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowStepExecutionQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`
	// StageExecution indicates to query WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"stageExecution"`

	// Refer holds the route path reference of the WorkflowStepExecution entity.
	Refer *object.Refer `path:"workflowstepexecution,default=" query:"-" json:"-"`
	// ID of the WorkflowStepExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the WorkflowStepExecution entity for querying,
// after validating.
func (wseqi *WorkflowStepExecutionQueryInput) Model() *WorkflowStepExecution {
	if wseqi == nil {
		return nil
	}

	return &WorkflowStepExecution{
		ID: wseqi.ID,
	}
}

// Validate checks the WorkflowStepExecutionQueryInput entity.
func (wseqi *WorkflowStepExecutionQueryInput) Validate() error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	return wseqi.ValidateWith(wseqi.inputConfig.Context, wseqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionQueryInput entity with the given context and client set.
func (wseqi *WorkflowStepExecutionQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	if wseqi.Refer != nil && *wseqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflowstepexecution.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStepExecutions().Query()

	// Validate when querying under the Project route.
	if wseqi.Project != nil {
		if err := wseqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflowstepexecution.ProjectID(wseqi.Project.ID))
		}
	}

	// Validate when querying under the StageExecution route.
	if wseqi.StageExecution != nil {
		if err := wseqi.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstepexecution.WorkflowStageExecutionID(wseqi.StageExecution.ID))
		}
	}

	if wseqi.Refer != nil {
		if wseqi.Refer.IsID() {
			q.Where(
				workflowstepexecution.ID(wseqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of workflowstepexecution")
		}
	} else if wseqi.ID != "" {
		q.Where(
			workflowstepexecution.ID(wseqi.ID))
	} else {
		return errors.New("invalid identify of workflowstepexecution")
	}

	q.Select(
		workflowstepexecution.FieldID,
	)

	var e *WorkflowStepExecution
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
			e = cv.(*WorkflowStepExecution)
		}
	}

	wseqi.ID = e.ID
	return nil
}

// WorkflowStepExecutionQueryInputs holds the query input of the WorkflowStepExecution entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowStepExecutionQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// StageExecution indicates to query WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the WorkflowStepExecutionQueryInputs entity.
func (wseqi *WorkflowStepExecutionQueryInputs) Validate() error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	return wseqi.ValidateWith(wseqi.inputConfig.Context, wseqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionQueryInputs entity with the given context and client set.
func (wseqi *WorkflowStepExecutionQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if wseqi.Project != nil {
		if err := wseqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	// Validate when querying under the StageExecution route.
	if wseqi.StageExecution != nil {
		if err := wseqi.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStepExecutionUpdateInput holds the modification input of the WorkflowStepExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStepExecutionUpdateInput struct {
	WorkflowStepExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Spec of the workflow step execution.
	Spec map[string]interface{} `path:"-" query:"-" json:"spec,omitempty"`
	// Number of times that this workflow step execution has been executed.
	Times int `path:"-" query:"-" json:"times,omitempty"`
	// Duration of the workflow step execution.
	Duration int `path:"-" query:"-" json:"duration,omitempty"`
	// Log record of the workflow step execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow step execution. It's the yaml file that defines the workflow step execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// Model returns the WorkflowStepExecution entity for modifying,
// after validating.
func (wseui *WorkflowStepExecutionUpdateInput) Model() *WorkflowStepExecution {
	if wseui == nil {
		return nil
	}

	_wse := &WorkflowStepExecution{
		ID:          wseui.ID,
		Description: wseui.Description,
		Labels:      wseui.Labels,
		Spec:        wseui.Spec,
		Times:       wseui.Times,
		Duration:    wseui.Duration,
		Record:      wseui.Record,
		Input:       wseui.Input,
	}

	return _wse
}

// Validate checks the WorkflowStepExecutionUpdateInput entity.
func (wseui *WorkflowStepExecutionUpdateInput) Validate() error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	return wseui.ValidateWith(wseui.inputConfig.Context, wseui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionUpdateInput entity with the given context and client set.
func (wseui *WorkflowStepExecutionUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wseui.WorkflowStepExecutionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// WorkflowStepExecutionUpdateInputs holds the modification input item of the WorkflowStepExecution entities.
type WorkflowStepExecutionUpdateInputsItem struct {
	// ID of the WorkflowStepExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Spec of the workflow step execution.
	Spec map[string]interface{} `path:"-" query:"-" json:"spec,omitempty"`
	// Number of times that this workflow step execution has been executed.
	Times int `path:"-" query:"-" json:"times"`
	// Duration of the workflow step execution.
	Duration int `path:"-" query:"-" json:"duration"`
	// Log record of the workflow step execution.
	Record string `path:"-" query:"-" json:"record"`
	// Input of the workflow step execution. It's the yaml file that defines the workflow step execution.
	Input string `path:"-" query:"-" json:"input"`
}

// ValidateWith checks the WorkflowStepExecutionUpdateInputsItem entity with the given context and client set.
func (wseui *WorkflowStepExecutionUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowStepExecutionUpdateInputs holds the modification input of the WorkflowStepExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowStepExecutionUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update WorkflowStepExecution entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// StageExecution indicates to update WorkflowStepExecution entity MUST under the StageExecution route.
	StageExecution *WorkflowStageExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowStepExecutionUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowStepExecution entities for modifying,
// after validating.
func (wseui *WorkflowStepExecutionUpdateInputs) Model() []*WorkflowStepExecution {
	if wseui == nil || len(wseui.Items) == 0 {
		return nil
	}

	_wses := make([]*WorkflowStepExecution, len(wseui.Items))

	for i := range wseui.Items {
		_wse := &WorkflowStepExecution{
			ID:          wseui.Items[i].ID,
			Description: wseui.Items[i].Description,
			Labels:      wseui.Items[i].Labels,
			Spec:        wseui.Items[i].Spec,
			Times:       wseui.Items[i].Times,
			Duration:    wseui.Items[i].Duration,
			Record:      wseui.Items[i].Record,
			Input:       wseui.Items[i].Input,
		}

		_wses[i] = _wse
	}

	return _wses
}

// IDs returns the ID list of the WorkflowStepExecution entities for modifying,
// after validating.
func (wseui *WorkflowStepExecutionUpdateInputs) IDs() []object.ID {
	if wseui == nil || len(wseui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wseui.Items))
	for i := range wseui.Items {
		ids[i] = wseui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowStepExecutionUpdateInputs entity.
func (wseui *WorkflowStepExecutionUpdateInputs) Validate() error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	return wseui.ValidateWith(wseui.inputConfig.Context, wseui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowStepExecutionUpdateInputs entity with the given context and client set.
func (wseui *WorkflowStepExecutionUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wseui == nil {
		return errors.New("nil receiver")
	}

	if len(wseui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowStepExecutions().Query()

	// Validate when updating under the Project route.
	if wseui.Project != nil {
		if err := wseui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflowstepexecution.ProjectID(wseui.Project.ID))
		}
	}

	// Validate when updating under the StageExecution route.
	if wseui.StageExecution != nil {
		if err := wseui.StageExecution.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowstepexecution.WorkflowStageExecutionID(wseui.StageExecution.ID))
		}
	}

	ids := make([]object.ID, 0, len(wseui.Items))

	for i := range wseui.Items {
		if wseui.Items[i] == nil {
			return errors.New("nil item")
		}

		if wseui.Items[i].ID != "" {
			ids = append(ids, wseui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowstepexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range wseui.Items {
		if err := wseui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowStepExecutionOutput holds the output of the WorkflowStepExecution entity.
type WorkflowStepExecutionOutput struct {
	ID                  object.ID              `json:"id,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Labels              map[string]string      `json:"labels,omitempty"`
	CreateTime          *time.Time             `json:"createTime,omitempty"`
	UpdateTime          *time.Time             `json:"updateTime,omitempty"`
	Status              status.Status          `json:"status,omitempty"`
	WorkflowExecutionID object.ID              `json:"workflowExecutionID,omitempty"`
	WorkflowID          object.ID              `json:"workflowID,omitempty"`
	Type                string                 `json:"type,omitempty"`
	Spec                map[string]interface{} `json:"spec,omitempty"`
	Times               int                    `json:"times,omitempty"`
	Duration            int                    `json:"duration,omitempty"`
	Record              string                 `json:"record,omitempty"`
	Input               string                 `json:"input,omitempty"`

	Project        *ProjectOutput                `json:"project,omitempty"`
	StageExecution *WorkflowStageExecutionOutput `json:"stageExecution,omitempty"`
}

// View returns the output of WorkflowStepExecution entity.
func (_wse *WorkflowStepExecution) View() *WorkflowStepExecutionOutput {
	return ExposeWorkflowStepExecution(_wse)
}

// View returns the output of WorkflowStepExecution entities.
func (_wses WorkflowStepExecutions) View() []*WorkflowStepExecutionOutput {
	return ExposeWorkflowStepExecutions(_wses)
}

// ExposeWorkflowStepExecution converts the WorkflowStepExecution to WorkflowStepExecutionOutput.
func ExposeWorkflowStepExecution(_wse *WorkflowStepExecution) *WorkflowStepExecutionOutput {
	if _wse == nil {
		return nil
	}

	wseo := &WorkflowStepExecutionOutput{
		ID:                  _wse.ID,
		Name:                _wse.Name,
		Description:         _wse.Description,
		Labels:              _wse.Labels,
		CreateTime:          _wse.CreateTime,
		UpdateTime:          _wse.UpdateTime,
		Status:              _wse.Status,
		WorkflowExecutionID: _wse.WorkflowExecutionID,
		WorkflowID:          _wse.WorkflowID,
		Type:                _wse.Type,
		Spec:                _wse.Spec,
		Times:               _wse.Times,
		Duration:            _wse.Duration,
		Record:              _wse.Record,
		Input:               _wse.Input,
	}

	if _wse.Edges.Project != nil {
		wseo.Project = ExposeProject(_wse.Edges.Project)
	} else if _wse.ProjectID != "" {
		wseo.Project = &ProjectOutput{
			ID: _wse.ProjectID,
		}
	}
	if _wse.Edges.StageExecution != nil {
		wseo.StageExecution = ExposeWorkflowStageExecution(_wse.Edges.StageExecution)
	} else if _wse.WorkflowStageExecutionID != "" {
		wseo.StageExecution = &WorkflowStageExecutionOutput{
			ID: _wse.WorkflowStageExecutionID,
		}
	}
	return wseo
}

// ExposeWorkflowStepExecutions converts the WorkflowStepExecution slice to WorkflowStepExecutionOutput pointer slice.
func ExposeWorkflowStepExecutions(_wses []*WorkflowStepExecution) []*WorkflowStepExecutionOutput {
	if len(_wses) == 0 {
		return nil
	}

	wseos := make([]*WorkflowStepExecutionOutput, len(_wses))
	for i := range _wses {
		wseos[i] = ExposeWorkflowStepExecution(_wses[i])
	}
	return wseos
}
