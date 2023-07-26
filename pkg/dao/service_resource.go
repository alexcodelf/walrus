package dao

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresourcerelationship"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/utils/strs"
)

// ServiceResourceInstancesEdgeSave saves the edge instances of model.ServiceResource entity.
func ServiceResourceInstancesEdgeSave(ctx context.Context, mc model.ClientSet, entity *model.ServiceResource) error {
	if entity.Edges.Instances == nil {
		return nil
	}

	// Delete stale items.
	_, err := mc.ServiceResources().Delete().
		Where(serviceresource.ClassID(entity.ID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Add new items.
	newItems := entity.Edges.Instances
	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil relationship")
		}
		newItems[i].ClassID = entity.ID
	}

	newItems, err = mc.ServiceResources().CreateBulk().
		Set(newItems...).
		Save(ctx)
	if err != nil {
		return err
	}

	entity.Edges.Instances = newItems // Feedback.

	return nil
}

// ServiceResourceShapeClassQuery wraps the given model.ServiceResource query
// to select all shape class resources and the owned components and dependencies of them.
func ServiceResourceShapeClassQuery(query *model.ServiceResourceQuery) *model.ServiceResourceQuery {
	var (
		order  = model.Desc(serviceresource.FieldCreateTime)
		wcOpts = func(q *model.ConnectorQuery) {
			q.Select(
				connector.FieldName,
				connector.FieldType,
				connector.FieldCategory,
				connector.FieldConfigVersion,
				connector.FieldConfigData,
			)
		}
	)

	return query.
		Where(
			serviceresource.Shape(types.ServiceResourceShapeClass),
			serviceresource.Mode(types.ServiceResourceModeManaged)).
		Order(order).
		WithInstances(func(iq *model.ServiceResourceQuery) {
			iq.
				Order(order).
				WithConnector(wcOpts).
				WithComponents(func(cq *model.ServiceResourceQuery) {
					cq.
						Order(order).
						WithConnector(wcOpts)
				})
		}).
		WithDependencies(func(rrq *model.ServiceResourceRelationshipQuery) {
			rrq.Select(
				serviceresourcerelationship.FieldServiceResourceID,
				serviceresourcerelationship.FieldDependencyID,
				serviceresourcerelationship.FieldType)
		})
}

// ServiceResourceToMap recursive set a map of service resources indexed by its unique index.
func ServiceResourceToMap(resources []*model.ServiceResource) map[string]*model.ServiceResource {
	m := make(map[string]*model.ServiceResource)

	stack := make([]*model.ServiceResource, 0)
	stack = append(stack, resources...)

	for len(stack) > 0 {
		res := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		key := ServiceResourceGetUniqueKey(res)
		if _, ok := m[key]; ok {
			continue
		}
		m[key] = res

		stack = append(stack, res.Edges.Components...)
		stack = append(stack, res.Edges.Instances...)
	}

	return m
}

// ServiceResourceGetUniqueKey returns the unique index key of the given model.ServiceResource.
func ServiceResourceGetUniqueKey(r *model.ServiceResource) string {
	// Align to schema definition.
	return strs.Join("-", string(r.ConnectorID), r.Shape, r.Mode, r.Type, r.Name)
}

func ServiceResourceUpdates(
	mc model.ClientSet,
	input ...*model.ServiceResource,
) ([]*model.ServiceResourceUpdateOne, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid input: empty list")
	}

	rrs := make([]*model.ServiceResourceUpdateOne, len(input))

	for i, r := range input {
		if r == nil {
			return nil, errors.New("invalid input: nil entity")
		}

		// Required.
		u := mc.ServiceResources().UpdateOneID(r.ID).
			SetDriftResult(r.DriftResult)

		rrs[i] = u
	}

	return rrs, nil
}
