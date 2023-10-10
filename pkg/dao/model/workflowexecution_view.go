// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowExecutionCreateInput holds the creation input of the WorkflowExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowExecutionCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to create WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Duration of the workflow execution.
	Duration int `path:"-" query:"-" json:"duration"`
	// Progress of the workflow. N/M format,N is number of stages completed, M is total number of stages.
	Progress int `path:"-" query:"-" json:"progress"`
	// ID of the subject that this workflow execution belongs to.
	Subject object.ID `path:"-" query:"-" json:"subject"`
	// ID of the project to belong.
	ProjectID object.ID `path:"-" query:"-" json:"projectID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID list of the stage executions that belong to this workflow execution.
	WorkflowStagesExecution []object.ID `path:"-" query:"-" json:"workflowStagesExecution,omitempty"`
	// Log record of the workflow execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow execution. It's the yaml file that defines the workflow execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// Model returns the WorkflowExecution entity for creating,
// after validating.
func (weci *WorkflowExecutionCreateInput) Model() *WorkflowExecution {
	if weci == nil {
		return nil
	}

	_we := &WorkflowExecution{
		Duration:                weci.Duration,
		Progress:                weci.Progress,
		Subject:                 weci.Subject,
		ProjectID:               weci.ProjectID,
		Name:                    weci.Name,
		Description:             weci.Description,
		Labels:                  weci.Labels,
		WorkflowStagesExecution: weci.WorkflowStagesExecution,
		Record:                  weci.Record,
		Input:                   weci.Input,
	}

	if weci.Workflow != nil {
		_we.WorkflowID = weci.Workflow.ID
	}

	return _we
}

// Validate checks the WorkflowExecutionCreateInput entity.
func (weci *WorkflowExecutionCreateInput) Validate() error {
	if weci == nil {
		return errors.New("nil receiver")
	}

	return weci.ValidateWith(weci.inputConfig.Context, weci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionCreateInput entity with the given context and client set.
func (weci *WorkflowExecutionCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Workflow route.
	if weci.Workflow != nil {
		if err := weci.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowExecutionCreateInputs holds the creation input item of the WorkflowExecution entities.
type WorkflowExecutionCreateInputsItem struct {
	// Duration of the workflow execution.
	Duration int `path:"-" query:"-" json:"duration"`
	// Progress of the workflow. N/M format,N is number of stages completed, M is total number of stages.
	Progress int `path:"-" query:"-" json:"progress"`
	// ID of the subject that this workflow execution belongs to.
	Subject object.ID `path:"-" query:"-" json:"subject"`
	// ID of the project to belong.
	ProjectID object.ID `path:"-" query:"-" json:"projectID"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID list of the stage executions that belong to this workflow execution.
	WorkflowStagesExecution []object.ID `path:"-" query:"-" json:"workflowStagesExecution,omitempty"`
	// Log record of the workflow execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow execution. It's the yaml file that defines the workflow execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// ValidateWith checks the WorkflowExecutionCreateInputsItem entity with the given context and client set.
