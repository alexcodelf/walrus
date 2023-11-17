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
	"github.com/seal-io/walrus/pkg/dao/model/resourcecomponentrelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceComponentRelationshipCreateInput holds the creation input of the ResourceComponentRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentRelationshipCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Type of the relationship.
	Type string `path:"-" query:"-" json:"type"`

	// Dependency specifies full inserting the new ResourceComponent entity of the ResourceComponentRelationship entity.
	Dependency *ResourceComponentQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ResourceComponentRelationship entity for creating,
// after validating.
func (rcrci *ResourceComponentRelationshipCreateInput) Model() *ResourceComponentRelationship {
	if rcrci == nil {
		return nil
	}

	_rcr := &ResourceComponentRelationship{
		Type: rcrci.Type,
	}

	if rcrci.Dependency != nil {
		_rcr.DependencyID = rcrci.Dependency.ID
	}
	return _rcr
}

// Validate checks the ResourceComponentRelationshipCreateInput entity.
func (rcrci *ResourceComponentRelationshipCreateInput) Validate() error {
	if rcrci == nil {
		return errors.New("nil receiver")
	}

	return rcrci.ValidateWith(rcrci.inputConfig.Context, rcrci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipCreateInput entity with the given context and client set.
func (rcrci *ResourceComponentRelationshipCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if rcrci.Dependency != nil {
		if err := rcrci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcrci.Dependency = nil
			}
		}
	}

	return nil
}

// ResourceComponentRelationshipCreateInputs holds the creation input item of the ResourceComponentRelationship entities.
type ResourceComponentRelationshipCreateInputsItem struct {
	// Type of the relationship.
	Type string `path:"-" query:"-" json:"type"`

	// Dependency specifies full inserting the new ResourceComponent entity.
	Dependency *ResourceComponentQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ResourceComponentRelationshipCreateInputsItem entity with the given context and client set.
func (rcrci *ResourceComponentRelationshipCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if rcrci.Dependency != nil {
		if err := rcrci.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcrci.Dependency = nil
			}
		}
	}

	return nil
}

// ResourceComponentRelationshipCreateInputs holds the creation input of the ResourceComponentRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentRelationshipCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentRelationshipCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponentRelationship entities for creating,
// after validating.
func (rcrci *ResourceComponentRelationshipCreateInputs) Model() []*ResourceComponentRelationship {
	if rcrci == nil || len(rcrci.Items) == 0 {
		return nil
	}

	_rcrs := make([]*ResourceComponentRelationship, len(rcrci.Items))

	for i := range rcrci.Items {
		_rcr := &ResourceComponentRelationship{
			Type: rcrci.Items[i].Type,
		}

		if rcrci.Items[i].Dependency != nil {
			_rcr.DependencyID = rcrci.Items[i].Dependency.ID
		}

		_rcrs[i] = _rcr
	}

	return _rcrs
}

