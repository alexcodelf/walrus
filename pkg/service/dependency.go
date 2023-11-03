package service

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/terraform/parser"
)

// GetServiceDependencyOutputsByID gets the dependency outputs of the service by service id.
func GetServiceDependencyOutputsByID(
	ctx context.Context,
	client model.ClientSet,
	serviceID object.ID,
	dependOutputs []string,
) (map[string]parser.OutputState, error) {
	entity, err := client.Services().Query().
		Where(service.ID(serviceID)).
		WithDependencies(func(sq *model.ServiceRelationshipQuery) {
			sq.Where(func(s *sql.Selector) {
				s.Where(sql.ColumnsNEQ(servicerelationship.FieldServiceID, servicerelationship.FieldDependencyID))
			})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	dependencyServiceIDs := make([]object.ID, 0, len(entity.Edges.Dependencies))

	for _, d := range entity.Edges.Dependencies {
		if d.Type != types.ServiceRelationshipTypeImplicit {
			continue
		}

		dependencyServiceIDs = append(dependencyServiceIDs, d.DependencyID)
	}

	return getServiceDependencyOutputs(ctx, client, dependencyServiceIDs, dependOutputs)
}

// getServiceDependencyOutputs gets the dependency outputs of the service.
func getServiceDependencyOutputs(
	ctx context.Context,
	client model.ClientSet,
	dependencyServiceIDs []object.ID,
	dependOutputs []string,
) (map[string]parser.OutputState, error) {
	dependencyRevisions, err := client.ServiceRevisions().Query().
		Select(
			servicerevision.FieldID,
			servicerevision.FieldAttributes,
			servicerevision.FieldOutput,
			servicerevision.FieldServiceID,
			servicerevision.FieldProjectID,
		).
		Where(func(s *sql.Selector) {
			sq := s.Clone().
				AppendSelectExprAs(
					sql.RowNumber().
						PartitionBy(servicerevision.FieldServiceID).
						OrderBy(sql.Desc(servicerevision.FieldCreateTime)),
					"row_number",
				).
				Where(s.P()).
				From(s.Table()).
				As(servicerevision.Table)

			// Query the latest revision of the service.
			s.Where(sql.EQ(s.C("row_number"), 1)).
				From(sq)
		}).Where(servicerevision.ServiceIDIn(dependencyServiceIDs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make(map[string]parser.OutputState, 0)
	dependSets := sets.NewString(dependOutputs...)

	for _, r := range dependencyRevisions {
		revisionOutput, err := parser.ParseStateOutputRawMap(r)
		if err != nil {
			return nil, err
		}

		for n, o := range revisionOutput {
			if dependSets.Has(n) {
				outputs[n] = o
			}
		}
	}

	return outputs, nil
}