func (weci *WorkflowExecutionCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowExecutionCreateInputs holds the creation input of the WorkflowExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowExecutionCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to create WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowExecutionCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowExecution entities for creating,
// after validating.
func (weci *WorkflowExecutionCreateInputs) Model() []*WorkflowExecution {
	if weci == nil || len(weci.Items) == 0 {
		return nil
	}

	_wes := make([]*WorkflowExecution, len(weci.Items))

	for i := range weci.Items {
		_we := &WorkflowExecution{
			Duration:                weci.Items[i].Duration,
			Progress:                weci.Items[i].Progress,
			Subject:                 weci.Items[i].Subject,
			ProjectID:               weci.Items[i].ProjectID,
			Name:                    weci.Items[i].Name,
			Description:             weci.Items[i].Description,
			Labels:                  weci.Items[i].Labels,
			WorkflowStagesExecution: weci.Items[i].WorkflowStagesExecution,
			Record:                  weci.Items[i].Record,
			Input:                   weci.Items[i].Input,
		}

		if weci.Workflow != nil {
			_we.WorkflowID = weci.Workflow.ID
		}

		_wes[i] = _we
	}

	return _wes
}

// Validate checks the WorkflowExecutionCreateInputs entity .
func (weci *WorkflowExecutionCreateInputs) Validate() error {
	if weci == nil {
		return errors.New("nil receiver")
	}

	return weci.ValidateWith(weci.inputConfig.Context, weci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionCreateInputs entity with the given context and client set.
func (weci *WorkflowExecutionCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weci == nil {
		return errors.New("nil receiver")
	}

	if len(weci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Workflow route.
	if weci.Workflow != nil {
		if err := weci.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				weci.Workflow = nil
			}
		}
	}

	for i := range weci.Items {
		if weci.Items[i] == nil {
			continue
		}

		if err := weci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowExecutionDeleteInput holds the deletion input of the WorkflowExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowExecutionDeleteInput struct {
	WorkflowExecutionQueryInput `path:",inline"`
}

// WorkflowExecutionDeleteInputs holds the deletion input item of the WorkflowExecution entities.
type WorkflowExecutionDeleteInputsItem struct {
	// ID of the WorkflowExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// WorkflowExecutionDeleteInputs holds the deletion input of the WorkflowExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowExecutionDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to delete WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowExecutionDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowExecution entities for deleting,
// after validating.
func (wedi *WorkflowExecutionDeleteInputs) Model() []*WorkflowExecution {
	if wedi == nil || len(wedi.Items) == 0 {
		return nil
	}

	_wes := make([]*WorkflowExecution, len(wedi.Items))
	for i := range wedi.Items {
		_wes[i] = &WorkflowExecution{
			ID: wedi.Items[i].ID,
		}
	}
	return _wes
}

// IDs returns the ID list of the WorkflowExecution entities for deleting,
// after validating.
func (wedi *WorkflowExecutionDeleteInputs) IDs() []object.ID {
	if wedi == nil || len(wedi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wedi.Items))
	for i := range wedi.Items {
		ids[i] = wedi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowExecutionDeleteInputs entity.
func (wedi *WorkflowExecutionDeleteInputs) Validate() error {
	if wedi == nil {
		return errors.New("nil receiver")
	}

	return wedi.ValidateWith(wedi.inputConfig.Context, wedi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionDeleteInputs entity with the given context and client set.
func (wedi *WorkflowExecutionDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wedi == nil {
		return errors.New("nil receiver")
	}

	if len(wedi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowExecutions().Query()

	// Validate when deleting under the Workflow route.
	if wedi.Workflow != nil {
		if err := wedi.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowexecution.WorkflowID(wedi.Workflow.ID))
		}
	}

	ids := make([]object.ID, 0, len(wedi.Items))

	for i := range wedi.Items {
		if wedi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wedi.Items[i].ID != "" {
			ids = append(ids, wedi.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	return nil
}

// WorkflowExecutionQueryInput holds the query input of the WorkflowExecution entity,
// please tags with `path:",inline"` if embedding.
type WorkflowExecutionQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to query WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"workflow"`

	// Refer holds the route path reference of the WorkflowExecution entity.
	Refer *object.Refer `path:"workflowexecution,default=" query:"-" json:"-"`
	// ID of the WorkflowExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`
}

// Model returns the WorkflowExecution entity for querying,
// after validating.
func (weqi *WorkflowExecutionQueryInput) Model() *WorkflowExecution {
	if weqi == nil {
		return nil
	}

	return &WorkflowExecution{
		ID: weqi.ID,
	}
}

// Validate checks the WorkflowExecutionQueryInput entity.
func (weqi *WorkflowExecutionQueryInput) Validate() error {
	if weqi == nil {
		return errors.New("nil receiver")
	}

	return weqi.ValidateWith(weqi.inputConfig.Context, weqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionQueryInput entity with the given context and client set.
func (weqi *WorkflowExecutionQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weqi == nil {
		return errors.New("nil receiver")
	}

	if weqi.Refer != nil && *weqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflowexecution.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowExecutions().Query()

	// Validate when querying under the Workflow route.
	if weqi.Workflow != nil {
		if err := weqi.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowexecution.WorkflowID(weqi.Workflow.ID))
		}
	}

	if weqi.Refer != nil {
		if weqi.Refer.IsID() {
			q.Where(
				workflowexecution.ID(weqi.Refer.ID()))
		} else {
			return errors.New("invalid identify refer of workflowexecution")
		}
	} else if weqi.ID != "" {
		q.Where(
			workflowexecution.ID(weqi.ID))
	} else {
		return errors.New("invalid identify of workflowexecution")
	}

	q.Select(
		workflowexecution.FieldID,
	)

	var e *WorkflowExecution
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
			e = cv.(*WorkflowExecution)
		}
	}

	weqi.ID = e.ID
	return nil
}

// WorkflowExecutionQueryInputs holds the query input of the WorkflowExecution entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowExecutionQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to query WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the WorkflowExecutionQueryInputs entity.
func (weqi *WorkflowExecutionQueryInputs) Validate() error {
	if weqi == nil {
		return errors.New("nil receiver")
	}

	return weqi.ValidateWith(weqi.inputConfig.Context, weqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionQueryInputs entity with the given context and client set.
func (weqi *WorkflowExecutionQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Workflow route.
	if weqi.Workflow != nil {
		if err := weqi.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowExecutionUpdateInput holds the modification input of the WorkflowExecution entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowExecutionUpdateInput struct {
	WorkflowExecutionQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID list of the stage executions that belong to this workflow execution.
	WorkflowStagesExecution []object.ID `path:"-" query:"-" json:"workflowStagesExecution,omitempty"`
	// Log record of the workflow execution.
	Record string `path:"-" query:"-" json:"record,omitempty"`
	// Input of the workflow execution. It's the yaml file that defines the workflow execution.
	Input string `path:"-" query:"-" json:"input,omitempty"`
}

// Model returns the WorkflowExecution entity for modifying,
// after validating.
func (weui *WorkflowExecutionUpdateInput) Model() *WorkflowExecution {
	if weui == nil {
		return nil
	}

	_we := &WorkflowExecution{
		ID:                      weui.ID,
		Description:             weui.Description,
		Labels:                  weui.Labels,
		WorkflowStagesExecution: weui.WorkflowStagesExecution,
		Record:                  weui.Record,
		Input:                   weui.Input,
	}

	return _we
}

// Validate checks the WorkflowExecutionUpdateInput entity.
func (weui *WorkflowExecutionUpdateInput) Validate() error {
	if weui == nil {
		return errors.New("nil receiver")
	}

	return weui.ValidateWith(weui.inputConfig.Context, weui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionUpdateInput entity with the given context and client set.
func (weui *WorkflowExecutionUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := weui.WorkflowExecutionQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// WorkflowExecutionUpdateInputs holds the modification input item of the WorkflowExecution entities.
type WorkflowExecutionUpdateInputsItem struct {
	// ID of the WorkflowExecution entity.
	ID object.ID `path:"-" query:"-" json:"id"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID list of the stage executions that belong to this workflow execution.
	WorkflowStagesExecution []object.ID `path:"-" query:"-" json:"workflowStagesExecution"`
	// Log record of the workflow execution.
	Record string `path:"-" query:"-" json:"record"`
	// Input of the workflow execution. It's the yaml file that defines the workflow execution.
	Input string `path:"-" query:"-" json:"input"`
}

// ValidateWith checks the WorkflowExecutionUpdateInputsItem entity with the given context and client set.
func (weui *WorkflowExecutionUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// WorkflowExecutionUpdateInputs holds the modification input of the WorkflowExecution entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowExecutionUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Workflow indicates to update WorkflowExecution entity MUST under the Workflow route.
	Workflow *WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowExecutionUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the WorkflowExecution entities for modifying,
// after validating.
func (weui *WorkflowExecutionUpdateInputs) Model() []*WorkflowExecution {
	if weui == nil || len(weui.Items) == 0 {
		return nil
	}

	_wes := make([]*WorkflowExecution, len(weui.Items))

	for i := range weui.Items {
		_we := &WorkflowExecution{
			ID:                      weui.Items[i].ID,
			Description:             weui.Items[i].Description,
			Labels:                  weui.Items[i].Labels,
			WorkflowStagesExecution: weui.Items[i].WorkflowStagesExecution,
			Record:                  weui.Items[i].Record,
			Input:                   weui.Items[i].Input,
		}

		_wes[i] = _we
	}

	return _wes
}

// IDs returns the ID list of the WorkflowExecution entities for modifying,
// after validating.
func (weui *WorkflowExecutionUpdateInputs) IDs() []object.ID {
	if weui == nil || len(weui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(weui.Items))
	for i := range weui.Items {
		ids[i] = weui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowExecutionUpdateInputs entity.
func (weui *WorkflowExecutionUpdateInputs) Validate() error {
	if weui == nil {
		return errors.New("nil receiver")
	}

	return weui.ValidateWith(weui.inputConfig.Context, weui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowExecutionUpdateInputs entity with the given context and client set.
func (weui *WorkflowExecutionUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if weui == nil {
		return errors.New("nil receiver")
	}

	if len(weui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.WorkflowExecutions().Query()

	// Validate when updating under the Workflow route.
	if weui.Workflow != nil {
		if err := weui.Workflow.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			q.Where(
				workflowexecution.WorkflowID(weui.Workflow.ID))
		}
	}

	ids := make([]object.ID, 0, len(weui.Items))

	for i := range weui.Items {
		if weui.Items[i] == nil {
			return errors.New("nil item")
		}

		if weui.Items[i].ID != "" {
			ids = append(ids, weui.Items[i].ID)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	if len(ids) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	idsCnt, err := q.Where(workflowexecution.IDIn(ids...)).
		Count(ctx)
	if err != nil {
		return err
	}

	if idsCnt != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range weui.Items {
		if err := weui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowExecutionOutput holds the output of the WorkflowExecution entity.
type WorkflowExecutionOutput struct {
	ID                      object.ID         `json:"id,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Description             string            `json:"description,omitempty"`
	Labels                  map[string]string `json:"labels,omitempty"`
	CreateTime              *time.Time        `json:"createTime,omitempty"`
	UpdateTime              *time.Time        `json:"updateTime,omitempty"`
	Status                  status.Status     `json:"status,omitempty"`
	ProjectID               object.ID         `json:"projectID,omitempty"`
	Subject                 object.ID         `json:"subject,omitempty"`
	Progress                int               `json:"progress,omitempty"`
	Duration                int               `json:"duration,omitempty"`
	WorkflowStagesExecution []object.ID       `json:"workflowStagesExecution,omitempty"`
	Record                  string            `json:"record,omitempty"`
	Input                   string            `json:"input,omitempty"`

	Workflow *WorkflowOutput `json:"workflow,omitempty"`
}

// View returns the output of WorkflowExecution entity.
func (_we *WorkflowExecution) View() *WorkflowExecutionOutput {
	return ExposeWorkflowExecution(_we)
}

// View returns the output of WorkflowExecution entities.
func (_wes WorkflowExecutions) View() []*WorkflowExecutionOutput {
	return ExposeWorkflowExecutions(_wes)
}

// ExposeWorkflowExecution converts the WorkflowExecution to WorkflowExecutionOutput.
func ExposeWorkflowExecution(_we *WorkflowExecution) *WorkflowExecutionOutput {
	if _we == nil {
		return nil
	}

	weo := &WorkflowExecutionOutput{
		ID:                      _we.ID,
		Name:                    _we.Name,
		Description:             _we.Description,
		Labels:                  _we.Labels,
		CreateTime:              _we.CreateTime,
		UpdateTime:              _we.UpdateTime,
		Status:                  _we.Status,
		ProjectID:               _we.ProjectID,
		Subject:                 _we.Subject,
		Progress:                _we.Progress,
		Duration:                _we.Duration,
		WorkflowStagesExecution: _we.WorkflowStagesExecution,
		Record:                  _we.Record,
		Input:                   _we.Input,
	}

	if _we.Edges.Workflow != nil {
		weo.Workflow = ExposeWorkflow(_we.Edges.Workflow)
	} else if _we.WorkflowID != "" {
		weo.Workflow = &WorkflowOutput{
			ID: _we.WorkflowID,
		}
	}
	return weo
}

// ExposeWorkflowExecutions converts the WorkflowExecution slice to WorkflowExecutionOutput pointer slice.
func ExposeWorkflowExecutions(_wes []*WorkflowExecution) []*WorkflowExecutionOutput {
	if len(_wes) == 0 {
		return nil
	}

	weos := make([]*WorkflowExecutionOutput, len(_wes))
	for i := range _wes {
		weos[i] = ExposeWorkflowExecution(_wes[i])
	}
	return weos
}