// Validate checks the ResourceComponentRelationshipCreateInputs entity .
func (rcrci *ResourceComponentRelationshipCreateInputs) Validate() error {
	if rcrci == nil {
		return errors.New("nil receiver")
	}

	return rcrci.ValidateWith(rcrci.inputConfig.Context, rcrci.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipCreateInputs entity with the given context and client set.
func (rcrci *ResourceComponentRelationshipCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrci == nil {
		return errors.New("nil receiver")
	}

	if len(rcrci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range rcrci.Items {
		if rcrci.Items[i] == nil {
			continue
		}

		if err := rcrci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceComponentRelationshipDeleteInput holds the deletion input of the ResourceComponentRelationship entity,
// please tags with `path:",inline"` if embedding.
type ResourceComponentRelationshipDeleteInput struct {
	ResourceComponentRelationshipQueryInput `path:",inline"`
}

// ResourceComponentRelationshipDeleteInputs holds the deletion input item of the ResourceComponentRelationship entities.
type ResourceComponentRelationshipDeleteInputsItem struct {
	// ID of the ResourceComponentRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ResourceComponentRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`
}

// ResourceComponentRelationshipDeleteInputs holds the deletion input of the ResourceComponentRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentRelationshipDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentRelationshipDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponentRelationship entities for deleting,
// after validating.
func (rcrdi *ResourceComponentRelationshipDeleteInputs) Model() []*ResourceComponentRelationship {
	if rcrdi == nil || len(rcrdi.Items) == 0 {
		return nil
	}

	_rcrs := make([]*ResourceComponentRelationship, len(rcrdi.Items))
	for i := range rcrdi.Items {
		_rcrs[i] = &ResourceComponentRelationship{
			ID: rcrdi.Items[i].ID,
		}
	}
	return _rcrs
}

// IDs returns the ID list of the ResourceComponentRelationship entities for deleting,
// after validating.
func (rcrdi *ResourceComponentRelationshipDeleteInputs) IDs() []object.ID {
	if rcrdi == nil || len(rcrdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rcrdi.Items))
	for i := range rcrdi.Items {
		ids[i] = rcrdi.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceComponentRelationshipDeleteInputs entity.
func (rcrdi *ResourceComponentRelationshipDeleteInputs) Validate() error {
	if rcrdi == nil {
		return errors.New("nil receiver")
	}

	return rcrdi.ValidateWith(rcrdi.inputConfig.Context, rcrdi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipDeleteInputs entity with the given context and client set.
func (rcrdi *ResourceComponentRelationshipDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrdi == nil {
		return errors.New("nil receiver")
	}

	if len(rcrdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponentRelationships().Query()

	ids := make([]object.ID, 0, len(rcrdi.Items))
	ors := make([]predicate.ResourceComponentRelationship, 0, len(rcrdi.Items))
	indexers := make(map[any][]int)

	for i := range rcrdi.Items {
		if rcrdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if rcrdi.Items[i].ID != "" {
			ids = append(ids, rcrdi.Items[i].ID)
			ors = append(ors, resourcecomponentrelationship.ID(rcrdi.Items[i].ID))
			indexers[rcrdi.Items[i].ID] = append(indexers[rcrdi.Items[i].ID], i)
		} else if rcrdi.Items[i].Type != "" {
			ors = append(ors, resourcecomponentrelationship.And(
				resourcecomponentrelationship.Type(rcrdi.Items[i].Type)))
			indexerKey := fmt.Sprint("/", rcrdi.Items[i].Type)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := resourcecomponentrelationship.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = resourcecomponentrelationship.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			resourcecomponentrelationship.FieldID,
			resourcecomponentrelationship.FieldType,
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
			indexerKey := fmt.Sprint("/", es[i].Type)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			rcrdi.Items[j].ID = es[i].ID
			rcrdi.Items[j].Type = es[i].Type
		}
	}

	return nil
}

// ResourceComponentRelationshipQueryInput holds the query input of the ResourceComponentRelationship entity,
// please tags with `path:",inline"` if embedding.
type ResourceComponentRelationshipQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Refer holds the route path reference of the ResourceComponentRelationship entity.
	Refer *object.Refer `path:"resourcecomponentrelationship,default=" query:"-" json:"-"`
	// ID of the ResourceComponentRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ResourceComponentRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`
}

// Model returns the ResourceComponentRelationship entity for querying,
// after validating.
func (rcrqi *ResourceComponentRelationshipQueryInput) Model() *ResourceComponentRelationship {
	if rcrqi == nil {
		return nil
	}

	return &ResourceComponentRelationship{
		ID:   rcrqi.ID,
		Type: rcrqi.Type,
	}
}

// Validate checks the ResourceComponentRelationshipQueryInput entity.
func (rcrqi *ResourceComponentRelationshipQueryInput) Validate() error {
	if rcrqi == nil {
		return errors.New("nil receiver")
	}

	return rcrqi.ValidateWith(rcrqi.inputConfig.Context, rcrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipQueryInput entity with the given context and client set.
func (rcrqi *ResourceComponentRelationshipQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrqi == nil {
		return errors.New("nil receiver")
	}

	if rcrqi.Refer != nil && *rcrqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", resourcecomponentrelationship.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponentRelationships().Query()

	if rcrqi.Refer != nil {
		if rcrqi.Refer.IsID() {
			q.Where(
				resourcecomponentrelationship.ID(rcrqi.Refer.ID()))
		} else if refers := rcrqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				resourcecomponentrelationship.Type(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of resourcecomponentrelationship")
		}
	} else if rcrqi.ID != "" {
		q.Where(
			resourcecomponentrelationship.ID(rcrqi.ID))
	} else if rcrqi.Type != "" {
		q.Where(
			resourcecomponentrelationship.Type(rcrqi.Type))
	} else {
		return errors.New("invalid identify of resourcecomponentrelationship")
	}

	q.Select(
		resourcecomponentrelationship.FieldID,
		resourcecomponentrelationship.FieldType,
	)

	var e *ResourceComponentRelationship
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
			e = cv.(*ResourceComponentRelationship)
		}
	}

	rcrqi.ID = e.ID
	rcrqi.Type = e.Type
	return nil
}

// ResourceComponentRelationshipQueryInputs holds the query input of the ResourceComponentRelationship entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type ResourceComponentRelationshipQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`
}

// Validate checks the ResourceComponentRelationshipQueryInputs entity.
func (rcrqi *ResourceComponentRelationshipQueryInputs) Validate() error {
	if rcrqi == nil {
		return errors.New("nil receiver")
	}

	return rcrqi.ValidateWith(rcrqi.inputConfig.Context, rcrqi.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipQueryInputs entity with the given context and client set.
func (rcrqi *ResourceComponentRelationshipQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// ResourceComponentRelationshipUpdateInput holds the modification input of the ResourceComponentRelationship entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentRelationshipUpdateInput struct {
	ResourceComponentRelationshipQueryInput `path:",inline" query:"-" json:"-"`

	// Dependency indicates replacing the stale ResourceComponent entity.
	Dependency *ResourceComponentQueryInput `uri:"-" query:"-" json:"dependency"`
}

// Model returns the ResourceComponentRelationship entity for modifying,
// after validating.
func (rcrui *ResourceComponentRelationshipUpdateInput) Model() *ResourceComponentRelationship {
	if rcrui == nil {
		return nil
	}

	_rcr := &ResourceComponentRelationship{
		ID:   rcrui.ID,
		Type: rcrui.Type,
	}

	if rcrui.Dependency != nil {
		_rcr.DependencyID = rcrui.Dependency.ID
	}
	return _rcr
}

// Validate checks the ResourceComponentRelationshipUpdateInput entity.
func (rcrui *ResourceComponentRelationshipUpdateInput) Validate() error {
	if rcrui == nil {
		return errors.New("nil receiver")
	}

	return rcrui.ValidateWith(rcrui.inputConfig.Context, rcrui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipUpdateInput entity with the given context and client set.
func (rcrui *ResourceComponentRelationshipUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := rcrui.ResourceComponentRelationshipQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	if rcrui.Dependency != nil {
		if err := rcrui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcrui.Dependency = nil
			}
		}
	}

	return nil
}

// ResourceComponentRelationshipUpdateInputs holds the modification input item of the ResourceComponentRelationship entities.
type ResourceComponentRelationshipUpdateInputsItem struct {
	// ID of the ResourceComponentRelationship entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Type of the ResourceComponentRelationship entity, a part of the unique index.
	Type string `path:"-" query:"-" json:"type,omitempty"`

	// Dependency indicates replacing the stale ResourceComponent entity.
	Dependency *ResourceComponentQueryInput `uri:"-" query:"-" json:"dependency"`
}

// ValidateWith checks the ResourceComponentRelationshipUpdateInputsItem entity with the given context and client set.
func (rcrui *ResourceComponentRelationshipUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	if rcrui.Dependency != nil {
		if err := rcrui.Dependency.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				rcrui.Dependency = nil
			}
		}
	}

	return nil
}

// ResourceComponentRelationshipUpdateInputs holds the modification input of the ResourceComponentRelationship entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type ResourceComponentRelationshipUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*ResourceComponentRelationshipUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the ResourceComponentRelationship entities for modifying,
// after validating.
func (rcrui *ResourceComponentRelationshipUpdateInputs) Model() []*ResourceComponentRelationship {
	if rcrui == nil || len(rcrui.Items) == 0 {
		return nil
	}

	_rcrs := make([]*ResourceComponentRelationship, len(rcrui.Items))

	for i := range rcrui.Items {
		_rcr := &ResourceComponentRelationship{
			ID:   rcrui.Items[i].ID,
			Type: rcrui.Items[i].Type,
		}

		if rcrui.Items[i].Dependency != nil {
			_rcr.DependencyID = rcrui.Items[i].Dependency.ID
		}

		_rcrs[i] = _rcr
	}

	return _rcrs
}

// IDs returns the ID list of the ResourceComponentRelationship entities for modifying,
// after validating.
func (rcrui *ResourceComponentRelationshipUpdateInputs) IDs() []object.ID {
	if rcrui == nil || len(rcrui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(rcrui.Items))
	for i := range rcrui.Items {
		ids[i] = rcrui.Items[i].ID
	}
	return ids
}

// Validate checks the ResourceComponentRelationshipUpdateInputs entity.
func (rcrui *ResourceComponentRelationshipUpdateInputs) Validate() error {
	if rcrui == nil {
		return errors.New("nil receiver")
	}

	return rcrui.ValidateWith(rcrui.inputConfig.Context, rcrui.inputConfig.Client, nil)
}

// ValidateWith checks the ResourceComponentRelationshipUpdateInputs entity with the given context and client set.
func (rcrui *ResourceComponentRelationshipUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if rcrui == nil {
		return errors.New("nil receiver")
	}

	if len(rcrui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.ResourceComponentRelationships().Query()

	ids := make([]object.ID, 0, len(rcrui.Items))
	ors := make([]predicate.ResourceComponentRelationship, 0, len(rcrui.Items))
	indexers := make(map[any][]int)

	for i := range rcrui.Items {
		if rcrui.Items[i] == nil {
			return errors.New("nil item")
		}

		if rcrui.Items[i].ID != "" {
			ids = append(ids, rcrui.Items[i].ID)
			ors = append(ors, resourcecomponentrelationship.ID(rcrui.Items[i].ID))
			indexers[rcrui.Items[i].ID] = append(indexers[rcrui.Items[i].ID], i)
		} else if rcrui.Items[i].Type != "" {
			ors = append(ors, resourcecomponentrelationship.And(
				resourcecomponentrelationship.Type(rcrui.Items[i].Type)))
			indexerKey := fmt.Sprint("/", rcrui.Items[i].Type)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := resourcecomponentrelationship.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = resourcecomponentrelationship.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			resourcecomponentrelationship.FieldID,
			resourcecomponentrelationship.FieldType,
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
			indexerKey := fmt.Sprint("/", es[i].Type)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			rcrui.Items[j].ID = es[i].ID
			rcrui.Items[j].Type = es[i].Type
		}
	}

	for i := range rcrui.Items {
		if err := rcrui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// ResourceComponentRelationshipOutput holds the output of the ResourceComponentRelationship entity.
type ResourceComponentRelationshipOutput struct {
	ID         object.ID  `json:"id,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	Type       string     `json:"type,omitempty"`

	Dependency *ResourceComponentOutput `json:"dependency,omitempty"`
}

// View returns the output of ResourceComponentRelationship entity.
func (_rcr *ResourceComponentRelationship) View() *ResourceComponentRelationshipOutput {
	return ExposeResourceComponentRelationship(_rcr)
}

// View returns the output of ResourceComponentRelationship entities.
func (_rcrs ResourceComponentRelationships) View() []*ResourceComponentRelationshipOutput {
	return ExposeResourceComponentRelationships(_rcrs)
}

// ExposeResourceComponentRelationship converts the ResourceComponentRelationship to ResourceComponentRelationshipOutput.
func ExposeResourceComponentRelationship(_rcr *ResourceComponentRelationship) *ResourceComponentRelationshipOutput {
	if _rcr == nil {
		return nil
	}

	rcro := &ResourceComponentRelationshipOutput{
		ID:         _rcr.ID,
		CreateTime: _rcr.CreateTime,
		Type:       _rcr.Type,
	}

	if _rcr.Edges.Dependency != nil {
		rcro.Dependency = ExposeResourceComponent(_rcr.Edges.Dependency)
	} else if _rcr.DependencyID != "" {
		rcro.Dependency = &ResourceComponentOutput{
			ID: _rcr.DependencyID,
		}
	}
	return rcro
}

// ExposeResourceComponentRelationships converts the ResourceComponentRelationship slice to ResourceComponentRelationshipOutput pointer slice.
func ExposeResourceComponentRelationships(_rcrs []*ResourceComponentRelationship) []*ResourceComponentRelationshipOutput {
	if len(_rcrs) == 0 {
		return nil
	}

	rcros := make([]*ResourceComponentRelationshipOutput, len(_rcrs))
	for i := range _rcrs {
		rcros[i] = ExposeResourceComponentRelationship(_rcrs[i])
	}
	return rcros
}
